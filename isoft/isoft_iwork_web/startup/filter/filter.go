package filter

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"isoft/isoft_iwork_web/controllers"
	"isoft/isoft_iwork_web/core/iworkcache"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkdata/entry"
	"isoft/isoft_iwork_web/core/iworkrun"
	"isoft/isoft_iwork_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup/memory"
	"strings"
)

func FilterFunc(ctx *context.Context) {
	work_name := ctx.Input.Param(":work_name")
	workCache, err := iworkcache.GetWorkCacheWithName(work_name)
	if err != nil {
		panic(err)
	}
	// TODO 是否有序？
	memory.FilterMap.Range(func(k, v interface{}) bool {
		filterWorkName := k.(string)
		fs := v.([]models.Filters)
		if intercept(fs, workCache, ctx) {
			if workCache, err := iworkcache.GetWorkCacheWithName(filterWorkName); err == nil {
				mapData := controllers.ParseParam(ctx, workCache.Steps)
				mapData[iworkconst.HTTP_REQUEST_OBJECT] = ctx.Request // 传递 request 对象
				trackingId, receiver := iworkrun.RunOneWork(workCache.WorkId, &entry.Dispatcher{TmpDataMap: mapData})
				// 将执行过的所有 filter_trackingId 记录到 ctx 中去
				recordFilterStackData(ctx, workCache.Work.WorkName, trackingId)
				ctx.ResponseWriter.Header().Add(iworkconst.TRACKING_ID, trackingId)
				if receiver != nil {
					tempDataMap := receiver.TmpDataMap
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
		// 每个 filter 都要进行判断
		return true
	})
}

func intercept(fs []models.Filters, workCache *iworkcache.WorkCache, ctx *context.Context) bool {
	// fs 有两条记录,一条简单过滤器配置,一条复杂过滤器配置
	for _, filter := range fs {
		if filter.WorkName != "" {
			workNames := strings.Split(filter.WorkName, ",")
			for _, workName := range workNames {
				if workName == workCache.Work.WorkName {
					return true
				}
			}
		}
		if filter.ComplexWorkName != "" {
			complexWorkNames := strings.Split(filter.ComplexWorkName, ",")
			for _, complexWorkName := range complexWorkNames {
				if strings.HasPrefix(complexWorkName, workCache.Work.WorkName) {
					return interceptWithParameter(complexWorkName, ctx)
				}
			}
		}
	}
	return false
}

// 根据参数拦截
// 支持以下几种场景
// workName?paramName=paramValue
func interceptWithParameter(urlpattern string, ctx *context.Context) bool {
	urlparamStr := urlpattern[strings.Index(urlpattern, "?")+1:]
	urlparams := strings.Split(urlparamStr, "=")
	if len(urlparams) == 2 && ctx.Input.Query(urlparams[0]) == urlparams[1] {
		return true
	}
	return false
}

func recordFilterStackData(ctx *context.Context, filterWorkName, trackingId string) {
	filterTrackingIds := ctx.Request.Header.Get(iworkconst.FILTER_TRACKING_ID_STACK)
	filterTrackingIds = fmt.Sprintf("%s,%s[<span style='color:blue;'>%s</span>]", filterTrackingIds, filterWorkName, trackingId)
	filterTrackingIds = strings.TrimPrefix(filterTrackingIds, ",")
	ctx.Request.Header.Set(iworkconst.FILTER_TRACKING_ID_STACK, filterTrackingIds)

}
