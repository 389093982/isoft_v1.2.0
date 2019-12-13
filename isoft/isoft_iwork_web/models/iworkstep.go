package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"strings"
	"time"
)

type CDATA struct {
	Value string `xml:",cdata"`
}

type WorkStep struct {
	Id                      int64     `json:"id"`
	WorkId                  int64     `json:"work_id"`
	WorkStepId              int64     `json:"work_step_id"`
	WorkSubId               int64     `json:"work_sub_id"` // 子流程 id
	WorkStepName            string    `json:"work_step_name"`
	WorkStepDesc            string    `json:"work_step_desc" orm:"type(text);default('')"`
	WorkStepType            string    `json:"work_step_type"`
	WorkStepIndent          int       `json:"work_step_indent"` // 调整缩进级别
	WorkStepInput           string    `json:"work_step_input" orm:"type(text)"`
	WorkStepInputXml        CDATA     `xml:"work_step_input_xml" json:"work_step_input_xml" orm:"-"`
	WorkStepOutput          string    `json:"work_step_output" orm:"type(text)"`
	WorkStepOutputXml       CDATA     `xml:"work_step_output_xml" json:"work_step_output_xml" orm:"-"`
	IsDefer                 string    `json:"is_defer"`
	WorkStepParamMapping    string    `json:"work_step_param_mapping" orm:"type(text)"`
	WorkStepParamMappingXml CDATA     `xml:"work_step_param_mapping_xml" json:"work_step_param_mapping_xml" orm:"-"`
	CreatedBy               string    `json:"created_by"`
	CreatedTime             time.Time `xml:"-" json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy           string    `json:"last_updated_by"`
	LastUpdatedTime         time.Time `xml:"-" json:"last_updated_time"`
}

// 多字段唯一键
func (u *WorkStep) TableUnique() [][]string {
	return [][]string{
		{"WorkId", "WorkStepName"},
	}
}

func SearchWorkIdsFromWorkStep(search string, o orm.Ormer) []int64 {
	workIds := make([]int64, 0)
	if search != "" {
		lst := make(orm.ParamsList, 0)
		var cond = orm.NewCondition()
		cond = cond.And("work_step_input__icontains", strings.TrimSpace(search)).
			Or("work_step_output__icontains", strings.TrimSpace(search)).
			Or("work_step_type", strings.TrimSpace(search))
		_, err := o.QueryTable("work_step").SetCond(cond).ValuesFlat(&lst, "work_id")
		if err == nil {
			for _, pl := range lst {
				workIds = append(workIds, pl.(int64))
			}
		}
	}
	return workIds
}

func DeleteAllWorkStep(work_id int64, o orm.Ormer) error {
	_, err := o.QueryTable("work_step").Filter("work_id", work_id).Delete()
	return err
}

func InsertOrUpdateWorkStep(step *WorkStep, o orm.Ormer) (id int64, err error) {
	if step.Id > 0 {
		id, err = o.Update(step)
	} else {
		id, err = o.Insert(step)
	}
	return
}

func QueryWorkStep(condArr map[string]interface{}, o orm.Ormer) (steps []WorkStep, err error) {
	qs := o.QueryTable("work_step")
	if work_id, ok := condArr["work_id"]; ok {
		qs = qs.Filter("work_id", work_id)
	}
	qs = qs.OrderBy("work_step_id")
	qs.All(&steps)
	return
}

func QueryOneWorkStep(work_id int64, work_step_id int64, o orm.Ormer) (step WorkStep, err error) {
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).One(&step)
	return
}

func QueryAllWorkStepInfo(work_id int64, o orm.Ormer) (steps []WorkStep, err error) {
	_, err = o.QueryTable("work_step").Filter("work_id", work_id).OrderBy("work_step_id").All(&steps)
	if err == nil {
		for index, step := range steps {
			steps[index].WorkStepInputXml = CDATA{Value: Marshal(step.WorkStepInput, new(iworkmodels.ParamInputSchema))}
			steps[index].WorkStepOutputXml = CDATA{Value: Marshal(step.WorkStepOutput, new(iworkmodels.ParamOutputSchema))}
			steps[index].WorkStepParamMappingXml = CDATA{Value: Marshal(step.WorkStepParamMapping, new(iworkmodels.ParamMapping))}
		}
	}
	return
}

func Marshal(s string, v interface{}) string {

	json.Unmarshal([]byte(s), v)
	bytes, err := xml.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func MarshalJsonToXml(s string) string {
	return ""
}

func QueryWorkStepInfo(work_id int64, work_step_id int64, o orm.Ormer) (step WorkStep, err error) {
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).One(&step)
	return
}

func QueryWorkStepByStepName(work_id int64, work_step_name string, o orm.Ormer) (step WorkStep, err error) {
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_name", work_step_name).One(&step)
	return
}

// mod 只支持 +、- 符号
func BatchChangeWorkStepIdOrder(work_id, work_step_id int64, mod string, o orm.Ormer) error {
	query := fmt.Sprintf("UPDATE work_step SET work_step_id = work_step_id %s 1 WHERE work_id = ? and work_step_id > ?", mod)
	_, err := o.Raw(query, work_id, work_step_id).Exec()
	return err
}

func DeleteWorkStepByWorkStepId(work_id, work_step_id int64, o orm.Ormer) error {
	_, err := o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).Delete()
	if err == nil {
		err = BatchChangeWorkStepIdOrder(work_id, work_step_id, "-", o)
	}
	return err
}

// 获取前置节点信息
func QueryAllPreStepInfo(work_id int64, work_step_id int64, o orm.Ormer) (steps []WorkStep, err error) {
	_, err = o.QueryTable("work_step").Filter("work_id", work_id).
		Filter("work_step_id__lt", work_step_id).OrderBy("work_step_id").All(&steps)
	return
}

func QueryAllWorkStepByWorkName(work_name string, o orm.Ormer) (steps []WorkStep, err error) {
	work, err := QueryWorkByName(work_name, o)
	if err != nil {
		return nil, err
	}
	steps, err = QueryAllWorkStepInfo(work.Id, o)
	if err != nil {
		return nil, err
	}
	return
}

func CopyWorkStepInfo(step WorkStep) *WorkStep {
	newStep := &WorkStep{
		WorkStepName:         step.WorkStepName,
		WorkStepType:         step.WorkStepType,
		WorkStepDesc:         step.WorkStepDesc,
		WorkStepInput:        step.WorkStepInput,
		WorkStepOutput:       step.WorkStepOutput,
		WorkStepParamMapping: step.WorkStepParamMapping,
		CreatedBy:            step.CreatedBy,
		CreatedTime:          step.CreatedTime,
		LastUpdatedBy:        step.LastUpdatedBy,
		LastUpdatedTime:      step.LastUpdatedTime,
	}
	return newStep
}
