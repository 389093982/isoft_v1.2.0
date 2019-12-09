package mail

import (
	"fmt"
	"testing"
)

func Test_Log(t *testing.T) {
	mailConn := map[string]string{
		"user": "389093982@qq.com",
		"pass": "adnngcotnlwucabh",
		"host": "smtp.qq.com",
		"port": "465",
	}
	//定义收件人
	mailTo := []string{
		"13590331630@163.com",
		"1875112921@qq.com",
	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "领创科技专门测试邮件！！！"
	err := SendMail(mailConn, "lc keji", subject, body, mailTo)
	if err != nil {
		fmt.Println(err.Error())
	}
}
