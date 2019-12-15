package mail

import (
	"gopkg.in/gomail.v2"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/models"
	"strconv"
	"strings"
)

type SendMailNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *SendMailNode) Execute(trackingId string) {
	mailConn_user := this.TmpDataMap[iworkconst.STRING_PREFIX+"mailConn_user"].(string)
	mailConn_pass := this.TmpDataMap[iworkconst.STRING_PREFIX+"mailConn_pass"].(string)
	mailConn_host := this.TmpDataMap[iworkconst.STRING_PREFIX+"mailConn_host"].(string)
	mailConn_port := this.TmpDataMap[iworkconst.STRING_PREFIX+"mailConn_port"].(string)
	from_label := this.TmpDataMap[iworkconst.STRING_PREFIX+"from_label?"].(string)
	mail_to := this.TmpDataMap[iworkconst.STRING_PREFIX+"mail_to"].(string)
	subject := this.TmpDataMap[iworkconst.STRING_PREFIX+"subject"].(string)
	body := this.TmpDataMap[iworkconst.STRING_PREFIX+"body"].(string)
	mailConn := map[string]string{
		"user": mailConn_user,
		"pass": mailConn_pass,
		"host": mailConn_host,
		"port": mailConn_port,
	}
	mailTo := strings.Split(mail_to, ",")
	err := SendMail(mailConn, from_label, subject, body, mailTo)
	paramMap := make(map[string]interface{}, 0)
	if err != nil {
		paramMap["flag"] = "FAILED"
	} else {
		paramMap["flag"] = "SUCCESS"
	}
	// 将数据数据存储到数据中心
	this.DataStore.CacheDatas(this.WorkStep.WorkStepName, paramMap)
}

func (this *SendMailNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "mailConn_user", "发送邮箱服务器账号信息"},
		2: {iworkconst.STRING_PREFIX + "mailConn_pass", "发送邮箱服务器密码或者授权码"},
		3: {iworkconst.STRING_PREFIX + "mailConn_host", "发送邮箱服务器"},
		4: {iworkconst.STRING_PREFIX + "mailConn_port", "发送端口"},
		5: {iworkconst.STRING_PREFIX + "from_label?", "发送用户别名"},
		6: {iworkconst.STRING_PREFIX + "mail_to", "被发送用户,多个用逗号分隔"},
		7: {iworkconst.STRING_PREFIX + "subject", "邮件主题"},
		8: {iworkconst.STRING_PREFIX + "body", "邮件正文"},
	}
	return this.BPIS1(paramMap)
}

func (this *SendMailNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{"flag"})
}

//定义邮箱服务器连接信息,如果是阿里邮箱 pass填密码,qq邮箱填授权码
//	mailConn := map[string]string {
//		"user": "389093982@qq.com",
//		"pass": "adnngcotnlwucabh",
//		"host": "smtp.qq.com",
//		"port": "465",
//	}
func SendMail(mailConn map[string]string, fromLabel, subject, body string, mailTo []string) error {
	//转换端口类型为int
	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	//这种方式可以添加别名,如 fromLabel,也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("From", fromLabel+"<"+mailConn["user"]+">")
	//发送给多个用户
	m.SetHeader("To", mailTo...)
	//设置邮件主题
	m.SetHeader("Subject", subject)
	//设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
