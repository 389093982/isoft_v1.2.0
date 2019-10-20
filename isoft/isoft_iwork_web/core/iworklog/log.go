package iworklog

import (
	"fmt"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup"
	"time"
)

const cacheLen = 5

type CacheLoggerWriter struct {
	caches   []*models.RunLogDetail
	logOrder int64
}

func (this *CacheLoggerWriter) cleanCaches() {
	this.caches = make([]*models.RunLogDetail, 0)
}

func (this *CacheLoggerWriter) Write(trackingId, workStepName, logLevel, detail string) {
	if this.caches == nil {
		this.cleanCaches()
	}
	this.logOrder++
	log := &models.RunLogDetail{
		TrackingId:      trackingId,
		WorkStepName:    workStepName,
		LogLevel:        logLevel,
		Detail:          detail,
		LogOrder:        this.logOrder,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	this.caches = append(this.caches, log)
	if len(this.caches) >= cacheLen {
		this.flush()
		this.cleanCaches()
	}
}

func (this *CacheLoggerWriter) flush() {
	caches := this.caches // 使用临时变量进行参数传递
	startup.RunLogPool.JobQueue <- func() {
		if _, err := models.InsertMultiRunLogDetail(caches); err != nil {
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@" + err.Error())
		}
	}
}

func (this *CacheLoggerWriter) Close() {
	this.flush()
}

// 统计操作所花费的时间方法
func (this *CacheLoggerWriter) RecordCostTimeLog(operateName, trackingId string, start time.Time) {
	this.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO,
		fmt.Sprintf("%s total cost time :%v ms", operateName, time.Now().Sub(start).Nanoseconds()/1e6))
}
