package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viswals/core/interfaces"
	"github.com/viswals/core/models"
	"github.com/viswals/core/pkg/logger"
	"github.com/viswals/core/pkg/utils"
)

const (
	defaultHttpPort = "8080"
)

type IConsumerService interface {
	GetAllUsers(ctx context.Context, paginationParams utils.PaginationParams, filters []utils.Filter) (users []models.User, totalUsers int, err error)
	GetUserById(ctx context.Context, id int64) (user models.User, err error)
}

type Controller struct {
	logger         interfaces.ILogger
	corsMiddleware gin.HandlerFunc
	usecase        IConsumerService
	httpMux        *http.ServeMux
	httpPort       string
}

func (c *Controller) setDefaults() {
	if c.logger == nil {
		logger, err := logger.NewDefaultLogger()
		if err != nil {
			panic(err)
		}

		c.logger = logger
	}

	if c.httpPort == "" {
		c.httpPort = defaultHttpPort
	}
}

type Option func(*Controller)

// WithCorsMiddleware add the cors middleware to existing controllers
func WithCorsMiddleware(corsMiddleware gin.HandlerFunc) func(*Controller) {
	return func(c *Controller) {
		c.corsMiddleware = corsMiddleware
	}
}

func WithHttpMux(httpMux *http.ServeMux) func(*Controller) {
	return func(c *Controller) {
		c.httpMux = httpMux
	}
}

func WithHttpPort(port string) func(*Controller) {
	return func(c *Controller) {
		c.httpPort = port
	}
}

func New(usecase IConsumerService, opts ...Option) *Controller {
	ac := &Controller{
		usecase: usecase,
	}

	for _, opt := range opts {
		opt(ac)
	}

	// set default options
	ac.setDefaults()

	return ac
}

func (c *Controller) Start() error {
	c.registerRoutes()

	server := &http.Server{
		Handler:           c.httpMux,
		Addr:              fmt.Sprintf(":%v", c.httpPort),
		ReadHeaderTimeout: 3 * time.Second,
	}

	return server.ListenAndServe()
}

func (c *Controller) registerRoutes() {
	router := gin.Default()

	// define cors middleware if provided
	if c.corsMiddleware != nil {
		router.Use(c.corsMiddleware)
	}

	routes := router.Group("")
	{
		routes.GET("/users", c.GetAllUsers)
		routes.GET("/users/:id", c.GetUserById)
	}

	c.httpMux.Handle("/", router)
}
