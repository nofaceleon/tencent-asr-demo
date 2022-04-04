package controller

import (
	"asr/pkg"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"net/http"
	"os"
)

type IndexController struct {
}

// NotifyData 腾讯回调请求数据
type NotifyData struct {
	TaskId int    `json:"TaskId"`
	Result []*Res `json:"Result"`
}

type Res struct {
	VoiceId   string   `json:"VoiceId"`
	SliceType int      `json:"SliceType"`
	Text      string   `json:"Text"`
	StartTime int      `json:"StartTime"`
	EndTime   int      `json:"EndTime"`
	WorldList []string `json:"WorldList"`
}

var client *asr.Client

func Initialize() {
	client = pkg.GetClient()
}

// Start 开始语音识别
func (i *IndexController) Start(ctx *gin.Context) {
	var (
		EngineType string
		Url        string
	)
	//识别引擎类型
	EngineType = ctx.PostForm("engine_type")
	//拉流地址
	Url = ctx.PostForm("url")

	if len(EngineType) == 0 {
		ctx.JSON(400, gin.H{
			"err": "engine_type 参数不能为空",
		})
		return
	}

	if len(Url) == 0 {
		ctx.JSON(400, gin.H{
			"err": "url(拉流地址) 参数不能为空",
		})
		return
	}

	request := asr.NewCreateAsyncRecognitionTaskRequest()

	request.EngineType = common.StringPtr(EngineType)
	request.Url = common.StringPtr(Url)

	callbackUrl := viper.GetString("web.url") + ":" + viper.GetString("web.port") + "/notify"
	request.CallbackUrl = common.StringPtr(callbackUrl)
	request.AudioData = common.BoolPtr(false)

	response, err := client.CreateAsyncRecognitionTask(request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	fmt.Printf("%s", response.ToJsonString()+"\r\n")
	ctx.JSON(200, pkg.JsonDecode(response.ToJsonString()))
}

// Stop 停止语音识别
func (i *IndexController) Stop(ctx *gin.Context) {
	var TaskId int64
	TaskId = cast.ToInt64(ctx.PostForm("task_id"))
	fmt.Println(TaskId)
	request := asr.NewCloseAsyncRecognitionTaskRequest()
	request.TaskId = common.Int64Ptr(TaskId)
	response, err := client.CloseAsyncRecognitionTask(request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	fmt.Printf("%s", response.ToJsonString()+"\r\n")
	ctx.JSON(200, pkg.JsonDecode(response.ToJsonString()))
	return

}

// List 语音识别列表
func (i *IndexController) List(ctx *gin.Context) {
	request := asr.NewDescribeAsyncRecognitionTasksRequest()
	response, err := client.DescribeAsyncRecognitionTasks(request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	fmt.Printf("%s", response.ToJsonString()+"\r\n")
	ctx.JSON(200, pkg.JsonDecode(response.ToJsonString()))
	return
}

func (i *IndexController) Notify(context *gin.Context) {
	//获取json数据
	var reqInfo *NotifyData
	if err := context.ShouldBindJSON(&reqInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var startTime int
	var text string

	if len(reqInfo.Result) != 0 {
		startTime = reqInfo.Result[0].StartTime
		text = reqInfo.Result[0].Text
	} else {
		context.JSON(200, gin.H{
			"code": 1,
			"msg":  "空数据",
		})
		return
	}
	filePath := viper.GetString("web.filename")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 1,
			"msg":  "文件写入失败",
		})
		return
	}
	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	timeStr := pkg.ResolveTime(startTime)
	_, writeErr := write.WriteString(timeStr + "   " + text + "\r\n")
	if writeErr != nil {
		fmt.Println(writeErr.Error())
		return

	}
	//Flush将缓存的文件真正写入到文件中
	FlushErr := write.Flush()
	if FlushErr != nil {
		fmt.Println(FlushErr.Error())
		return
	}
	context.JSON(200, gin.H{
		"code": 0,
		"msg":  reqInfo.Result[0].Text,
	})
	return
}
