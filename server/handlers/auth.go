package handlers

import (
	"hillel_auction/repository/models"
	"hillel_auction/server/httpmodels"
	"hillel_auction/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Signup godoc
// @Summary      Sign up user
// @Description  Sign up user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        params body httpmodels.SignUpRequest true "request body"
// @Success      201  {object}  httpmodels.SignUpResponse
// @Failure      400  {object}  httpmodels.Error "[INCORRECT_REQUEST_BODY]"
// @Failure      500  {object}  httpmodels.Error "[INTERNAL_SERVER_ERROR]"
// @Router       /signup [post]
func (h *Handlers) SighUp(c echo.Context) error {

	var req httpmodels.SignUpRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr(err))
	}

	user, err := h.userService.CreateUser(&models.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		h.log.Error("failed to create use", err)
		return c.JSON(http.StatusInternalServerError, httpmodels.InternalServerErrorErr(err))
	}

	return c.JSON(http.StatusCreated, httpmodels.SignUpResponse{
		Id:    user.ID,
		Email: user.Email,
	})

}

// Signin godoc
// @Summary      Sign in user
// @Description  Sign in user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        params body httpmodels.SignInRequest true "request body"
// @Success      201  {object}  httpmodels.SignInResponse
// @Failure      400  {object}  httpmodels.Error "[INCORRECT_REQUEST_BODY]"
// @Failure      500  {object}  httpmodels.Error "[INTERNAL_SERVER_ERROR]"
// @Router       /signin [post]
func (h *Handlers) SignIn(c echo.Context) error {
	var req httpmodels.SignInRequest
	err := c.Bind(&req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr(err))
	}

	user, err := h.userService.FindUserByEmailPass(req.Email, req.Password)
	if err != nil {
		h.log.Error("failed to find user", err)
		return c.JSON(http.StatusInternalServerError, httpmodels.InternalServerErrorErr(err))
	}

	if user == nil {
		h.log.Error("no user with such email and password")
		return c.NoContent(http.StatusNoContent)
	}

	jwtPair, err := services.GenerateJWTPairTokens(user.ID)
	if err != nil {
		h.log.Error("failed to create tokens", err)
		return c.JSON(http.StatusInternalServerError, httpmodels.InternalServerErrorErr(err))
	}
	h.log.Info("token created")
	return c.JSON(http.StatusCreated, httpmodels.SignInResponse{
		AccessToken:  jwtPair.AccessToken,
		RefreshToken: jwtPair.RefreshToken,
	})

}

// RefreshJWT godoc
// @Summary      Refresh jwt token
// @Description  Refresh jwt token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        params body httpmodels.RefreshJWTRequest true "request body"
// @Success      201  {object}  httpmodels.RefreshJWTResponse
// @Failure      400  {object}  httpmodels.Error "[INCORRECT_REQUEST_BODY]"
// @Failure      500  {object}  httpmodels.Error "[INTERNAL_SERVER_ERROR]"
// @Router       /refresh [post]
func (h *Handlers) RefreshJWT(c echo.Context) error {
	var req httpmodels.RefreshJWTRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr)
	}

	jwtPair, err := services.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		h.log.Error("failed to create access token", err)
		return c.JSON(http.StatusInternalServerError, httpmodels.InternalServerErrorErr)
	}

	return c.JSON(http.StatusCreated, httpmodels.RefreshJWTResponse{
		AccessToken:  jwtPair.AccessToken,
		RefreshToken: jwtPair.RefreshToken,
	})
}
