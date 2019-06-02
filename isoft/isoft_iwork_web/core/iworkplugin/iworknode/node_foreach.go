package iworknode

import (
	"github.com/pkg/errors"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/block"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkdata/schema"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iwork_web/models/iwork"
	"strings"
)

type ForeachNode struct {
	BaseNode
	WorkStep         *iwork.WorkStep
	BlockStep        *block.BlockStep
	BlockStepRunFunc func(args *iworkprotocol.RunOneStepArgs) (receiver *entry.Receiver)
}

func (this *ForeachNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	if this.BlockStep.HasChildren {
		foreach_datas := tmpDataMap[iworkconst.FOREACH_PREFIX+"foreach_data"].([]map[string]interface{})
		paramMap := make(map[string]interface{})
		for index, foreach_data := range foreach_datas {
			paramMap[iworkconst.NUMBER_PREFIX+"foreach_index"] = index

			for key, value := range foreach_data {
				_key := strings.TrimSpace(strings.ReplaceAll(key, ";", ""))
				_key = _key[strings.LastIndex(_key, ".")+1:]
				paramMap["item."+_key] = value
			}
			this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)

			bsoRunner := BlockStepOrdersRunner{
				ParentStepId: this.WorkStep.WorkStepId,
				WorkCache:    this.WorkCache,
				TrackingId:   trackingId,
				Logwriter:    this.LogWriter,
				Store:        this.DataStore, // 获取数据中心
				Dispatcher:   nil,
				RunOneStep:   this.BlockStepRunFunc,
			}
			bsoRunner.Run()
		}
	} else {
		panic(errors.New("empty foreach was found!"))
	}
}

func (this *ForeachNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.FOREACH_PREFIX + "foreach_data", "迭代的元素"},
		2: {iworkconst.COMPLEX_PREFIX + "foreach_data_attr", "迭代元素属性值"},
	}
	return this.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *ForeachNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	parser := schema.WorkStepSchemaParser{WorkStep: this.WorkStep, ParamSchemaParser: &WorkStepFactory{WorkStep: this.WorkStep}}
	inputSchema := parser.GetCacheParamInputSchema()

	var foreach_data string
	var foreach_data_attr string
	for _, item := range inputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.FOREACH_PREFIX+"foreach_data" {
			foreach_data = item.ParamValue
		} else if item.ParamName == iworkconst.COMPLEX_PREFIX+"foreach_data_attr" {
			foreach_data_attr = item.ParamValue
		}
	}
	foreach_data = strings.TrimSpace(strings.ReplaceAll(foreach_data, ";", ""))
	foreach_data_attr = strings.TrimSpace(strings.ReplaceAll(foreach_data_attr, foreach_data+".", ""))
	attrs := strings.Split(foreach_data_attr, ";")
	for _, attr := range attrs {
		if strings.TrimSpace(attr) != "" {
			// 每个字段放入 items 中
			items = append(items, iworkmodels.ParamOutputSchemaItem{
				ParamName: strings.TrimSpace(attr), ParentPath: "item",
			})
		}
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *ForeachNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BuildParamOutputSchemaWithSlice([]string{iworkconst.NUMBER_PREFIX + "foreach_index"})
}
