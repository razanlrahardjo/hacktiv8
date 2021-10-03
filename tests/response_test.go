//go:build !integration
// +build !integration

package lib

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestSend(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return Send(c, 502, "Bad gateway")
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if nil != err {
		t.Error(err)
		return
	}

	utils.AssertEqual(t, 502, response.StatusCode, "Example 502 response")
}

func TestErrorBadRequest(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return ErrorBadRequest(c)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if nil != err {
		t.Error(err)
		return
	}

	utils.AssertEqual(t, 400, response.StatusCode, "Example 400 response")
}

func TestErrorNotFound(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return ErrorNotFound(c)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if nil != err {
		t.Error(err)
		return
	}

	utils.AssertEqual(t, 404, response.StatusCode, "Example 404 response")
}

func TestErrorConflict(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return ErrorConflict(c)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if nil != err {
		t.Error(err)
		return
	}

	utils.AssertEqual(t, 409, response.StatusCode, "Example 409 response")
}

func TestErrorInternal(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return ErrorInternal(c)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if nil != err {
		t.Error(err)
		return
	}

	utils.AssertEqual(t, 500, response.StatusCode, "Example 500 response")
}

func TestOK(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		return OK(c)
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	if nil != err {
		t.Error(err)
		return
	}

	utils.AssertEqual(t, 200, response.StatusCode, "Example 200 response")
}
