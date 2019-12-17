package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iwork_web/core/iworkutil/errorutil"
	"isoft/isoft_iwork_web/core/iworkutil/fileutil"
	"isoft/isoft_iwork_web/models"
	"isoft/isoft_iwork_web/startup/runtimecfg"
	"path"
)

func (this *WorkController) Import() {
	f, h, _ := this.GetFile("file")
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer f.Close()
	tempFileName := stringutil.RandomUUID() + path.Ext(h.Filename)
	//保存文件到指定的位置
	tempFilePath := path.Join(runtimecfg.FileSavePath, tempFileName)
	err1 := this.SaveToFile("file", tempFilePath)
	err2 := this.ImportFile(tempFilePath)
	if err1 == nil && err2 == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": errorutil.GetFirstError(err1, err2).Error()}
	}
	this.ServeJSON()
}

func (this *WorkController) ImportFile(filepath string) error {
	defer fileutil.RemoveFileOrDirectory(filepath)
	if path.Ext(filepath) == ".placement" {
		return this.ImportPlacementFile(filepath)
	}
	return nil
}

func (this *WorkController) ImportPlacementFile(filepath string) error {
	bytes, err := ioutil.ReadFile(fileutil.ChangeToLinuxSeparator(filepath))
	if err != nil {
		return err
	}
	var pem models.PlacementElementMppaer
	xml.Unmarshal(bytes, &pem)
	// 导入或者更新 Placement
	placement, err := models.QueryPlacementByName(pem.Placement.PlacementName)
	if err == nil {
		pem.Placement.Id = placement.Id
	} else {
		pem.Placement.Id = 0
	}
	_, err = models.InsertOrUpdatePlacement(&pem.Placement)
	if err != nil {
		return err
	}
	// 导入或者更新 Element
	for _, element := range pem.Elements {
		elm, err := models.QueryElementByPlacementAndElementName(placement.PlacementName, element.ElementName)
		if err == nil {
			element.Id = elm.Id
		} else {
			element.Id = 0
		}
		_, err = models.InsertOrUpdateElement(&element)
		if err != nil {
			return err
		}
	}
	return nil
}
