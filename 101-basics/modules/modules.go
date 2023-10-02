package main

import "github.com/gofiber/fiber/v2"

// import'a kendi modülümüzü ekliyoruz fakat herhangi bir verisini kullanmayacaksak, sadece init fonksiyonunu kullanacaksak "_" vereceğiz import'unun başına.
// Kendi paketimizi import ederken import alanına yazdık paket adını ancak fonksiyon içinde kullanırken her defasında paket adını yazmamak için
// connection string başına "." işareti koyarak direkt olarak erişebiliriz

func main() {

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world!")
	})

	_ = app.Listen(":3000")
}
