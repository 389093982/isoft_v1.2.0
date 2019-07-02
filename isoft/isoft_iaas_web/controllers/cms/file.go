package cms

import (
	"fmt"
	"path"
	"strings"
)

func (this *CMSController) FileUpload() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]string{"status": "ERROR", "errorMsg": fmt.Sprint(err)}
			this.ServeJSON()
		}
	}()
	// 判断是否是文件上传
	f, h, err := this.GetFile("file")
	if err == nil {
		//得到文件的名称
		fileName := h.Filename
		arr := strings.Split(fileName, ":")
		if len(arr) > 1 {
			index := len(arr) - 1
			fileName = arr[index]
		}
		//关闭上传的文件，不然的话会出现临时文件不能清除的情况
		f.Close()
		//保存文件到指定的位置,static/uploadfile,这个是文件的地址,第一个static前面不要有/
		err = this.SaveToFile("file", path.Join("static/uploadfile", fileName))
		if err == nil {
			this.Data["json"] = &map[string]interface{}{
				"status":   "SUCCESS",
				"filename": h.Filename,
				"filepath": "http://localhost:8087/static/uploadfile/" + h.Filename,
			}
		} else {
			this.Data["json"] = &map[string]string{"status": "ERROR", "errorMsg": err.Error()}
		}
	} else {
		this.Data["json"] = &map[string]string{"status": "ERROR", "errorMsg": err.Error()}
	}
	this.ServeJSON()
}
