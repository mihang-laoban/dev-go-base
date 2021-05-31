package biz

import (
	"dev-go-base/common"
	"dev-go-base/db"
	"fmt"
)

func Seeder(tblName string) bool {
	d, err := db.NewDb("s")
	common.ErrorHandling(err)
	rows, _ := d.GetAllCols(tblName)

	f := fmt.Sprintf("INSERT INTO %s (", tblName)

	var (
		entries, values string
		rowLen          int = len(rows)
	)

	for i := 0; i < rowLen; i++ {
		if rows[i].ColName == "id" {
			continue
		}
		if i == rowLen-1 {
			entries += fmt.Sprintf("%s", rows[i].ColName)
		} else {
			entries += fmt.Sprintf("%s, ", rows[i].ColName)
		}
		// todo 可以由请求参数修改长度
		lenLimit, _ := common.NumGetter(rows[i].ColType)
		var limit int
		if lenLimit == 0 {
			limit = 5
		}else {
			limit = lenLimit/2
		}

		switch rows[i].DataType {
		case "int", "double":
			tmpNum := common.GetRandomString(limit, common.Digit)
			if i == rowLen-1 {
				values += fmt.Sprintf("'%s'", tmpNum)
			} else {
				values += fmt.Sprintf("'%s', ", tmpNum)
			}
		case "varchar":
			tmpStr := common.GetRandomString(limit, common.LowChar, common.HighChar)
			if i == rowLen-1 {
				values += fmt.Sprintf("'%s'", tmpStr)
			} else {
				values += fmt.Sprintf("'%s', ", tmpStr)
			}
		case "text":
			tmpStr := common.GetRandomString(10, common.LowChar, common.HighChar)
			if i == rowLen-1 {
				values += fmt.Sprintf("'%s'", tmpStr)
			} else {
				values += fmt.Sprintf("'%s', ", tmpStr)
			}
		}
	}
	sql := fmt.Sprintf("%s%s) VALUES (%s)", f, entries, values)
	if err := d.Execute(sql); err != nil {
		return false
	}
	return true
}
