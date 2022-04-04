package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// JsonDecode json解析
func JsonDecode(str string) map[string]interface{} {
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return nil
	}
	return m
}

// ResolveTime 时间格式转换
func ResolveTime(seconds int) string {
	seconds = seconds / 1000
	hour := seconds / 3600
	minute := (seconds - hour*3600) / 60
	second := ((seconds - 3600*hour) - 60*minute) % 60
	time := []string{fmt.Sprintf("%02d", hour), fmt.Sprintf("%02d", minute), fmt.Sprintf("%02d", second)}
	timeStr := strings.Join(time, ":")
	return timeStr
}

// WriteFile 数据写入文件
func WriteFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	if _, err := write.WriteString(content); err != nil {
		return err

	}
	//Flush将缓存的文件真正写入到文件中
	if err := write.Flush(); err != nil {
		return err
	}
	return nil
}
