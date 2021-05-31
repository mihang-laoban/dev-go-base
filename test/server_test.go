package test

import (
	"dev-go-base/common"
	"dev-go-base/router"
	server2 "dev-go-base/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestJd(t *testing.T) {
	for i := 0; i < 10; i++ {
		tmpNum := common.GetRandomString(10, common.Digit)
		fmt.Println(tmpNum)
	}
}

func TestServer(t *testing.T) {
	//db.Init()
	server := server2.NewServer()
	server.POST("/getTableNames", router.GetTableNames)
	server.POST("/generateData", router.GenerateData)
	server.POST("/setDb", router.SetDb)
	server.GET("/g", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "world",
		})
	})
	if err := server.Run(":8090"); err != nil {
		common.ErrorHandling(err)
	}
}
