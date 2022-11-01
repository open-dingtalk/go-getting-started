package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/open-dingtalk/go-getting-started/pkg/logger"
	"github.com/open-dingtalk/go-getting-started/pkg/manager"
	"github.com/open-dingtalk/go-getting-started/pkg/oapi"
	"github.com/open-dingtalk/go-getting-started/pkg/service"
	"time"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 6:50 PM
 */

func ActiveInitHttpRouter(port int, staticFilePath string) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery(), MiddlewareLogger)

	coolAppService := service.NewCoolAppService(manager.NewCoolAppManager(oapi.NewDTOAPIClient(nil)))

	//api
	cardGroup := router.Group("/api")
	{
		cardGroup.POST("/sendText", coolAppService.SendTextMessage)
		cardGroup.POST("/sendMessageCard", coolAppService.SendMessageCard)
		cardGroup.POST("/sendTopCard", coolAppService.SendTopCard)
		cardGroup.GET("/getUserInfo", coolAppService.GetUserInfo)
	}

	//健康检查
	router.GET("/status", CheckStatus)

	//静态文件
	router.StaticFile("", staticFilePath+"/index.html")
	router.StaticFile("/index.html", staticFilePath+"/index.html")
	router.StaticFile("/style.css", staticFilePath+"/style.css")
	router.StaticFile("/main.js", staticFilePath+"/main.js")

	return router.Run(fmt.Sprintf(":%d", port))
}

func CheckStatus(c *gin.Context) {
	c.String(200, "ok")
}

func MiddlewareLogger(c *gin.Context) {
	start := time.Now()
	defer func() {
		logger.GetDefaultLogger().Info("http request. url=[%s] rt=[%d] status_code=[%d].",
			c.Request.URL,
			time.Now().Sub(start).Milliseconds(),
			c.Writer.Status())
	}()

	c.Next()
}
