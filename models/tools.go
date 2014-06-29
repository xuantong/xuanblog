package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	DEBUG = true // 输出调试信息
	// DEBUG = false    // 不输出调试信息
	PAGE_SIZE = 10 //每页显示数据
)

func Check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func RedirectToError(this *beego.Controller, err error) {
	if err != nil {
		this.Data["err"] = err
		this.TplNames = "/error.tpl"
		this.Finish()
	}
}

// 当DEBUG开关打开时，输出调试信息
// 使用方法：`debug('格式化字符串', 参数1, 参数2, ...)`
// 输出结果：`DEBUG: 提示信息`(会自动换行)
// 相当于加强版的`fmt.Println()`，支持格式化字符串，输出以`DEBUG: `开头
func Debug(infos ...interface{}) {
	if DEBUG {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

// 把用`, `间隔的字符串转换为字符串列表(例如关键字列表)
func Str2slice(str string) []string {
	return strings.Split(str, ",")
}

// 把字符串转换为time.Time对象
func Str2date(timeStr string) (timeObj time.Time, err error) {
	layout := "2006-01-02"
	timeObj, err = time.Parse(layout, timeStr)
	//Check(err)
	return
}

// 判断字符串切片中是否包含某个字符串
func SliceContains(strSlice []string, str string) bool {
	strMap := make(map[string]bool)
	for _, strItem := range strSlice {
		strMap[strItem] = true
	}
	if _, ok := strMap[str]; ok {
		return true
	}
	return false
}

// convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
