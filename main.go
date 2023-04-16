// https://www.youtube.com/watch?v=E1yiCT2Rdj8
// https://www.youtube.com/watch?v=zlrnwGZMBbU

//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
// npm rum build
//https://docs.gofiber.io/

package main

//import "fmt" // formateara string por consola 
import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port ==""{
		port = "3000"
	}

	app:= fiber.New()

	app.Use(cors.New())

	app.Static("/","./client/dist")

	app.Get("/users",func(c *fiber.Ctx) error{
		return c.JSON(&fiber.Map{
			"data":"User desde el baken",
		})
	})
	app.Listen(":"+port)
	fmt.Println("Server on port 3000")
}
