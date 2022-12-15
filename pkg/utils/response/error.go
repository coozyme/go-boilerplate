package response

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type BaseErrorResponse struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

func ResponseError(status int, message string, code int64) *BaseErrorResponse {

	return &BaseErrorResponse{
		Data: nil,
		Status: Status{
			Code:    code,
			Message: errors.New(message).Error(),
		},
	}
}

func (s *BaseErrorResponse) Send(c echo.Context) error {
	return c.JSON(int(s.Status.Code), s)
}
