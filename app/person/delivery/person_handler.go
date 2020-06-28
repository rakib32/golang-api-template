package delivery

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"test-api/app/errors"
	"test-api/app/person/usecase"
)

// PersonHandler  represent the httphandler for order
type PersonHandler struct {
	Usecase usecase.PersonUsecase
}

// NewOrderHandler will initialize the orders/ resources endpoint
func NewPersonHandler(e *echo.Echo, us usecase.PersonUsecase) {
	handler := &PersonHandler{
		Usecase: us,
	}

	// API groups
	v1 := e.Group("/api/v1")
	persons := v1.Group("/persons")

	persons.GET("", handler.GetAll)
	persons.GET("/:id", handler.GetByID)
	persons.POST("", handler.Create)
	persons.DELETE("/:id", handler.Remove)
	persons.PUT("/:id", handler.Update)
}

// GetAll will fetch all
func (o *PersonHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	lists, err := o.Usecase.GetAll()
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, lists)
}

// GetByID will get order by given id
func (o *PersonHandler) GetByID(c echo.Context) error {
	personID, _ := strconv.Atoi(c.Param("id"))

	ord, err := o.Usecase.GetById(personID)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, ord)
}

// Create will create new person row
func (o *PersonHandler) Create(c echo.Context) error {
	req := CreatePersonRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.RespondError(err))
	}

	person, err := o.Usecase.Create(req.FirstName, req.LastName)

	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, person)
}

// Update will update the data
func (o *PersonHandler) Update(c echo.Context) error {
	req := CreatePersonRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.RespondError(err))
	}

	personID, _ := strconv.Atoi(c.Param("id"))
	person, err := o.Usecase.Update(personID, req.FirstName, req.LastName)

	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, person)
}

// Update will update the data
func (o *PersonHandler) Remove(c echo.Context) error {

	personID, _ := strconv.Atoi(c.Param("id"))
	err := o.Usecase.Remove(personID)

	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "updated",
		"error":   "",
	})
}

// CreatePersonRequest ...
type CreatePersonRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
