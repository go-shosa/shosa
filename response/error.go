package response

import (
	"fmt"

	"github.com/labstack/echo"
)

type Error struct {
	HTTPStatus       int    `json:"status",xml:"status"`
	Code             int    `json:"code",xml:"code"`
	Property         string `json:"property,omitempty",xml:"property,omitempty"`
	Message          string `json:"message",xml:"message"`
	DeveloperMessage string `json:"developerMessage",xml:"developerMessage"`
	MoreInfo         string `json:"moreInfo,omitempty",xml:"moreInfo,omitempty"`
}

type ErrorResponse struct {
	Context          echo.Context
	ErrorInformation Error
}

func NewErrorResponse(c echo.Context, e Error) *ErrorResponse {
	return &ErrorResponse{
		Context:          c,
		ErrorInformation: e,
	}
}

func (er *ErrorResponse) Error() string {
	return er.ErrorInformation.DeveloperMessage
}

func (er *ErrorResponse) JSON() (err error) {
	ei := er.ErrorInformation
	return er.Context.JSON(ei.HTTPStatus, ei)
}

func (er *ErrorResponse) XML() (err error) {
	ei := er.ErrorInformation
	return er.Context.XML(ei.HTTPStatus, ei)
}

func (er *ErrorResponse) String() (err error) {
	ei := er.ErrorInformation
	str := fmt.Sprintf("status: %d, code: %d, property: %s, message: %s, develop: %s, more: %s", ei.HTTPStatus, ei.Code, ei.Property, ei.Message, ei.DeveloperMessage, ei.MoreInfo)
	return er.Context.String(ei.HTTPStatus, str)
}
