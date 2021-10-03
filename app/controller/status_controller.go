package controller

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/razanlrahardjo/hacktiv8/app/lib"
	"github.com/razanlrahardjo/hacktiv8/app/model"
	"github.com/razanlrahardjo/hacktiv8/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostStatus godoc
// @Summary Create new status
// @Description Create new status
// @Param X-Status-ID header string false "Status ID"
// @Param data body model.Status true "Status data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Status data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /status [post]
// @Tags Status
func PostStatus(c *fiber.Ctx) error {
	status := model.Status{}
	if err := c.BodyParser(&status); err != nil {
		return lib.ErrorBadRequest(c, err.Error())
	}

	// check required / not null field status
	validation := status.Validation("create")
	if len(validation) != 0 {
		return lib.ErrorBadRequest(c, validation)
	}

	db := services.DB
	// Create Data Status
	if tx := db.Create(&status); tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate") || strings.Contains(strings.ToLower(tx.Error.Error()), "unique") {
			return lib.ErrorConflict(c, "Duplicate Status")
		}
		return lib.ErrorInternal(c, tx.Error.Error())
	}

	return lib.OK(c, status)
}

// GetStatus godoc
// @Summary List of status features
// @Description List of status features
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} []model.Status List of status features
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /status [get]
// @Tags Status
func GetStatus(c *fiber.Ctx) error {
	db := services.DB

	var status []model.Status
	db.Model(&model.Status{}).Find(&status)

	return lib.OK(c, status)
}

// GetStatusID godoc
// @Summary Get an status feature by id
// @Description Get an status feature by id
// @Param id path string true "Status ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Status data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /status/{id} [get]
// @Tags Status
func GetStatusID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	status := model.Status{}
	result := db.Model(&status).Where("id = ?", &id).First(&status)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, status)
}

// PutStatus godoc
// @Summary Update status feature by id
// @Description Update status feature by id
// @Param X-Status-ID header string false "Status ID"
// @Param id path string true "Status ID"
// @Param data body model.Status true "Status data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Status data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /status/{id} [put]
// @Tags Status
func PutStatus(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	status := model.Status{}
	// check id if exist
	result := db.Where("id = ?", id).First(&status)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c, fmt.Sprintf("%s %s", result.Error.Error(), id))
	}
	if err := json.Unmarshal(c.Body(), &status); err != nil {
		return lib.ErrorBadRequest(c, err.Error())
	}
	if tx := db.Updates(&status); tx.Error != nil {
		return lib.ErrorConflict(c, tx.Error.Error())
	}
	return lib.OK(c, status)
}

// DeleteStatus godoc
// @Summary StatusDelete status feature by id
// @Description StatusDelete status feature by id
// @Param id path string true "Status ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /status/{id} [delete]
// @Tags Status
func DeleteStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	status := model.Status{}
	result := db.Model(&status).Where("id = ?", &id).First(&status)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&status)

	return lib.OK(c)
}
