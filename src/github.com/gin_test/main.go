package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin_test/controller"
	"github.com/gin_test/system"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(system.LoggerToFile())

	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.POST("/:role/:action", func(c *gin.Context) {
		c.Request.ParseForm()
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(data))
		system.Logger().WithFields(logrus.Fields{
			"name": "yancy",
		}).Info("post请求的输入参数%v", string(data))
		req := controller.Create(string(data))
		system.Logger().WithFields(logrus.Fields{
			"name": "yancy",
		}).Info("After Create %s", req)
		rst := req.BasicController()
		c.String(http.StatusOK, "rst== %v", rst)
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
