package rest

import (
	"context"
	"fmt"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"utfpr_db/config"
	"utfpr_db/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type healthCheckController interface {
	HealthCheck(c *gin.Context)
}

type userController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type authMiddleware interface {
	Auth() gin.HandlerFunc
}

type Controllers struct {
	UserController        userController
	HealthCheckController healthCheckController
}

type Middlewares struct {
	AuthMiddleware authMiddleware
}

type ServerRest struct {
	httpServer  *http.Server
	engine      *gin.Engine
	config      *config.Config
	controllers *Controllers
	middlewares *Middlewares
}

func NewRestServer(cfg *config.Config, controllers *Controllers, middlewares *Middlewares) *ServerRest {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	docs.SwaggerInfo.Title = "Facade - API"
	docs.SwaggerInfo.Description = "API para teste"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	return &ServerRest{
		engine:      engine,
		config:      cfg,
		controllers: controllers,
		middlewares: middlewares,
	}
}

func (s *ServerRest) registerRoutes() {
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.engine.Use(sentryGin.New(sentryGin.Options{
		WaitForDelivery: true,
	}))

	s.engine.GET("/health-check", s.controllers.HealthCheckController.HealthCheck)

	routeV1 := s.engine.Group("/v1")
	{
		userGroup := routeV1.Group("user")
		{
			userGroup.POST("/", s.controllers.UserController.Create)
			userGroup.GET("/:id", s.controllers.UserController.GetByID)
			userGroup.PUT("/", s.controllers.UserController.Update)
			userGroup.DELETE("/:id", s.controllers.UserController.Delete)
		}
	}
}

func (s *ServerRest) StartListening() {
	s.registerRoutes()

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.RestPort),
		Handler: s.engine,
	}

	fmt.Println("Listening on port", s.config.RestPort)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err.Error())
	}
}

func (s *ServerRest) StopListening(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctxWithTimeout)
	if err != nil {
		fmt.Println("http server forced to shutdown due to timeout")
	}

	fmt.Println("http server was gracefully stopped")
}
