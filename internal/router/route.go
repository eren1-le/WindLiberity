package router

import (
	"WindLiberity/pkg/app"
	"WindLiberity/pkg/middleware"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter
func NewRouter() *gin.Engine {
	g := gin.New()

	// using middleware
	g.Use(middleware.NoCache)
	g.Use(middleware.Cors)
	g.Use(middleware.Secure)

	g.NoRoute(app.RouteNotFound)
	g.NoMethod(app.RouteNotFound)

	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// HealthCheck 健康检查路由
	g.GET("/health", app.HealthCheck)

	g.StaticFS("/group1",http.Dir("data"))
	if app.IsLocal() {
		g.Use(gin.Logger(), middleware.Logging())
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// pprof router 性能分析路由
		// 默认关闭，开发环境下可以打开
		// 访问方式: HOST/debug/pprof
		// 通过 HOST/debug/pprof/profile 生成profile
		// 查看分析图 go tool pprof -http=:5000 profile
		// see: https://github.com/gin-contrib/pprof
		pprof.Register(g)
	}

	v1 := g.Group("/v1")

	v1.Use(
		middleware.RequestID(),
		middleware.Prom(middleware.WithNamespace("WindLiberity")),
	)
	if app.IsProd() {
		v1.Use(middleware.HandleErrors)
		v1.Use(middleware.VerifySign)
	}
	setApiv1(v1)

	return g
}
