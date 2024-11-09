package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Sistema averiado simulado
var systems = map[string]string{
	"navigation":       "NAV-01",
	"communications":   "COM-02",
	"life_support":     "LIFE-03",
	"engines":          "ENG-04",
	"deflector_shield": "SHLD-05",
}

func main() {
	app := fiber.New()

	// Ruta GET /status
	app.Get("/status", func(c *fiber.Ctx) error {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		keys := make([]string, 0, len(systems))
		for key := range systems {
			keys = append(keys, key)
		}
		damagedSystem := keys[r.Intn(len(keys))]

		response := fiber.Map{
			"damaged_system": damagedSystem,
		}
		return c.JSON(response)
	})

	// Ruta GET /repair-bay
	app.Get("/repair-bay", func(c *fiber.Ctx) error {
		damagedSystem := c.Query("system", "navigation") // Obtiene el sistema averiado desde la query
		code, exists := systems[damagedSystem]
		if !exists {
			code = "UNKNOWN"
		}

		htmlContent := `<!DOCTYPE html>
						<html>
						<head>
						    <title>Repair</title>
						</head>
						<body>
						<div class="anchor-point">` + code + `</div>
						</body>
						</html>`

		c.Set("Content-Type", "text/html")
		return c.SendString(htmlContent)
	})

	// Ruta POST /teapot
	app.Post("/teapot", func(c *fiber.Ctx) error {
		return c.Status(http.StatusTeapot).SendString("I'm a teapot")
	})

	// Iniciar servidor en el puerto 3000
	app.Listen(":3000")
}
