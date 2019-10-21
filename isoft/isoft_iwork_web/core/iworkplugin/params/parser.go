package params

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
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup/memory"
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
	repeatDatas := make([]interface{}, 0)
	foreachRefer := this.TmpDataMap[this.Item.ForeachRefer]
	if foreachRefer != nil {
		// 获取 item.ForeachRefer 对应的 repeat 切片数据,作为迭代参数,而不再从前置节点输出获取
		repeatDatas = datatypeutil.InterfaceConvertToSlice(foreachRefer)
	}
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
	// 将 paramValue 解析成对象值 []*OobjectAttrs
	objectAttrs := this.parseToObjectAttrs(paramVaule)
	// 存储 []*OobjectAttrs 转换后的 map[string]interface{}
	objectMap := make(map[string]interface{}, 0)
	// 存储 []*AttrObjects 转换后的 []interface{}
	parseValues := make([]interface{}, 0)
	for _, objectAttr := range objectAttrs {
		objectAttr.attrParseValue = this.parseParamVaule(paramName, objectAttr.attrPureValue, replaceMap...)
		// 此处禁止使用 datatypeutil.InterfaceConvertToSlice(), 因为 parseValues 中的元素可以是个 interface{} 也可以是个 []interface{}
		parseValues, objectMap[objectAttr.attrName] = append(parseValues, objectAttr.attrParseValue), objectAttr.attrParseValue
	}
	// 单值
	if len(parseValues) == 1 {
		return parseValues[0]
	}
	// 空值
	if len(parseValues) == 0 {
		return nil
	}
	return parseValues
}

func (this *PisItemDataParser) parseToObjectAttr(index int, paramValue string) *ObjectAttr {
	var attrName, attrPureValue string
	if strings.Contains(paramValue, "::") {
		attrName = paramValue[:strings.Index(paramValue, "::")]
		attrPureValue = paramValue[strings.Index(paramValue, "::")+2:]
	} else if strings.Contains(paramValue, "$") {
		attrName = strings.ReplaceAll(paramValue[strings.LastIndex(paramValue, ".")+1:], ";", "")
		attrPureValue = paramValue
	} else {
		attrName, attrPureValue = string(index), paramValue
	}
	return &ObjectAttr{index: index, attrName: attrName, attrPureValue: attrPureValue}
}

type ObjectAttr struct {
	index          int
	attrName       string      // 对象属性名
	attrPureValue  string      // 对象属性纯文本值
	attrParseValue interface{} // 对象属性解析值
}

// 将 paramVaule 转行成 对象值 map[string]interface{}, 即 []*ObjectAttr
func (this *PisItemDataParser) parseToObjectAttrs(paramVaule string) []*ObjectAttr {
	objectAttrs := make([]*ObjectAttr, 0)
	// 对转义字符 \, \; \( \) 等进行编码
	paramVaule = iworkfunc.EncodeSpecialForParamVaule(paramVaule)
	multiVals, err := iworkfunc.SplitWithLexerAnalysis(paramVaule)
	if err != nil {
		panic(err)
	}
	for index, value := range multiVals {
		if _value := this.trim(value); strings.TrimSpace(_value) != "" {
			objectAttr := this.parseToObjectAttr(index, strings.TrimSpace(_value))
			objectAttrs = append(objectAttrs, objectAttr)
		}
	}
	return objectAttrs
}

func (this *PisItemDataParser) parseParamVaule(paramName, paramVaule string, replaceMap ...map[string]interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			str := "<span style='color:red;'>execute func with expression is %s, error msg is :%s</span>"
			panic(fmt.Sprintf(str, paramVaule, err.(error).Error()))
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
		parser := &SimpleParser{
			paramName:  paramName,
			paramVaule: paramVaule,
			DataStore:  this.DataStore,
		}
		return parser.parseParamValue(replaceMap...)
	} else {
		return this.parseParamVauleWithCallers(callers, paramName, replaceMap...)
	}
}

func (this *PisItemDataParser) parseParamVauleWithCallers(callers []*iworkfunc.FuncCaller, paramName string, replaceMap ...map[string]interface{}) interface{} {
	historyFuncResultMap := make(map[string]interface{}, 0)
	var lastFuncResult interface{}
	// 按照顺序依次执行函数
	for _, caller := range callers {
		args := this.getCallerArgs(caller, historyFuncResultMap, paramName, replaceMap...)
		// 执行函数并记录结果,供下一个函数执行使用
		result := iworkfunc.ExecuteFuncCaller(caller, args)
		historyFuncResultMap["$func."+caller.FuncUUID], lastFuncResult = result, result
	}
	return lastFuncResult
}

// 函数参数替换成实际意义上的值
func (this *PisItemDataParser) getCallerArgs(caller *iworkfunc.FuncCaller,
	historyFuncResultMap map[string]interface{}, paramName string, replaceMap ...map[string]interface{}) []interface{} {
	args := make([]interface{}, 0)
	for _, arg := range caller.FuncArgs {
		// 判断参数是否来源于 historyFuncResultMap
		if _arg, ok := historyFuncResultMap[arg]; ok {
			args = append(args, _arg)
		} else {
			// 是直接参数,不需要函数进行特殊处理
			parser := &SimpleParser{
				paramName:  paramName,
				paramVaule: arg,
				DataStore:  this.DataStore,
			}
			_arg := parser.parseParamValue(replaceMap...)
			args = append(args, _arg)
		}
	}
	return args
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

type SimpleParser struct {
	paramName  string
	paramVaule string
	DataStore  *datastore.DataStore
}

func (this *SimpleParser) parseParamVauleWithReplaceProviderNode(replaceMap ...map[string]interface{}) interface{} {
	for replaceProviderNodeName, replaceProviderData := range replaceMap[0] {
		replaceProviderNodeName = strings.ReplaceAll(replaceProviderNodeName, ";", "")
		if strings.HasPrefix(this.paramVaule, replaceProviderNodeName) {
			attr := strings.Replace(this.paramVaule, replaceProviderNodeName+".", "", 1)
			attr = strings.ReplaceAll(attr, ";", "")
			return replaceProviderData.(map[string]interface{})[attr]
		}
	}
	return nil
}

// paramValue 来源于前置节点
func (this *SimpleParser) parseParamVauleWithPrefixNode() interface{} {
	// 格式校验
	if !strings.HasPrefix(this.paramVaule, "$") {
		panic(errors.New(fmt.Sprintf("%s ~ %s is not start with $", this.paramName, this.paramVaule)))
	}
	resolver := param.ParamVauleParser{ParamValue: this.paramVaule}
	nodeName := resolver.GetNodeNameFromParamValue()
	this.paramName = resolver.GetParamNameFromParamValue()
	paramValue := this.DataStore.GetData(nodeName, this.paramName) // 作为直接对象, dataStore 里面可以直接获取
	if paramValue != nil {
		return paramValue
	}
	if strings.Contains(this.paramName, ".") {
		_paramName := this.paramName[:strings.LastIndex(this.paramName, ".")]
		datas := this.DataStore.GetData(nodeName, _paramName) // 作为对象属性
		if datas == nil {
			return nil // 对象直接不存在，后续不必执行
		}
		attr := this.paramName[strings.LastIndex(this.paramName, ".")+1:]
		if reflect.TypeOf(datas).Kind() == reflect.Slice {
			return reflect.ValueOf(datas).Index(0).Interface().(map[string]interface{})[attr]
		}
		return datas.(map[string]interface{})[attr]
	} else {
		return paramValue
	}
}

// 尽量从缓存中获取
func (this *SimpleParser) parseParamVauleFromResource() interface{} {
	resource_name := strings.TrimPrefix(this.paramVaule, "$RESOURCE.")
	if resource, ok := memory.ResourceMap.Load(resource_name); ok {
		resource := resource.(models.Resource)
		if resource.ResourceType == "db" {
			return resource.ResourceDsn
		} else if resource.ResourceType == "sftp" || resource.ResourceType == "ssh" {
			return resource
		}
		return ""
	} else {
		panic(errors.New(fmt.Sprintf("can't find resource for %s", resource_name)))
	}
}

// 尽量从缓存中获取
func (this *SimpleParser) parseParamVauleFromGlobalVar() interface{} {
	gvName := strings.TrimPrefix(this.paramVaule, "$Global.")
	if gv, ok := memory.GlobalVarMap.Load(gvName); ok {
		return gv.(models.GlobalVar).Value
	} else {
		panic(errors.New(fmt.Sprintf("can't find globalVar for %s", gvName)))
	}
}

// 是直接参数,不需要函数进行特殊处理
func (this *SimpleParser) parseParamValue(replaceMap ...map[string]interface{}) interface{} {
	this.paramVaule = iworkfunc.DncodeSpecialForParamVaule(this.paramVaule)
	// 变量
	if strings.HasPrefix(strings.ToUpper(this.paramVaule), "$GLOBAL.") {
		return this.parseParamVauleFromGlobalVar()
	} else if strings.HasPrefix(strings.ToUpper(this.paramVaule), "$RESOURCE.") {
		return this.parseParamVauleFromResource()
	} else if strings.HasPrefix(strings.ToUpper(this.paramVaule), "$WORK.") {
		return iworkutil.GetWorkSubNameFromParamValue(this.paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(this.paramVaule), "$ENTITY.") {
		return iworkutil.GetParamValueForEntity(this.paramVaule)
	} else if strings.HasPrefix(strings.ToUpper(this.paramVaule), "$") {
		if len(replaceMap) > 0 {
			if paramVaule := this.parseParamVauleWithReplaceProviderNode(replaceMap...); paramVaule != nil {
				return paramVaule
			}
		}
		return this.parseParamVauleWithPrefixNode()
	} else if strings.HasPrefix(this.paramVaule, "`") && strings.HasSuffix(this.paramVaule, "`") {
		// 字符串
		return this.paramVaule[1 : len(this.paramVaule)-1]
	} else {
		// 数字
		return this.paramVaule
	}
}
