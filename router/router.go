package router

import (
	"dev-go-base/biz"
	"dev-go-base/common"
	"dev-go-base/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func GetTableNames(c *gin.Context) {
	d, err := db.NewDb("s")
	common.ErrorHandling(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Failure",
			"data":    fmt.Sprintf("Origin: %T - %v", errors.Cause(err), errors.Cause(err)),
		})
		return
	}
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

func SetDb(c *gin.Context) {
	type Req struct {
		User, Password, Host, Database string
		Port                           int
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ErrorHandling(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    1000,
			"message": "Failure",
			"data":    "Check your parameters first, please",
		})
		return
	}
	db.NewConFromReq(req.User, req.Password, req.Host, req.Database, req.Port)
	if db.Cons["s"] == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Failure",
			"data":    "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"data":    "",
		})
	}
}
