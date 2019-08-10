package filter

import (
	"github.com/astaxie/beego/context"
	"isoft/isoft_iwork_web/controllers"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
)

func FilterFunc(ctx *context.Context) {
	work_name := ctx.Input.Param(":work_name")
	workCache, err := iworkcache.GetWorkCacheWithName(work_name)
	if err != nil {
		panic(err)
	}
	for _, filterName := range workCache.FilterNames {
		if workCache, err := iworkcache.GetWorkCacheWithName(filterName); err == nil {
			mapData := controllers.ParseParam(ctx, workCache.Steps)
			mapData[iworkconst.HTTP_REQUEST_OBJECT] = ctx.Request // 传递 request 对象
			receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: mapData})
			if receiver != nil {
				tempDataMap := receiver.TmpDataMap
				if data, ok := tempDataMap[iworkconst.TRACKING_ID]; ok {
					ctx.ResponseWriter.Header().Add(iworkconst.TRACKING_ID, data.(string))
				}
				if data, ok := tempDataMap[iworkconst.DO_ERROR_FILTER]; ok {
					tmpDataMap := data.(map[string]interface{})
					for key, value := range tmpDataMap {
						if key == "headerCode" {
							ctx.ResponseWriter.WriteHeader(datatypeutil.InterfaceConvertToInt(value, 200))
						}
					}
				}
			}
		}
	}
}
