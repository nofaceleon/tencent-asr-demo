### 腾讯asr异步实时语音识别接口测试demo

#### 修改配置文件 config.toml,配置好secretId和secretKey

```toml
[tx]
endpoint = "asr.tencentcloudapi.com"
# 腾讯后台的api密钥信息
secretId = ""
secretKey = ""

[web]
# web监听的端口号
port = "8080"
url = "http://127.0.0.1"
# 识别结果保存文件名
filename = "./notify.txt"

```

#### 启动

```shell
nohup ./asr_web_linux_amd64 >> ./log.txt 2>&1 &
```



#### 接口

1. 	POST  /start

描述: 启动语音识别

| 参数        | 必填 | 描述                          |
| ----------- | ---- | ----------------------------- |
| engine_type | 是   | 识别引擎类型,具体参考腾讯文档 |
| url         | 是   | 拉流地址                      |



2. GET /list

描述: 查看当前任务list



3. POST /stop

描述: 停止识别任务

| 参数    | 必填 | 描述   |
| ------- | ---- | ------ |
| task_id | 是   | 任务id |



4. POST /notify

描述: 用于接收腾讯语音识别结果