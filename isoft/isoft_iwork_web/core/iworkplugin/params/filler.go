package params

import (
	"isoft/isoft_iwork_web/core/iworkdata/datastore"
	"isoft/isoft_iwork_web/core/iworkmodels"
)

// 将 ParamInputSchema 填充数据并返回临时的数据中心 tmpDataMap
func FillParamInputSchemaDataToTmp(pis *iworkmodels.ParamInputSchema, dataStore *datastore.DataStore) map[string]interface{} {
	// 存储节点中间数据
	tmpDataMap, pureTextTmpDataMap := make(map[string]interface{}), make(map[string]string)
	for _, item := range pis.ParamInputSchemaItems {
		fillParamInputSchemaItemDataToTmp(pureTextTmpDataMap, tmpDataMap, item, dataStore)
	}
	return tmpDataMap
}

func fillParamInputSchemaItemDataToTmp(pureTextTmpDataMap map[string]string, tmpDataMap map[string]interface{},
	item iworkmodels.ParamInputSchemaItem, DataStore *datastore.DataStore) {
	parser := &PisItemDataParser{
		DataStore:          DataStore,
		Item:               item,
		PureTextTmpDataMap: pureTextTmpDataMap,
		TmpDataMap:         tmpDataMap,
	}
	parser.FillPisItemDataToTmp()
}
