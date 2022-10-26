package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// func main() {
// 	// Hello world, the web server

// 	helloHandler := func(w http.ResponseWriter, req *http.Request) {
// 		io.WriteString(w, "Hello, world!\n")
// 	}

// 	http.HandleFunc("/hello", helloHandler)
// 	log.Println("Listing for requests at http://localhost:8000/hello")
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// Intit APP
	app := fiber.New()

	// Loading .env files
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// https://www.mongodb.com/docs/drivers/go/current/quick-start/ => for golang MongoDB tutorial

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Exam api")
	})

	app.Listen(":3000")
}
