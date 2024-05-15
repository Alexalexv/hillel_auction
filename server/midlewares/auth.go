package midlewares

import (
	"hillel_auction/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Authorization")

		_, err := services.ParseAccessToken(accessToken)
		if err != nil {
			log.Error("failed to parse access token")
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
