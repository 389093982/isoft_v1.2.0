package bash

import (
	"fmt"
	"isoft/isoft_iwork_web/core/iworkconst"
	"isoft/isoft_iwork_web/core/iworklog"
	"isoft/isoft_iwork_web/core/iworkmodels"
	"isoft/isoft_iwork_web/core/iworkplugin/node"
	"isoft/isoft_iwork_web/core/iworkutil/cmdutil"
	"isoft/isoft_iwork_web/models"
	"os"
	"strings"
)

type RunCmdLogWriter struct {
	LogType    string
	TrackingId string
	logwriter  *iworklog.CacheLoggerWriter
}

func (this *RunCmdLogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	messages := strings.Split(message, "\n")
	for _, messageInfo := range messages {
		if strings.TrimSpace(messageInfo) != "" {
			this.logwriter.Write(this.TrackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("%s -- %s", this.LogType, strings.TrimSpace(messageInfo)))
		}
	}
	return len(p), nil
}

type RunCmdNode struct {
	node.BaseNode
	WorkStep *models.WorkStep
}

func (this *RunCmdNode) Execute(trackingId string) {
	if cd := this.TmpDataMap[iworkconst.STRING_PREFIX+"cd?"].(string); cd != "" {
		if err := os.Chdir(cd); err != nil {
			panic(err)
		}
	}

	stdout := &RunCmdLogWriter{
		LogType:    "INFO",
		logwriter:  this.LogWriter,
		TrackingId: trackingId,
	}
	stderr := &RunCmdLogWriter{
		LogType:    "ERROR",
		logwriter:  this.LogWriter,
		TrackingId: trackingId,
	}

	command_name := this.TmpDataMap[iworkconst.STRING_PREFIX+"command_name"].(string)
	command_args := this.TmpDataMap[iworkconst.STRING_PREFIX+"command_args"].(string)
	args := strings.Split(command_args, " ")
	// 记录当前正在执行的命令
	this.LogWriter.Write(trackingId, "", iworkconst.LOG_LEVEL_INFO, fmt.Sprintf("current cmd command is ==> %s %s", command_name, strings.Join(args, " ")))
	err := cmdutil.RunCommand(stdout, stderr, command_name, args...)
	if err != nil {
		panic(err)
	}
}

func (this *RunCmdNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "cd?", "切换目录"},
		2: {iworkconst.STRING_PREFIX + "command_name", "执行命令"},
		3: {iworkconst.STRING_PREFIX + "command_args", "执行命令参数列表"},
	}
	return this.BPIS1(paramMap)
}

func (this *RunCmdNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return this.BPOS1([]string{iworkconst.STRING_PREFIX + "command_result"})
}
