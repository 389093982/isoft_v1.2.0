package pis

import (
	"errors"
	"fmt"
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkdata/param"
	"isoft/isoft_iwork_web/core/iworkfunc"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkutil"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/core/iworkvalid"
	"isoft/isoft_iwork_web/models/iwork"
	"reflect"
	"strings"
)

// iworkmodels.ParamInputSchemaItem 解析类
type PisItemDataParser struct {
	DataStore          *datastore.DataStore
	Item               iworkmodels.ParamInputSchemaItem
	PureTextTmpDataMap map[string]string
	TmpDataMap         map[string]interface{}
}

func (this *PisItemDataParser) FillPisItemDataToTmp() {
	this.checkEmpty()
	this.FillPisItemDataToPureTmp()
	this.FillPisItemDataToNPureTmp()
}

func (this *PisItemDataParser) FillPisItemDataToPureTmp() {
	this.PureTextTmpDataMap[this.Item.ParamName] = this.Item.ParamValue
}

func (this *PisItemDataParser) checkEmpty() {
	// 对参数进行非空校验
	if ok, checkResults := iworkvalid.CheckEmptyForItem(this.Item); !ok {
		panic(strings.Join(checkResults, ";"))
	}
}

func (this *PisItemDataParser) FillPisItemDataToNPureTmp() {
	// tmpDataMap 存储解析值
	if this.Item.PureText {
		this.TmpDataMap[this.Item.ParamName] = this.Item.ParamValue
	} else {
		// 判断当前参数是否是 repeat 参数
		if !this.Item.Repeatable {
			this.TmpDataMap[this.Item.ParamName] = this.ParseAndGetParamVaule(this.Item.ParamName, this.Item.ParamValue) // 输入数据存临时
		} else {
			this.ForeachFillPisItemDataToTmp()
		}
	}
}

func (this *PisItemDataParser) ForeachFillPisItemDataToTmp() {
	// 获取 item.ForeachRefer 对应的 repeat 切片数据,作为迭代参数,而不再从前置节点输出获取
	repeatDatas := datatypeutil.InterfaceConvertToSlice(this.TmpDataMap[this.Item.ForeachRefer])
	if len(repeatDatas) > 0 {
		paramValues := make([]interface{}, 0)
		for _, repeatData := range repeatDatas {
			// 替代的节点名称、替代的对象
			replaceProviderNodeName := strings.ReplaceAll(strings.TrimSpace(this.PureTextTmpDataMap[this.Item.ForeachRefer]), ";", "")
			replaceMap := map[string]interface{}{replaceProviderNodeName: repeatData}
			paramValue := this.ParseAndGetParamVaule(this.Item.ParamName, this.Item.ParamValue, replaceMap) // 输入数据存临时
			paramValues = append(paramValues, paramValue)
		}
		this.TmpDataMap[this.Item.ParamName] = paramValues // 所得值则是个切片
	} else {
		this.TmpDataMap[this.Item.ParamName] = this.ParseAndGetParamVaule(this.Item.ParamName, this.Item.ParamValue) // 输入数据存临时
	}
}

// 解析 paramVaule 并从 dataStore 中获取实际值
// 可能的情况有多种：单值 interface{}, 多值 []interface{}, 对象值 map[string]interface{}
func (this *PisItemDataParser) ParseAndGetParamVaule(paramName, paramVaule string, replaceMap ...map[string]interface{}) interface{} {
	// 将 paramValue 解析成 []*AttrObjects
	attrObjects := this.parseParamValueToAttrObjects(paramVaule)
	// 存储 []*AttrObjects 转换后的 map[string]interface{}
	resultObjectMap := make(map[string]interface{}, 0)
	// 存储 []*AttrObjects 转换后的 []interface{}
	results := make([]interface{}, 0)
	for _, attrResult := range attrObjects {
		value := this.parseAndGetSingleParamVaule(paramName, attrResult.attrValue, replaceMap...)
		resultObjectMap[attrResult.attrName] = value
		results = append(results, value)
	}
	// 对象值, 将 []*AttrObjects 转换成 map[string]interface{}
	if this.Item.ParamType == "objects" {
		return resultObjectMap
	} else {
		// 单值
		if len(results) == 1 {
			return results[0]
		}
		return results
	}
}

func (this *PisItemDataParser) parseAttrNameAndValueWithSingleParamValue(index int, paramValue string) (attrName string, value string) {
	if strings.Contains(paramValue, "::") {
		attrName := paramValue[:strings.Index(paramValue, "::")]
		value := paramValue[strings.Index(paramValue, "::")+2:]
		return attrName, value
	} else if strings.Contains(paramValue, "$") {
		attrName := strings.ReplaceAll(paramValue[strings.LastIndex(paramValue, ".")+1:], ";", "")
		return attrName, paramValue
	} else {
		return string(index), paramValue
	}
}

type AttrObject struct {
	index     int
	attrName  string
	attrValue string
}

// 将 paramVaule 转行成 对象值 map[string]interface{}, 即 []*AttrObject
func (this *PisItemDataParser) parseParamValueToAttrObjects(paramVaule string) []*AttrObject {
	attrObjects := make([]*AttrObject, 0)
	// 对转义字符 \, \; \( \) 等进行编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	multiVals, err := iworkfunc.SplitWithLexerAnalysis(paramVaule)
	if err != nil {
		panic(err)
	}
	for index, value := range multiVals {
		if _value := this.trim(value); strings.TrimSpace(_value) != "" {
			attrName, value := this.parseAttrNameAndValueWithSingleParamValue(index, strings.TrimSpace(_value))
			attrObjects = append(attrObjects, &AttrObject{index: index, attrName: attrName, attrValue: value})
		}
	}
	return attrObjects
}

func (this *PisItemDataParser) callParseAndGetSingleParamVaule(paramName, paramVaule string, replaceMap ...map[string]interface{}) interface{} {
	paramVaule = iworkfunc.DncodeSpecialForParamVaule(paramVaule)
	// 变量
	if strings.HasPrefix(strings.ToUpper(paramVaule), "$RESOURCE.") {
		return this.parseAndFillParamVauleWithResource(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$WORK.") {
		return iworkutil.GetWorkSubNameFromParamValue(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$ENTITY.") {
		return iworkutil.GetParamValueForEntity(paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(paramVaule), "$") {
		if len(replaceMap) > 0 {
			if paramVaule := this.parseAndFillParamVauleWithReplaceProviderNode(paramVaule, replaceMap...); paramVaule != nil {
				return paramVaule
			}
		}
		return this.parseAndFillParamVauleWithPrefixNode(paramName, paramVaule)
	} else if strings.HasPrefix(paramVaule, "`") && strings.HasSuffix(paramVaule, "`") {
		// 字符串
		return paramVaule[1 : len(paramVaule)-1]
	} else {
		// 数字
		return paramVaule
	}
}

func (this *PisItemDataParser) parseAndGetSingleParamVaule(paramName, paramVaule string, replaceMap ...map[string]interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("<span style='color:red;'>execute func with expression is %s, error msg is :%s</span>", paramVaule, err.(error).Error()))
		}
	}()
	// 对单个 paramVaule 进行特殊字符编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	callers, err := iworkfunc.ParseToFuncCallers(paramVaule)
	if err != nil {
		panic(err)
	}
	if callers == nil || len(callers) == 0 {
		// 是直接参数,不需要函数进行特殊处理
		return this.callParseAndGetSingleParamVaule(paramName, paramVaule, replaceMap...)
	}
	historyFuncResultMap := make(map[string]interface{}, 0)
	var lastFuncResult interface{}
	// 按照顺序依次执行函数
	for _, caller := range callers {
		args := make([]interface{}, 0)
		// 函数参数替换成实际意义上的值
		for _, arg := range caller.FuncArgs {
			// 判断参数是否来源于 historyFuncResultMap
			if _arg, ok := historyFuncResultMap[arg]; ok {
				args = append(args, _arg)
			} else {
				args = append(args, this.callParseAndGetSingleParamVaule(paramName, arg, replaceMap...))
			}
		}
		// 执行函数并记录结果,供下一个函数执行使用
		result := iworkfunc.ExecuteFuncCaller(caller, args)
		historyFuncResultMap["$func."+caller.FuncUUID] = result
		lastFuncResult = result
	}
	return lastFuncResult
}

// paramValue 来源于 iwork 模块
func (this *PisItemDataParser) parseAndFillParamVauleWithResource(paramVaule string) interface{} {
	resource, err := iwork.QueryResourceByName(strings.Replace(paramVaule, "$RESOURCE.", "", -1))
	if err == nil {
		if resource.ResourceType == "db" {
			return resource.ResourceDsn
		} else if resource.ResourceType == "sftp" || resource.ResourceType == "ssh" {
			return resource
		}
	}
	return ""
}

// paramValue 来源于前置节点
func (this *PisItemDataParser) parseAndFillParamVauleWithPrefixNode(paramName, paramVaule string) interface{} {
	// 格式校验
	if !strings.HasPrefix(paramVaule, "$") {
		panic(errors.New(fmt.Sprintf("%s ~ %s is not start with $", paramName, paramVaule)))
	}
	resolver := param.ParamVauleParser{ParamValue: paramVaule}
	nodeName := resolver.GetNodeNameFromParamValue()
	paramName = resolver.GetParamNameFromParamValue()
	paramValue := this.DataStore.GetData(nodeName, paramName) // 作为直接对象, dataStore 里面可以直接获取
	if paramValue != nil {
		return paramValue
	}
	_paramName := paramName[:strings.LastIndex(paramName, ".")]
	datas := this.DataStore.GetData(nodeName, _paramName) // 作为对象属性
	attr := paramName[strings.LastIndex(paramName, ".")+1:]
	if reflect.TypeOf(datas).Kind() == reflect.Slice {
		return reflect.ValueOf(datas).Index(0).Interface().(map[string]interface{})[attr]
	}
	return datas.(map[string]interface{})[attr]
}

func (this *PisItemDataParser) parseAndFillParamVauleWithReplaceProviderNode(paramVaule string, replaceMap ...map[string]interface{}) interface{} {
	for replaceProviderNodeName, replaceProviderData := range replaceMap[0] {
		replaceProviderNodeName = strings.ReplaceAll(replaceProviderNodeName, ";", "")
		if strings.HasPrefix(paramVaule, replaceProviderNodeName) {
			attr := strings.Replace(paramVaule, replaceProviderNodeName+".", "", 1)
			attr = strings.ReplaceAll(attr, ";", "")
			return replaceProviderData.(map[string]interface{})[attr]
		}
	}
	return nil
}

// 去除不合理的字符
func (this *PisItemDataParser) trim(paramValue string) string {
	// 先进行初次的 trim
	paramValue = strings.TrimSpace(paramValue)
	// 去除前后的 \n
	paramValue = strings.TrimPrefix(paramValue, "\n")
	paramValue = strings.TrimSuffix(paramValue, "\n")
	// 再进行二次 trim
	paramValue = strings.TrimSpace(paramValue)
	return paramValue
}
