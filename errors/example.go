package errors

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
//     app := fiber.New(fiber.Config{
//         ErrorHandler: ErrorHandle,
//     })

//     // SetError will not display an exit message but will take it to the log
//     // SetErrorMessage will display a message in the response
//     // SetErrorWithData will return message and data in response

//     app.Get("/error", func(c *fiber.Ctx) error {
//         return SetError(http.StatusBadRequest, "Error Message")
//     })
//     app.Get("/error-message", func (c *fiber.Ctx) error {
//         return SetErrorMessage(http.StatusBadRequest, "Error Message")
//     })
//     app.Get("/error-message-data", func(c *fiber.Ctx) error {
//         return SetErrorMessageWithData(http.StatusNotFound, "Error Message", "Not a data")
//     })

//     log.Fatal(app.Listen(":3000"))
// }