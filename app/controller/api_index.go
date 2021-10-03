package controller

import (
	"github.com/razanlrahardjo/hacktiv8/app/lib"

	"github.com/gofiber/fiber/v2"
)

// ApiIndexGet index page
// @Summary show basic response
// @Description show basic response
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Failure 409 {object} lib.Response "conflict"
// @Failure 500 {object} lib.Response "internal error"
// @Router / [get]
// @Tags Index
func ApiIndexGet(c *fiber.Ctx) error {
	return lib.OK(c)
}
