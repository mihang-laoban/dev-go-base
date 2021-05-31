package test

import (
	"dev-go-base/common"
	"dev-go-base/db"
	"dev-go-base/router"
	server2 "dev-go-base/server"
	"fmt"
	"testing"
)

func TestJd(t *testing.T) {
	for i := 0; i < 10; i++ {
		tmpNum := common.GetRandomString(10, common.Digit)
		fmt.Println(tmpNum)
	}
}

func TestServer(t *testing.T) {
	db.Init()
	server := server2.NewServer()
	server.POST("/getTableNames", router.GetTableNames)
	server.POST("/generateData", router.GenerateData)
	if err := server.Run(":3000"); err != nil {
		common.ErrorHandling(err)
	}
}
