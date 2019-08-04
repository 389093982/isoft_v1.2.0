package models

import "time"

type Filters struct {
	Id              int64     `json:"id"`
	FilterWorkId    int64     `json:"filter_work_id"`
	WorkId          int64     `json:"work_id"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}
