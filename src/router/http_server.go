package router

import (
	"go-study-server/src/apis"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init_router() http.Handler {
	router := gin.Default()
	application_routers := router.Group("/go-study")
	{
		// 工具模块
		utils_routers := application_routers.Group("/utils")
		{
			utils_routers.GET("/uuid", apis.Get_uuid_64)
			utils_routers.GET("/uuid32", apis.Get_uuid_32)
			utils_routers.Any("/test", test_http_api)
		}
	}

	return router
}

func test_http_api(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "is alive!",
	})
}

func start_http_server(router http.Handler, port string) {
	log.Print("server will started on :", port)
	http.ListenAndServe(":"+port, router)
}

func Start_http_server(port int) {
	// 初始化路由
	router := init_router()
	// 启动服务器
	start_http_server(router, strconv.Itoa(port))
}
