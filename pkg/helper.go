package pkg

import (
	"encoding/json"
	"strconv"
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
	time := []string{strconv.Itoa(hour), strconv.Itoa(minute), strconv.Itoa(second)}
	timeStr := strings.Join(time, ":")
	return timeStr
}
