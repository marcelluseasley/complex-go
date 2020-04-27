package main

import (
	"database/sql"
	"fmt"
	"log"

	
	_ "github.com/lib/pq"

	"os"
	"strconv"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

// type Fibonacci struct {
// 	num int `json:"number"`
// 	result int `json:"result"`
// }

func main() {
	PGUser := os.Getenv("PGUSER")
	PGHost := os.Getenv("PGHOST")
	PGDatabase := os.Getenv("PGDATABASE")
	PGPassword := os.Getenv("PGPASSWORD")
	PGPort := os.Getenv("PGPORT")

	connStr := fmt.Sprintf("user=%s host=%s dbname=%s password=%s port=%s sslmode=disable",PGUser, PGHost, PGDatabase, PGPassword, PGPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fibo (number INT, factorial INT)")
	if err != nil {
		log.Fatal(err)
	}
	
	// Fiber route handlers

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("hello world")
	})

	app.Get("/values/all", func(c *fiber.Ctx) {
		//get all values from postgress
		// rows, err := db.Query("SELECT * from fibo")
		// if err != nil {
		// 	log.Println(err)
		// }
		c.Send("NOT IMPLEMENTED")
		
	})

	app.Post("/values", func(c *fiber.Ctx) {
		index := c.FormValue("index")
		log.Printf("index is %s", index)
		i, _ := strconv.ParseInt(index, 10, 32)

		if i > 40 {
			c.Status(422).Send("Index too high")
		}
		log.Printf("i is %d", i)
		result := fib(i)
		log.Printf("result is %d", result)
		
		_, err := db.Exec(`INSERT INTO fibo (number, factorial) VALUES ($1, $2)`,i, result)
		if err != nil {
			log.Println(err)
		}

		
		c.Send(result)
	})

	app.Listen(5000)

}

func fib(i int64) int64 {
	if i < 2 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}
