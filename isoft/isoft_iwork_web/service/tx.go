package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

func renderError(err interface{}) error {
	if errorMsg, ok := err.(string); ok {
		return errors.New(errorMsg)
	} else if _, ok := err.(error); ok {
		return err.(error)
	}
	panic(err)
}

func ExecuteWithTx(serviceArgs map[string]interface{}, serviceFunc func(args map[string]interface{}) error) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err == nil {
		defer o.Rollback()
		defer func() {
			if err1 := recover(); err1 != nil {
				err = renderError(err1)
			}
		}()
		serviceArgs["o"] = o
		if err = serviceFunc(serviceArgs); err == nil {
			o.Commit()
		}
	}
	return
}

func ExecuteResultServiceWithTx(serviceArgs map[string]interface{},
	serviceFunc func(args map[string]interface{}) (map[string]interface{}, error)) (result map[string]interface{}, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err == nil {
		defer o.Rollback()
		defer func() {
			if err1 := recover(); err1 != nil {
				err = renderError(err1)
			}
		}()
		serviceArgs["o"] = o
		if result, err = serviceFunc(serviceArgs); err == nil {
			o.Commit()
		}
	}
	return
}
