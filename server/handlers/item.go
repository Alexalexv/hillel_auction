package handlers

import (
	"hillel_auction/server/httpmodels"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateBook godoc
// @Summary      Create item
// @Description  Create new item endpoint
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param        params body httpmodels.CreateItemRequest true "request body"
// @Success      201
// @Failure      400  {object}  httpmodels.Error "[INCORRECT_REQUEST_BODY], [VALIDATION_FAILED]"
// @Router       /item [post]
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
