package v1

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/mahfuz110244/project2/entity"
	"github.com/mahfuz110244/project2/lib"

	"github.com/labstack/echo/v4"
)

func GetMessageInfoHandler(c echo.Context) error {
	var b bytes.Buffer
	if _, err := io.Copy(&b, c.Request().Body); err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: "failed to copy response body",
		})

	}
	res, err := lib.GetTextInfo(b)
	if err != nil {
		if strings.Contains(err.Error(), "connect: connection refused") {
			return c.JSON(http.StatusServiceUnavailable, &entity.Response{
				Success: false,
				Message: "Can't get response from Project1",
			})
		}
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: "Something went wrong!",
		})
	}
	return c.JSON(http.StatusOK, res)
}
