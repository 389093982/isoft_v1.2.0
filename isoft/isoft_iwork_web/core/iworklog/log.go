package iworklog

import (
	"fmt"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/models"
	"time"
)

const cacheLen = 10

type CacheLoggerWriter struct {
	caches []*models.RunLogDetail
}

func (this *CacheLoggerWriter) cleanCaches() {
	this.caches = make([]*models.RunLogDetail, 0)
}

func (this *CacheLoggerWriter) Write(trackingId, workStepName, logLevel, detail string) {
	if this.caches == nil {
		this.cleanCaches()
	}
	log := &models.RunLogDetail{
		TrackingId:      trackingId,
		WorkStepName:    workStepName,
		LogLevel:        logLevel,
		Detail:          detail,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	this.caches = append(this.caches, log)
	if len(this.caches) >= cacheLen {
		this.Flush()
		this.cleanCaches()
	}
}

func (this *CacheLoggerWriter) Flush() {
	go func() {
		if _, err := models.InsertMultiRunLogDetail(this.caches); err != nil {
			fmt.Println(err.Error())
		}
	}()
}

func (this *CacheLoggerWriter) Close() {
	this.Flush()
}

// 统计操作所花费的时间方法
func (this *CacheLoggerWriter) RecordCostTimeLog(operateName, trackingId string, start time.Time) {
	this.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO,
		fmt.Sprintf("%s total cost time :%v ms", operateName, time.Now().Sub(start).Nanoseconds()/1e6))
}
