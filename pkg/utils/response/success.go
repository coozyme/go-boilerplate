package response

import (
	"github.com/labstack/echo/v4"
)

type BaseSuccessResponse struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}
type BaseSuccessResponsePagination struct {
	Data   paginationData `json:"data"`
	Status Status         `json:"status"`
}

type paginationData struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta,omitempty"`
}

func ResponsepSuccessPagination(status int, code int64, message string, data interface{}, pagination *Pagination) *BaseSuccessResponsePagination {
	return &BaseSuccessResponsePagination{
		Data: paginationData{
			Data: data,
			Meta: Meta{
				Pagination: pagination,
			},
		},
		Status: Status{
			Message: message,
			Code:    code,
		},
	}
}

func ResponseSuccess(status int, code int64, message string, data interface{}) *BaseSuccessResponse {
	return &BaseSuccessResponse{
		Data: data,
		Status: Status{
			Code:    code,
			Message: message,
		},
	}
}

func (s *BaseSuccessResponsePagination) Send(c echo.Context) error {
	return c.JSON(int(s.Status.Code), s)
}

func (s *BaseSuccessResponse) Send(c echo.Context) error {
	return c.JSON(int(s.Status.Code), s)
}
