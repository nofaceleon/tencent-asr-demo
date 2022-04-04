package main

import (
	"asr/app/http/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//配置初始化
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("web.filename", "./notify.txt")
	viper.SetDefault("web.file_row", "./notify_row.txt")
	viper.SetDefault("web.port", "8080")
	configErr := viper.ReadInConfig() //读取配置
	if configErr != nil {
		panic(configErr.Error())
	}
	viper.WatchConfig()
}

//交叉编译
//GOOS=linux GOARCH=amd64 go build -o asr_web_linux_amd64 main.go

func main() {
	//gin.SetMode(gin.ReleaseMode) //设置运行模式
	//初始化配置信息
	r := gin.Default()
	controller.Initialize()
	suc := new(controller.IndexController)
	r.POST("/notify", suc.Notify) //接收腾讯回调
	r.GET("/list", suc.List)      //查询当前进行的任务
	r.POST("/start", suc.Start)
	r.POST("/stop", suc.Stop)
	err := r.Run(":" + viper.GetString("web.port"))
	if err != nil {
		fmt.Println(err.Error())
	} //默认8080端口

}
