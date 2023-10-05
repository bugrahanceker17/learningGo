package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// https://docs.gofiber.io/api/middleware/proxy
	// Make request within handler
	// localhost:3000/user/fkalsf/slşadşdla şeklinde gelecek istekleri key'ine göre yönlendirme yapmak için
	//bir nevi controller mantığı
	app.Get("/:key/*", ProxyHandler)

	// TODO: implement handler!
	app.Delete("cache/:key/*", EvictCacheHandler)

	// TODO: implement handler!
	app.Delete("limit/:key/*", ResetLimitHandler)

	app.Listen(":3000")
}
