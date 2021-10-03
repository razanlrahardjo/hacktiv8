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

// PostUser godoc
// @Summary Create new user
// @Description Create new user
// @Param X-User-ID header string false "User ID"
// @Param data body model.User true "User data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users [post]
// @Tags User
func PostUser(c *fiber.Ctx) error {
	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return lib.ErrorBadRequest(c, err.Error())
	}

	// check required / not null field user
	validation := user.Validation("create")
	if len(validation) != 0 {
		return lib.ErrorBadRequest(c, validation)
	}

	db := services.DB
	// Create Data User
	if tx := db.Create(&user); tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate") || strings.Contains(strings.ToLower(tx.Error.Error()), "unique") {
			return lib.ErrorConflict(c, "Duplicate User")
		}
		return lib.ErrorInternal(c, tx.Error.Error())
	}

	return lib.OK(c, user)
}

// GetUser godoc
// @Summary List of user features
// @Description List of user features
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} []model.User List of user features
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users [get]
// @Tags User
func GetUser(c *fiber.Ctx) error {
	db := services.DB

	var users []model.User
	db.Model(&model.User{}).Find(&users)

	return lib.OK(c, users)
}

// PutUser godoc
// @Summary Update user feature by id
// @Description Update user feature by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "User ID"
// @Param data body model.User true "User data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id} [put]
// @Tags User
func PutUser(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	user := model.User{}
	// check id if exist
	result := db.Where("id = ?", id).First(&user)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c, fmt.Sprintf("%s %s", result.Error.Error(), id))
	}
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return lib.ErrorBadRequest(c, err.Error())
	}
	if tx := db.Updates(&user); tx.Error != nil {
		return lib.ErrorConflict(c, tx.Error.Error())
	}
	return lib.OK(c, user)
}

// DeleteUser godoc
// @Summary UserDelete user feature by id
// @Description UserDelete user feature by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id} [delete]
// @Tags User
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	user := model.User{}
	result := db.Model(&user).Where("id = ?", &id).First(&user)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&user)

	return lib.OK(c)
}
