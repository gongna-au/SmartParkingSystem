package router

import (
	"backend/handler/action"
	"backend/handler/login"
	"backend/handler/parking"
	"backend/handler/recommend"
	"backend/handler/user"
	"net/http"
	"strings"

	"backend/handler/signup"

	"backend/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.MaxMultipartMemory = 200 << 20 // 64 MiB

}

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

// 注册全局中间件
func registerGlobalMiddleWare(g *gin.Engine) {
	g.Use(middlewares.Logger())
	g.Use(Cors())
	//g.Use(middlewares.Cors()) //开启中间件 允许使用跨域请求
}

// 设置所有地域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 注册 API 路由
func RegisterAPIRoutes(g *gin.Engine) {
	g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g1 := g.Group("/api/v1/login")
	{
		//使用手机号和密码登陆
		g1.POST("/common", login.CommonUserLoginByPhone)
		g1.POST("/user", login.UpdateUser)
	}
	g2 := g.Group("/api/v1/sign")
	{
		//使用手机号和密码注册
		g2.POST("/common", signup.CommonUserSignupUsingPhone)
	}

	g3 := g.Group("/api/v1/recommend")
	{
		g3.POST("/browse", recommend.GetResourceByBrowse)
		g3.POST("/collect", recommend.GetResourceByCollect)
		g3.POST("/like", recommend.GetResourceByLike)
		g3.POST("/click", recommend.GetResourceByClick)
		g3.POST("/save", recommend.GetResourceBySave)
	}
	g4 := g.Group("/api/v1/action")
	{
		g4.POST("/browse", action.UpdateUserBrowse)
		g4.POST("/collect", action.UpdateUserCollect)
		g4.POST("/like", action.UpdateUserLike)
		g4.POST("/click", action.UpdateUserClick)
		g4.POST("/save", action.UpdateUserSave)
		g4.POST("/search", action.UpdateUserSearch)
	}
	g5 := g.Group("/api/v1/parking")
	{
		g5.POST("/search", parking.SearchParkingLotsByLocation)
		g5.GET("/history", parking.SearchHistoryByUserID)
		g5.GET("/reserve", parking.SearchReserveByUserID)
		g5.GET("/cancel", parking.CancelReserveByID)
	}

	// http://localhost:8083/api/parking/search
	g6 := g.Group("/api/v1/user")
	{
		g6.POST("/card", user.AddNewCard)
		g6.GET("/cards", user.GetBoundedCard)
	}

}

// 配置 404 路由
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
