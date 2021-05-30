package router

import (
	"dev-go-base/biz"
	"dev-go-base/common"
	"dev-go-base/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTableNames(c *gin.Context) {
	d := db.NewDb("r")
	tbl := d.GetAllTableNames()
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    tbl,
	})
}

func GenerateData(c *gin.Context) {
	type Req struct {
		TableName string
		Nums      int
	}
	var req Req

	if err := c.ShouldBindJSON(&req); err != nil {
		common.ErrorHandling(err)
	}
	for i := 0; i < req.Nums; i++ {
		if isCreated := biz.Seeder(req.TableName); !isCreated {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Failure",
				"data":    "",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    "",
	})
}
