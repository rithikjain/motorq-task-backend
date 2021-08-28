package view

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rithikjain/motorq-task-backend/pkg/utils"
	"net/http"
)

type ErrView struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//noinspection ALL
var (
	ErrMethodNotAllowed = errors.New("Error: Method is not allowed")
	ErrInvalidToken     = errors.New("Error: Invalid Authorization token")
	ErrUserExists       = errors.New("Error: User already exists")
)

var ErrHTTPStatusMap = map[string]int{
	utils.ErrNotFound.Error():     http.StatusNotFound,
	utils.ErrInvalidSlug.Error():  http.StatusBadRequest,
	utils.ErrExists.Error():       http.StatusConflict,
	utils.ErrNoContent.Error():    http.StatusNotFound,
	utils.ErrDatabase.Error():     http.StatusInternalServerError,
	utils.ErrUnauthorized.Error(): http.StatusUnauthorized,
	utils.ErrForbidden.Error():    http.StatusForbidden,
	utils.ErrClash.Error():        http.StatusConflict,
	ErrMethodNotAllowed.Error():   http.StatusMethodNotAllowed,
	ErrInvalidToken.Error():       http.StatusBadRequest,
	ErrUserExists.Error():         http.StatusConflict,
}

func Wrap(err error, c *fiber.Ctx) error {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	if code == 0 {
		code = http.StatusInternalServerError
	}

	errView := ErrView{
		Message: msg,
		Status:  code,
	}

	return c.Status(code).JSON(errView)
}
