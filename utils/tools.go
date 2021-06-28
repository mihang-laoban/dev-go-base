package utils

import (
	"bufio"
	"dev-go-base/constants"
	"fmt"
	"os"
	"time"
)

func FileReader(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer DeferHandler(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return err
}

func DeferHandler(file interface{}) {
	switch file.(type) {
	case *os.File:
		err := file.(*os.File).Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetTimeStr(t time.Time) string {
	return t.Format(constants.STD_TIME_FORMAT)
}

// ======= 将时间字符串转换为时间戳 =======
func Time2Unix(t string) int64 {
	stamp, _ := time.ParseInLocation(constants.STD_TIME_FORMAT, t, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp.Unix()
}

func Unix2Time(u int64) string {
	return time.Unix(u, 0).Format(constants.STD_TIME_FORMAT) //输出：2019-01-08 13:50:30
}

func GetNowStr() string {
	return Unix2Time(time.Now().Unix())
}