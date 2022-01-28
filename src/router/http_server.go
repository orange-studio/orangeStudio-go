package router

import (
	"go-study-server/src/apis"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initRouter() http.Handler {
	router := gin.Default()
	applicationRouters := router.Group("/orangestudio")
	{
		// 工具模块
		utilsRouters := applicationRouters.Group("/utils")
		{
			utilsRouters.GET("/uuid", apis.GetUUID64)
			utilsRouters.GET("/uuid32", apis.GetUUID32)
			utilsRouters.GET("/test", testHttpApi)
		}
		permissionRouters := applicationRouters.Group("/permission")
		{
			permissionRouters.GET("/getuser/:id", apis.GetUser)
		}
	}

	return router
}

func testHttpApi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "is alive!",
	})
}

func startHttpServer(router http.Handler, port string) {
	log.Print("server will started on :", port)
	http.ListenAndServe(":"+port, router)
}

func StartHttpServer(port int) {
	// 初始化路由
	router := initRouter()
	// 启动服务器
	startHttpServer(router, strconv.Itoa(port))
}
