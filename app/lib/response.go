package lib

import (
	"github.com/gofiber/fiber/v2"
)

// Response http response
type Response struct {
	Status  int    `json:"status"`  // http status
	Message string `json:"message"` // response message
}

// Send response
func Send(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Status:  status,
		Message: message,
	})
}

// ErrorBadRequest send http 400 bad request
func ErrorBadRequest(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message = append(message, "Bad request")
	}

	return Send(c, 400, message[0])
}

// ErrorNotFound send http 404 not found
func ErrorNotFound(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message = append(message, "Not found")
	}

	return Send(c, 404, message[0])
}

// ErrorInternal send http 500 internal server error
func ErrorInternal(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message = append(message, "Internal server error")
	}

	return Send(c, 500, message[0])
}

// ErrorConflict send http 409 conflict
func ErrorConflict(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message = append(message, "Conflict")
	}

	return Send(c, 409, message[0])
}

// OK send http 200 response
func OK(c *fiber.Ctx, result ...interface{}) error {
	if len(result) == 0 {
		result = append(result, Response{
			Status:  200,
			Message: "success",
		})
	}

	return c.Status(200).JSON(result[0])
}
