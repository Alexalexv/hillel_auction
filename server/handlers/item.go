package handlers

import (
	"hillel_auction/server/httpmodels"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateItem godoc
// @Summary      Create item
// @Description  Create new item endpoint
// @Tags         Item
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        params body httpmodels.CreateItemRequest true "request body"
// @Success      201
// @Failure      400  {object}  httpmodels.Error "Message: INCORRECT_REQUEST_BODY, VALIDATION_FAILED"
// @Router       /items [post]
func (h *Handlers) CreateItem(c echo.Context) error {
	h.log.Info(c.Request().Method, c.Request().URL)

	req := new(httpmodels.CreateItemRequest)

	err := c.Bind(req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr(err))
	}

	err = req.Validate()
	if err != nil {
		h.log.Error("validation failed: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.ValidationFailedErr(err))
	}

	return c.NoContent(http.StatusCreated)
}

// GetAllItems godoc
// @Summary      Returns all items
// @Description  Returns all items of current user user
// @Tags         Item
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /items/all [get]
func (h *Handlers) GetAllItems(c echo.Context) error {
	h.log.Info(c.Request().Method, c.Request().URL)

	type all struct {
		All string
	}

	resp := all{
		All: "all",
	}

	return c.JSON(http.StatusOK, resp)
}

// GetItem godoc
// @Summary      Returns item
// @Description  Returns item by id
// @Tags         Item
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200
// @Failure      400  {object}  httpmodels.Error "Message: INCORRECT_PARAMETER"
// @Router       /items/{id} [get]
func (h *Handlers) GetItem(c echo.Context) error {
	h.log.Info(c.Request().Method, c.Request().URL)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("wrong parameter: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectParameter(err))
	}

	resp := httpmodels.GetItemRequest{
		Id:          id,
		Name:        "medal",
		Description: "something old",
		MinBuyPrice: 100,
		SoldPrice:   10000,
		Pictures:    []string{"https://some_domen.ua/medal.jpeg"},
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteItem godoc
// @Summary      Delete item
// @Description  Delete item by id
// @Tags         Item
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200
// @Failure      400  {object}  httpmodels.Error "Message: INCORRECT_PARAMETER"
// @Router       /items/{id} [delete]
func (h *Handlers) DeleteItem(c echo.Context) error {
	h.log.Info(c.Request().Method, c.Request().URL)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("wrong parameter: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectParameter(err))
	}
	h.log.Infof("item %d deleted", id)
	return c.NoContent(http.StatusAccepted)
}

// UpdateItem godoc
// @Summary Update item
// @Description Update item by ID
// @Tags Item
// @Security     ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param body body httpmodels.CreateItemRequest true "request body"
// @Success 200
// @Failure 400 {object} httpmodels.Error "Message: INCORRECT_REQUEST_BODY, VALIDATION_FAILED, INCAORRECT_PARAMETER"
// @Router /items/{id} [put]
func (h *Handlers) UpdateItem(c echo.Context) error {
	h.log.Info(c.Request().Method, c.Request().URL)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("wrong parameter: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectParameter(err))
	}
	req := new(httpmodels.CreateItemRequest)

	err = c.Bind(req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr(err))
	}

	err = req.Validate()
	if err != nil {
		h.log.Error("validation failed: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.ValidationFailedErr(err))
	}
	h.log.Infof("user id: %d updated", id)

	return c.JSON(http.StatusAccepted, req)
}
