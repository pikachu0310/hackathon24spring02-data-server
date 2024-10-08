// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// アイテムを取得
	// (GET /items)
	CreateItem(ctx echo.Context) error
	// アイテムを合成
	// (POST /items/combine)
	CombineItems(ctx echo.Context) error
	// アイテムを取得
	// (GET /items/{itemId})
	GetItem(ctx echo.Context, itemId openapi_types.UUID) error
	// アイテムを自機に合成
	// (POST /mech/merge)
	MergeItemToMech(ctx echo.Context) error
	// サーバーの生存確認
	// (GET /ping)
	PingServer(ctx echo.Context) error
	// テスト用
	// (GET /test)
	Test(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateItem converts echo context to params.
func (w *ServerInterfaceWrapper) CreateItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateItem(ctx)
	return err
}

// CombineItems converts echo context to params.
func (w *ServerInterfaceWrapper) CombineItems(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CombineItems(ctx)
	return err
}

// GetItem converts echo context to params.
func (w *ServerInterfaceWrapper) GetItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "itemId" -------------
	var itemId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "itemId", ctx.Param("itemId"), &itemId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter itemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetItem(ctx, itemId)
	return err
}

// MergeItemToMech converts echo context to params.
func (w *ServerInterfaceWrapper) MergeItemToMech(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.MergeItemToMech(ctx)
	return err
}

// PingServer converts echo context to params.
func (w *ServerInterfaceWrapper) PingServer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PingServer(ctx)
	return err
}

// Test converts echo context to params.
func (w *ServerInterfaceWrapper) Test(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Test(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/items", wrapper.CreateItem)
	router.POST(baseURL+"/items/combine", wrapper.CombineItems)
	router.GET(baseURL+"/items/:itemId", wrapper.GetItem)
	router.POST(baseURL+"/mech/merge", wrapper.MergeItemToMech)
	router.GET(baseURL+"/ping", wrapper.PingServer)
	router.GET(baseURL+"/test", wrapper.Test)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RYbVMTyRb+K1Tf+wGrIglo3bLyjYvlvakrQiF+uLVFWc2kSUYnM+NMB2GtWMwMC6yQ",
	"BYGgiGuQZXlJJKKsK4uIP6bpCfyLre4JkJeehHXd3Q9bReVl6D59znmec57TeQAkLaFrKlKxCcIPgCnF",
	"UQLyjx1aol9WUQSjhNmD7iWRidlj3dB0ZGAZ8UUyRonW21FkSoasY1lT2cOKr4DOTLgTM8RaJPak+3yE",
	"TnxPrAKxV4i9Spwx4iwTq3CUe+U+/c5dGAcBgId1BMLAxIasxkAqUDpDhQn0ecbpTJp+m/az3HZ+7w/3",
	"poqZXHGp8BkBtJ0jgDr2/WJIBYCB7iVlA0VB+KvyVAUE2FS4Igq/7/QArf8OkjDznhGgFniIsSH3J7Eg",
	"pGrX37xwR9abmlsf/otYBTfzk5vZvnAWiaxiFEMGO6kuEFVWiZMhzivi7BNnhr+OEXuL2L8QZ0IIQbS+",
	"xchVEAADmpGAGIRBMilHRVbEEJ6bbgY0ZDws4MD+e2Jl6N4anZkg1mgzC8xeoXtrF5qaQw9beda26dYT",
	"UdaqCcD8LsFbifwZYKeOiNDuRFK8Fu1+LalKsopMUxC+84Q7vEqcVYaEVaDT6eLi3uEn281sn7msJhP9",
	"Hs79SUVBuF2RB1GvLCyK/QMWdGGSVfWifbwwV8eMGlN8TRDrZ2JPEXuy+HqULr2tYwVjKN319WT+gzvr",
	"0EdL/gYiKkbGIFT8TNDtUXfWOV6YO34272/lhve9xsbh7ghdesHbwmvi5IjjsA/WOh3fI9Ziycl6yb4p",
	"f+2bJLq6Tqw0sTJ1tusIRf32H49kfTZHoRlvlySkIAOyLVeT3nutpWLOIlbanUsffnzOXHq0zK0WPPTZ",
	"IZn39Q4xDO3+NVlReqCwI9VQtJhdoeltzpAssQ5OedY4HHZSJxy6jtQYjv+mo5g+sVQ3iKVD05Sodl8V",
	"l0atfeKMcD68J86PpwnkPNknTp79114jzo5/HbFDO+HQNc2QkB8y5e771EEUDSDVbOzy8dO39OClnxHf",
	"wOMIKjhOrKkQsfLEyhHrB2Jl3UWbWHk6mSHWY/5wmr3W6xoDhiz5KExNat3Hm+7cmo+rnkMNzcR14Wa9",
	"B8WQ6lNXcZ3Ys6106QU92PRGA87UPLEm2d9JeM3EyRInV1yfvSA6IwHP166PdnLH49NiC0P/PV+UHjnE",
	"serafVFPq3XkeOQZ3d0trn+gkxkOLKMae933qxcDSZqs+DCXTqc5Meo3OFPYGQWi5m+kSoNPCuGUIuWJ",
	"LKN4FQtKgAXKxbaMrid5LHks6BWiVihsWnUac1U3qJLHKpmqlOBKrahV+RqlrNCmSizFk4kRQ3UvIuz9",
	"nwYaAGHwj+DZvSZYutQE+STLaF2aceqt5XNQNbJ8ozc3C1xkq2V1QGOmJU3FUOJ+ekMj0OW7UIonQ5da",
	"QyAAkobC6hxj3QwHgzEZx5P9LZKWCJYtC8ahdBfiuKa2XTZ1NkSG2i5GIYYXTWQMCkZmQJxpLgV5Yn8k",
	"zg6x1pka2JPEzrGvziYn9A6x33JxWC7Ob7DGMmK1R4i1y7Ukx5TD3vAmaWKPCvTs3ZujjQnX+YYuv+Et",
	"d5PYs8XCy+LMmNer2rsjxH53OpqzepExm9FAo4ACYBAZphdLa0uoJcRC1HSkQl0GYXCpJdRyiRUCxHEO",
	"epAhwT/FEM81YwQnciQKwqDDQBDzCyynl6lrqumxpS0UOkEJqXwn1HVFlvje4B3TkwePC+djVaoGDHdi",
	"hj7KNl1sKr8f0OkFevCEhXU5dLm27dzQcNM1LalGOfXMZCIBjeHqK4Y9W7ISABjGzJN7nwn62CYvJ8xR",
	"dnvnhaKZouSUXe+BR3Nk4n9r0eEvlhnRLwipyprCRhKl/kpwSvfvLw8Rt1sXogfsLRJN+fL3PwiXyKtD",
	"AyYQRgYz1fAeK7OnrEhO7oJh4B0FqlMfKEtjg9tvqu9vVEOs0QcTTHD8C4jrEXO8V+v0dOGPqKEK1fuT",
	"i6ekgY2Khx5MsSlyPOduZonzmMmM85K1f/sTa/+/E6eSXStfU1FcjT28dEZRvzLqltXYzROBaZAsjIZw",
	"UFegXJUmNAQTOtcwXVNjguqoSVK3Vvp9rjyyM1VkQjqfpVtPiyt7R7k0XxnEpdFGGEYv+6c4gMqTu/5X",
	"fa4z5ul5cX7Dc9XTW2EzcbY4dHni5Iuz23TFqZhWwsGgoklQiWsmDl8JXQkFmTSnAtVm3OevipmcwAAb",
	"d/7fdavndndP19VbHb2Rrhu3b/VcB6y9lICt19/auyOVTc0UHO5RpmIp50qqL/VrAAAA//9bq39M8BYA",
	"AA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
