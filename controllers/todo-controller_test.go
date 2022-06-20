package controllers

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/gavv/httpexpect/v2"
// )

// func irisTester(t *testing.T) *httpexpect.Expect {
// 	handler := IrisHandler()

// 	return httpexpect.WithConfig(httpexpect.Config{
// 		Client: &http.Client{
// 			Transport: httpexpect.NewBinder(handler),
// 			Jar:       httpexpect.NewJar(),
// 		},
// 		Reporter: httpexpect.NewAssertReporter(t),
// 		Printers: []httpexpect.Printer{
// 			httpexpect.NewDebugPrinter(t, true),
// 		},
// 	})
// }

// func TestGetBy(t *testing.T) {
// 	// req := httptest.NewRequest(http.MethodGet, "/todos/62ae5e458acd8d80ad433f97", nil)
// 	// t.Errorf(req.Response.Status)

// 	request := e.GET("/todos/62ae5e458acd8d80ad433f97").
// 		Expect().
// 		Status(http.StatusOK).JSON()
// }
