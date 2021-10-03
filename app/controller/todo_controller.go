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

// PostTodo godoc
// @Summary Create new todo
// @Description Create new todo
// @Param X-Todo-ID header string false "Todo ID"
// @Param data body model.Todo true "Todo data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Todo data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /todos [post]
// @Tags Todo
func PostTodo(c *fiber.Ctx) error {
	todo := model.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		return lib.ErrorBadRequest(c, err.Error())
	}

	// check required / not null field todo
	validation := todo.Validation("create")
	if len(validation) != 0 {
		return lib.ErrorBadRequest(c, validation)
	}

	db := services.DB
	// Create Data Todo
	if tx := db.Create(&todo); tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate") || strings.Contains(strings.ToLower(tx.Error.Error()), "unique") {
			return lib.ErrorConflict(c, "Duplicate Todo")
		}
		return lib.ErrorInternal(c, tx.Error.Error())
	}

	return lib.OK(c, todo)
}

// GetTodo godoc
// @Summary List of todo features
// @Description List of todo features
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} []model.Todo List of todo features
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /todos [get]
// @Tags Todo
func GetTodo(c *fiber.Ctx) error {
	db := services.DB

	var todos []model.Todo
	db.Model(&model.Todo{}).Find(&todos)

	return lib.OK(c, todos)
}

// GetTodoID godoc
// @Summary Get an todo feature by id
// @Description Get an todo feature by id
// @Param id path string true "Todo ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Todo data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /todos/{id} [get]
// @Tags Todo
func GetTodoID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	todo := model.Todo{}
	result := db.Model(&todo).Where("id = ?", &id).First(&todo)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, todo)
}

// PutTodo godoc
// @Summary Update todo feature by id
// @Description Update todo feature by id
// @Param X-Todo-ID header string false "Todo ID"
// @Param id path string true "Todo ID"
// @Param data body model.Todo true "Todo data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Todo data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /todos/{id} [put]
// @Tags Todo
func PutTodo(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	todo := model.Todo{}
	// check id if exist
	result := db.Where("id = ?", id).First(&todo)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c, fmt.Sprintf("%s %s", result.Error.Error(), id))
	}
	validation := todo.Validation("update")
	if len(validation) != 0 {
		return lib.ErrorBadRequest(c, validation)
	}
	if err := json.Unmarshal(c.Body(), &todo); err != nil {
		return lib.ErrorBadRequest(c, err.Error())
	}
	if tx := db.Updates(&todo); tx.Error != nil {
		return lib.ErrorConflict(c, tx.Error.Error())
	}
	return lib.OK(c, todo)
}

// DeleteTodo godoc
// @Summary TodoDelete todo feature by id
// @Description TodoDelete todo feature by id
// @Param id path string true "Todo ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /todos/{id} [delete]
// @Tags Todo
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	todo := model.Todo{}
	result := db.Model(&todo).Where("id = ?", &id).First(&todo)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&todo)

	return lib.OK(c)
}
