package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// https://www.yellowduck.be/posts/pretty-print-json-with-go
func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${method} | ${path}\n${resBody}\n\n",
	}))

	app.Get(
		"/api",
		func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		},
	)

	app.Post(
		"/api/*",
		func(c *fiber.Ctx) error {
			body := c.Body()
			prettyJson, err := formatJSON(body)

			if err != nil {
				log.Fatal(err)
			}

			return c.Type("json", "utf-8").Send(prettyJson)
		},
	)

	app.Listen("localhost:3099")
}
