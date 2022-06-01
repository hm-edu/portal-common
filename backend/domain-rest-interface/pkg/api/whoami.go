package api

import (
	"net/http"

	"github.com/hm-edu/portal-common/auth"
	"github.com/labstack/echo/v4"
)

// whoamiHandler godoc
// @Summary whoami Endpoint
// @Description Returns the username of the logged in user
// @Tags User
// @Accept json
// @Produce json
// @Router /whoami [get]
// @Security API
// @Success 200 {string} string "Username"
// @Failure 400 {object} echo.HTTPError "Unauthorized"
func (s *Server) whoamiHandler(c echo.Context) (err error) {
	sub, err := auth.UserFromRequest(c)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Request"}
	}
	return c.JSON(http.StatusOK, sub)
}
