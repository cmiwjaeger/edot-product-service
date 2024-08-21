package route

import (
	http "edot-monorepo/services/product-service/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	ProductController *http.ProductController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/product", c.ProductController.List)

}

func (c *RouteConfig) SetupAuthRoute() {
	// c.App.Use(c.AuthMiddleware)
}
