package response

import (
	"encoding/json"

	"github.com/labstack/echo"
)

type Response struct {
	Context    echo.Context
	HTTPStatus int
	Item       interface{}
}

func New(c echo.Context, status int, i interface{}) *Response {
	return &Response{
		Context:    c,
		HTTPStatus: status,
		Item:       i,
	}
}

// Debug returns Response value is encoded json.
func (r *Response) Debug() (result interface{}, err error) {
	debug := struct {
		HTTPStatus int
		Item       interface{}
	}{
		r.HTTPStatus,
		r.Item,
	}

	b, err := json.Marshal(debug)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// JSON sets Respose.Item is encoded json in echo.Context and returns error when happend some exception.
func (r *Response) JSON() (err error) {
	return JSON(r.Context, r.HTTPStatus, r.Item)
}

// XML sets Respose.Item is encoded XML in echo.Context and returns error when happend some exception.
func (r *Response) XML() (err error) {
	return XML(r.Context, r.HTTPStatus, r.Item)
}

// String sets Respose.Item in echo.Context and returns error when happend some exception.
func (r *Response) String() (err error) {
	return String(r.Context, r.HTTPStatus, r.Item)
}

func JSON(c echo.Context, status int, i interface{}) (err error) {
	return c.JSON(status, i)
}

func XML(c echo.Context, status int, i interface{}) (err error) {
	return c.XML(status, i)
}

func String(c echo.Context, status int, i interface{}) (err error) {
	var str string
	switch i.(type) {
	case nil:
		str = ""
	case string:
		str = i.(string)
	default:
		b, err := json.Marshal(i)
		if err != nil {
			return err
		}
		str = string(b)
	}
	return c.String(status, str)
}
