package response

import (
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
)

type (
	user struct {
		ID   int    `json:"id" xml:"id" form:"id"`
		Name string `json:"name" xml:"name" form:"name"`
	}
)

const (
	userJSON       = `{"id":1,"name":"Jon Snow"}`
	userXML        = `<user><id>1</id><name>Jon Snow</name></user>`
	userForm       = `id=1&name=Jon Snow`
	invalidContent = "invalid content"
)

func TestNew(t *testing.T) {
	e := echo.New()
	req := test.NewRequest("GET", "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	status := 200
	item := "hello, human!!"
	expected := &Response{
		Context:    c,
		HTTPStatus: status,
		Item:       item,
	}

	actual := New(c, status, item)
	if expected.HTTPStatus != actual.HTTPStatus {
		t.Fatalf("actual.HTTPStatus should be equal expected.HTTPStatus. expected:%d, actual:%d", expected.HTTPStatus, actual.HTTPStatus)
	}
	if expected.Item != actual.Item {
		t.Fatalf("actual.Item should be equal expected.Item. expected:%v, actual:%v", expected.Item, actual.Item)
	}
}

func TestJSON(t *testing.T) {
	e := echo.New()
	req := test.NewRequest("GET", "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	resp := New(c, http.StatusOK, user{1, "Jon Snow"})
	err := resp.JSON()
	if err != nil {
		t.Fatalf("Recieved unexpected error :\n%+v", err)
	}
	if http.StatusOK != rec.Status() {
		t.Errorf("Response http status is expected %d, but actual %d", http.StatusOK, rec.Status())
	}
	if echo.MIMEApplicationJSONCharsetUTF8 != rec.Header().Get(echo.HeaderContentType) {
		t.Errorf("Response header is expected %d, but actual %d", echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
	if userJSON != rec.Body.String() {
		t.Errorf("Response body is expected %d, but actual %d", userJSON, rec.Body.String())
	}
}

func TestXML(t *testing.T) {
	actual := xml.Header + userXML

	e := echo.New()
	req := test.NewRequest("GET", "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	resp := New(c, http.StatusOK, user{1, "Jon Snow"})
	err := resp.XML()
	if err != nil {
		t.Fatalf("Recieved unexpected error :\n%+v", err)
	}
	if http.StatusOK != rec.Status() {
		t.Errorf("Response http status is expected %d, but actual %d", http.StatusOK, rec.Status())
	}
	if echo.MIMEApplicationXMLCharsetUTF8 != rec.Header().Get(echo.HeaderContentType) {
		t.Errorf("Response header is expected %d, but actual %d", echo.MIMEApplicationXMLCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
	if actual != rec.Body.String() {
		t.Errorf("Response body is expected %d, but actual %d", actual, rec.Body.String())
	}
}

func TestString(t *testing.T) {
	e := echo.New()
	req := test.NewRequest("GET", "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	resp := New(c, http.StatusOK, "Hello, World!")
	err := resp.String()
	if err != nil {
		t.Fatalf("Recieved unexpected error :\n%+v", err)
	}
	if http.StatusOK != rec.Status() {
		t.Errorf("Response http status is expected %d, but actual %d", http.StatusOK, rec.Status())
	}
	if echo.MIMETextPlainCharsetUTF8 != rec.Header().Get(echo.HeaderContentType) {
		t.Errorf("Response header is expected %d, but actual %d", echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
	if "Hello, World!" != rec.Body.String() {
		t.Errorf("Response body is expected %d, but actual %d", "Hello, World!", rec.Body.String())
	}
}
