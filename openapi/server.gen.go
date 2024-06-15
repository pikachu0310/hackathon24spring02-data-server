// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
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
	. "github.com/pikachu0310/go-backend-template/openapi/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// OAuth認証コールバック
	// (GET /auth/callback)
	OauthCallback(ctx echo.Context, params OauthCallbackParams) error
	// ログイン
	// (GET /auth/login)
	Login(ctx echo.Context, params LoginParams) error
	// ログアウト
	// (POST /auth/logout)
	Logout(ctx echo.Context) error
	// イベントのリストを取得
	// (GET /events)
	GetEvents(ctx echo.Context) error
	// イベントを登録
	// (POST /events)
	PostEvent(ctx echo.Context) error
	// 開催中のイベントを取得
	// (GET /events/now)
	GetCurrentEvent(ctx echo.Context) error
	// イベントの情報を取得
	// (GET /events/{eventSlug})
	GetEvent(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントの情報を変更
	// (PATCH /events/{eventSlug})
	PatchEvent(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントとイベントに登録されているゲームの情報をCSV形式で取得
	// (GET /events/{eventSlug}/csv)
	GetEventCsv(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントに登録されているゲームのリストを取得
	// (GET /events/{eventSlug}/games)
	GetEventGames(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントの画像を取得
	// (GET /events/{eventSlug}/image)
	GetEventImage(ctx echo.Context, eventSlug EventSlugInPath) error
	// イベントに登録されているタームのリストを取得
	// (GET /events/{eventSlug}/terms)
	GetEventTerms(ctx echo.Context, eventSlug EventSlugInPath) error
	// ゲームのリストを取得 GET /games?termId=X&eventSlug=X&userId=X&includeUnpublished=true
	// (GET /games)
	GetGames(ctx echo.Context, params GetGamesParams) error
	// ゲームを登録
	// (POST /games)
	PostGame(ctx echo.Context) error
	// ゲーム情報を取得
	// (GET /games/{gameId})
	GetGame(ctx echo.Context, gameId GameIdInPath) error
	// ゲーム情報を変更
	// (PATCH /games/{gameId})
	PatchGame(ctx echo.Context, gameId GameIdInPath) error
	// ゲームのアイコン画像を取得
	// (GET /games/{gameId}/icon)
	GetGameIcon(ctx echo.Context, gameId GameIdInPath) error
	// ゲームの画像を取得
	// (GET /games/{gameId}/image)
	GetGameImage(ctx echo.Context, gameId GameIdInPath) error
	// サーバーの生存確認
	// (GET /ping)
	PingServer(ctx echo.Context) error
	// イベントに登録されているタームのリストを取得
	// (GET /terms)
	GetTerms(ctx echo.Context) error
	// タームを登録
	// (POST /terms)
	PostTerm(ctx echo.Context) error
	// ターム情報を取得
	// (GET /terms/{termId})
	GetTerm(ctx echo.Context, termId TermIdInPath) error
	// ターム情報を変更
	// (PATCH /terms/{termId})
	PatchTerm(ctx echo.Context, termId TermIdInPath) error
	// タームに登録されているゲームのリストを取得
	// (GET /terms/{termId}/games)
	GetTermGames(ctx echo.Context, termId TermIdInPath) error
	// テスト用
	// (GET /test)
	Test(ctx echo.Context) error
	// 自分のユーザー情報を取得
	// (GET /users/me)
	GetMe(ctx echo.Context) error
	// 自分が登録したゲームのリストを取得
	// (GET /users/me/games)
	GetMeGames(ctx echo.Context) error
	// ユーザー情報を取得
	// (GET /users/{userId})
	GetUser(ctx echo.Context, userId UserIdInPath) error
	// ユーザーが登録したゲームのリストを取得
	// (GET /users/{userId}/games)
	GetUserGames(ctx echo.Context, userId UserIdInPath) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// OauthCallback converts echo context to params.
func (w *ServerInterfaceWrapper) OauthCallback(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OauthCallbackParams
	// ------------- Required query parameter "code" -------------

	err = runtime.BindQueryParameter("form", true, true, "code", ctx.QueryParams(), &params.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter code: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OauthCallback(ctx, params)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginParams
	// ------------- Required query parameter "redirect" -------------

	err = runtime.BindQueryParameter("form", true, true, "redirect", ctx.QueryParams(), &params.Redirect)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter redirect: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx, params)
	return err
}

// Logout converts echo context to params.
func (w *ServerInterfaceWrapper) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Logout(ctx)
	return err
}

// GetEvents converts echo context to params.
func (w *ServerInterfaceWrapper) GetEvents(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEvents(ctx)
	return err
}

// PostEvent converts echo context to params.
func (w *ServerInterfaceWrapper) PostEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEvent(ctx)
	return err
}

// GetCurrentEvent converts echo context to params.
func (w *ServerInterfaceWrapper) GetCurrentEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCurrentEvent(ctx)
	return err
}

// GetEvent converts echo context to params.
func (w *ServerInterfaceWrapper) GetEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithOptions("simple", "eventSlug", ctx.Param("eventSlug"), &eventSlug, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEvent(ctx, eventSlug)
	return err
}

// PatchEvent converts echo context to params.
func (w *ServerInterfaceWrapper) PatchEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithOptions("simple", "eventSlug", ctx.Param("eventSlug"), &eventSlug, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchEvent(ctx, eventSlug)
	return err
}

// GetEventCsv converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventCsv(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithOptions("simple", "eventSlug", ctx.Param("eventSlug"), &eventSlug, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventCsv(ctx, eventSlug)
	return err
}

// GetEventGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithOptions("simple", "eventSlug", ctx.Param("eventSlug"), &eventSlug, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventGames(ctx, eventSlug)
	return err
}

// GetEventImage converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventImage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithOptions("simple", "eventSlug", ctx.Param("eventSlug"), &eventSlug, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventImage(ctx, eventSlug)
	return err
}

// GetEventTerms converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTerms(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventSlug" -------------
	var eventSlug EventSlugInPath

	err = runtime.BindStyledParameterWithOptions("simple", "eventSlug", ctx.Param("eventSlug"), &eventSlug, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventTerms(ctx, eventSlug)
	return err
}

// GetGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetGames(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetGamesParams
	// ------------- Optional query parameter "termId" -------------

	err = runtime.BindQueryParameter("form", true, false, "termId", ctx.QueryParams(), &params.TermId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// ------------- Optional query parameter "eventSlug" -------------

	err = runtime.BindQueryParameter("form", true, false, "eventSlug", ctx.QueryParams(), &params.EventSlug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventSlug: %s", err))
	}

	// ------------- Optional query parameter "userId" -------------

	err = runtime.BindQueryParameter("form", true, false, "userId", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// ------------- Optional query parameter "includeUnpublished" -------------

	err = runtime.BindQueryParameter("form", true, false, "includeUnpublished", ctx.QueryParams(), &params.IncludeUnpublished)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter includeUnpublished: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGames(ctx, params)
	return err
}

// PostGame converts echo context to params.
func (w *ServerInterfaceWrapper) PostGame(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostGame(ctx)
	return err
}

// GetGame converts echo context to params.
func (w *ServerInterfaceWrapper) GetGame(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "gameId", ctx.Param("gameId"), &gameId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGame(ctx, gameId)
	return err
}

// PatchGame converts echo context to params.
func (w *ServerInterfaceWrapper) PatchGame(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "gameId", ctx.Param("gameId"), &gameId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchGame(ctx, gameId)
	return err
}

// GetGameIcon converts echo context to params.
func (w *ServerInterfaceWrapper) GetGameIcon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "gameId", ctx.Param("gameId"), &gameId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGameIcon(ctx, gameId)
	return err
}

// GetGameImage converts echo context to params.
func (w *ServerInterfaceWrapper) GetGameImage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "gameId" -------------
	var gameId GameIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "gameId", ctx.Param("gameId"), &gameId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter gameId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetGameImage(ctx, gameId)
	return err
}

// PingServer converts echo context to params.
func (w *ServerInterfaceWrapper) PingServer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PingServer(ctx)
	return err
}

// GetTerms converts echo context to params.
func (w *ServerInterfaceWrapper) GetTerms(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTerms(ctx)
	return err
}

// PostTerm converts echo context to params.
func (w *ServerInterfaceWrapper) PostTerm(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTerm(ctx)
	return err
}

// GetTerm converts echo context to params.
func (w *ServerInterfaceWrapper) GetTerm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "termId" -------------
	var termId TermIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "termId", ctx.Param("termId"), &termId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTerm(ctx, termId)
	return err
}

// PatchTerm converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTerm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "termId" -------------
	var termId TermIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "termId", ctx.Param("termId"), &termId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTerm(ctx, termId)
	return err
}

// GetTermGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetTermGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "termId" -------------
	var termId TermIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "termId", ctx.Param("termId"), &termId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter termId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTermGames(ctx, termId)
	return err
}

// Test converts echo context to params.
func (w *ServerInterfaceWrapper) Test(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Test(ctx)
	return err
}

// GetMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMe(ctx)
	return err
}

// GetMeGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetMeGames(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMeGames(ctx)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx, userId)
	return err
}

// GetUserGames converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserGames(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserIdInPath

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserGames(ctx, userId)
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

	router.GET(baseURL+"/auth/callback", wrapper.OauthCallback)
	router.GET(baseURL+"/auth/login", wrapper.Login)
	router.POST(baseURL+"/auth/logout", wrapper.Logout)
	router.GET(baseURL+"/events", wrapper.GetEvents)
	router.POST(baseURL+"/events", wrapper.PostEvent)
	router.GET(baseURL+"/events/now", wrapper.GetCurrentEvent)
	router.GET(baseURL+"/events/:eventSlug", wrapper.GetEvent)
	router.PATCH(baseURL+"/events/:eventSlug", wrapper.PatchEvent)
	router.GET(baseURL+"/events/:eventSlug/csv", wrapper.GetEventCsv)
	router.GET(baseURL+"/events/:eventSlug/games", wrapper.GetEventGames)
	router.GET(baseURL+"/events/:eventSlug/image", wrapper.GetEventImage)
	router.GET(baseURL+"/events/:eventSlug/terms", wrapper.GetEventTerms)
	router.GET(baseURL+"/games", wrapper.GetGames)
	router.POST(baseURL+"/games", wrapper.PostGame)
	router.GET(baseURL+"/games/:gameId", wrapper.GetGame)
	router.PATCH(baseURL+"/games/:gameId", wrapper.PatchGame)
	router.GET(baseURL+"/games/:gameId/icon", wrapper.GetGameIcon)
	router.GET(baseURL+"/games/:gameId/image", wrapper.GetGameImage)
	router.GET(baseURL+"/ping", wrapper.PingServer)
	router.GET(baseURL+"/terms", wrapper.GetTerms)
	router.POST(baseURL+"/terms", wrapper.PostTerm)
	router.GET(baseURL+"/terms/:termId", wrapper.GetTerm)
	router.PATCH(baseURL+"/terms/:termId", wrapper.PatchTerm)
	router.GET(baseURL+"/terms/:termId/games", wrapper.GetTermGames)
	router.GET(baseURL+"/test", wrapper.Test)
	router.GET(baseURL+"/users/me", wrapper.GetMe)
	router.GET(baseURL+"/users/me/games", wrapper.GetMeGames)
	router.GET(baseURL+"/users/:userId", wrapper.GetUser)
	router.GET(baseURL+"/users/:userId/games", wrapper.GetUserGames)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xcf2/bRtL+KgLfF2gLyJHTFIdAQHFIFTcVzokNySpa9AKBljYSW4lkyGUuuUCASSap",
	"UzmXNE3Spi3ObS7XuLbjpE17l2vj6MOsaTnf4rC7JEWKPyVRstvLP4ZFLXdmZ+aZfTg71EWmIjRFgQc8",
	"lJnsRUZkJbYJIJDIJ3AO8LDYUGp5fp6FdXyJ45ksI+IPaYZnm4DJ9kYxaUYCZxVOAlUmCyUFpBm5UgdN",
	"Ft9YBXJF4kTICXgGpN1H+l2kP0H6snkrvCDi2WQocXyNabXSTI1tgnw1VDYdMpDgH5H+DOnf5I8zaeaM",
	"IDVZyGQZReGqvjpAIDUjdKBDBtKhM5AOigykCB3okEF00L/DOmj/Qvozoka/2JZ1L4mEGexjEiCSIAIJ",
	"coBcrrIQhDv3xZ22oW3ufv5P50rxbVOQawKP3DRzfkpgRW6qIlRBDfBT4DyU2CnI1qi8RfNuOz6KymKT",
	"k2VO4OeBxAnVGb4a4nTjh9vd+78gdctor7746kr3Z23nlyu7X6++uPNZcvphtcqyrVdZJIqVAV8N1LoI",
	"WQnG1xsb9UF7QnrLRDcSDxiood7GI1KvOq8gdSt/vPvlJaSudzfaSP12Z/tad3vrtaE0JAoQVHKwERF4",
	"SN3COMNXlpG+kXp153k7m+pubh4+anz19+HkU7EEGz2cfUDVspRK0/gM83Nw5J62tRIWPwQVyLTSzAmC",
	"8H7kVSTAQkE6ZX4ZEDc721/vLt/YW7o81HJNGWWSY/CizQvzbA2UpEYcudgL+pck0zxF6lapMDuSJiJb",
	"A2VFahBtXLIDVUHq1t765u4Xfxsu1zgmJTI5uSJI1RJNtx6px+nXZNWh+TWWbDpZGef2MtdLHtHmT8Lk",
	"JBu47M1VR9tO40g1l8nJ88pig5PrwEemcXnzxZ02Um8jbQWp3yH1EtLaSG0j9XukXkFquyd5URAagOXj",
	"iZbLoi0UKyE22IoPuGguNr75affq0lBrpPP2yMVoBCFW3gJS04qggNRp7TB3sS3tOHIk0CQTJl2GRZvc",
	"mHI7v5dVnQnPjX3fpAkgoSy54rsFIIsCL5M195GcNJNvsjXgHGEbe5HjWemClxilmXkWVupk9gI4qwD5",
	"APAiPEOizChBbmHqFqRSQpRoMgoXbS7E4cDxo9WfEv/eRvom0nWkL3dvrXVv/Wro1526BQRXPMWo6GBC",
	"RklYqTBrXLmM1Ac7253urbV0SuG5swp4LXX4KKx3H14dUngUBbMI1pDTOxKGB9MEd5gNBcJuzKTI1NEp",
	"Zb9YkVsVS/QEWJGVaQbjRdWEeJEl3SVtYrzIkRicFucqUaZG2j2c+7UnSH+SdDKomB4ISkgONcaWhyZM",
	"1yz5Dqlj4mumpIkyNisTUkkjUjZwnm2K+GYmEVIXN0svAKkZmKUBXz0GQ0yI1JXuzxrSriOtvXtXi72t",
	"t9KOUmRUbUD2LTriUD4OzrBKI4x6dO/++mLlx927GlI3jKs/Iu0TY/szHNraVRzjWhsvH6nr+LP6eEFS",
	"gDfE8e6NyUSUJYwHbaQ+H9QSvo4RZPiStb5krS9Z6zhZ6+TqghjPibHhl1Q4MSr8koruJxU9yFzJmRui",
	"y1qmN4Og/9ulWOPlPU4jO0+mLalp0zh+dsU2nZwx4xRQAV8ts3TvHtHwsaRhEWV7s+SqoaveSrrmPxnm",
	"HfMsoGoq0xprxMY69cTSaRT41dGdQd6zYbyAL8lA8ga8KAlnuAYgtXFzF/OAWBJomgW80sSq1Egior0K",
	"TJphq03Ombt6Nyp2icj3K95kK+HAttsd7FtMjdIe5b3LJsHGnxEIRxJ4yFaIZ81OivfnSoVybu7UwrHc",
	"QvnUsZMzWAo2AlOHUJSzmUyNg3Vl8VBFaGbI4BP5hXdKb5VLxZkCHk8vFmbm54r5hbnC+3SSfjpC5RRn",
	"Cu/mczPl4zPFXCE/v5CfO2UTxr4heJbUqwvv5IupfDE1896xk/OzM68xaeYckGQ65eFD04emsSRBBDwr",
	"ckyWOXJo+tARbBUW1olvM6wC65kK22gsspWP8JUaIMvHAcBi3bBzmDk8LGeNSrv6gj4IPPVsE9xt7a1f",
	"21t7RjjGM6Rjzk46Vs4qgGzwpqFxxIc2rPSHwGk8mJ7UkKUcmX7di8gCqHISqMAUFFKCxNU4nm2kKMGq",
	"A7ZqNjbNChXWooXBXToaUtfoX1LAumRWr9RVP8LWInH1xvS0d6a32GrK2qrJmMPeMSUeW1yQuL+a1SxZ",
	"aTYxIcoyc8cUWHfbdAPpN/CDmvYI60GSxQcMnoA5je+lTm4INY53eNjXaSkyO9Ju7mx3kPoPpK4i/SHS",
	"HpMt5Qla0pwfjecrSN1A+jrSl8iVTaQ9InsO5lWlwizSbuIr2hoZ8ynSv0f6t4TndpD6YHflY2PrS7Sk",
	"Mem+cJslqkaEmVeucXnZps9+ISaZwTDeMHNZcrhYczsDPxL1jO4fbGmmCOBUThA+4kB4FP9KAuXfSH9g",
	"EgTsjkdIfzJ0CLui06lqeDAKCmWqguyTcGbp9322f336Da8up4RUTuAh4GFqKmWJv4e070gkfo43Yfx3",
	"FQeav67mYH91yX4qB2ZG62RZ9uo6be0oZq8cK4oNjjo/86HcHwEcBE1y4/9L4AyTZf4v02vFzJitdxna",
	"d9cr57GSxF6gW5jbKLvLN4xPVrFBnFxQX0fafzBQrt8xnn9Ok4+vRWHqbUHh+zNPfz+VNR3Sbpoz9ixo",
	"mu10Kx3gYbv+aKIRyPAtoXqhz2ZNpQE5kZVgBrOnqSoLWbfZwqzlKXH2MSYM/ZbHbYcHclsMb8X0DiWz",
	"THK7Bh50xDvobUFa5KpVwId5V7tpquPj0R4qMrzwlzBk5BRJArzDzSPgYzRD00r1ztOHpObRW+iQSAia",
	"LhQJDrtdtFl6KzKzePdBP1P0hmT627Q9W9hkTe+0z65+2fjmh2TSD50rKvewsFL3ST52z05C1h1P/vJ0",
	"FsVKYFF7pHH/6u5XP1knsObuONG8k5TjyULiQy1Tkc9Fwi0nn5sE4iA4Dy194kWDXw9bGO5yxXeN7XvG",
	"s+uYaieHuzX3xw26T/Sf5ztqu7a7nAoNmiUzNbYJolnYCTJq3xNmLC5HGrmHpnIxDJ8s3RtAXkhOpuch",
	"IX62DwtC/UzKOZPwM1EnI/K1+Dh195DGZn/kBCSZnZHONQQXyUAgNaNRtkBG/TZQRmr4yaOsM2GUdUZA",
	"WWTqDMiaYS1NfuUVu4O756LIl9nCji4CBbnOcQILN+nId9z85rZLyQNMvPv1utXVtuU4K7xp3NhA2pKz",
	"tw0taRxfaShVgNQNs/6lPkDqNaS1jaX7SH30isLbrx28gudTO3jEkorUz0hIrHQvfYvUT8n/j42nT5G6",
	"Yd+A1HXsS3rbXucWUu8GLNPUotQT5rdkuzvpwO6I9vG5Dwj3g6iG7YapEzMLKYrGP1KsvPnen5Xp6df/",
	"YAe0dYFGofXJ66w3oXmWZeLegntYseWEeS4yvlqLs/1kwqUWGjreUMlJgIWWr/enqtLLB56SiidLZy7S",
	"N5hbUel64N3X9fL0WCsCQa7wotb7VLKPkA0uJzjAFVJNSMor46skDAzP32EhwePs/hJCMCQzVjdVGC7z",
	"FYEfOzbH/FAS1h3Ww8dIe2PIvPFcEfWUSHwx1EPigXVGkg6IZ3SRo6vyNfI8x9eKQDpH+jviFLvEBsv1",
	"bSO9FjtR4Gv+vzrhXty8YB53Otf1M1nXDfwXL23VePhF994ve+vXyMjoh1rrefYAPo0al9fIU+DWgX3k",
	"pNYNp55k2WGb2/Bso7/3ccK8kzr0QPLOTjDvtHxmoyNzkT6QtKJgMnA6df1gzlh5Z5ArHMm0E8Q7BwJP",
	"J4owOiARQhiTMucYMNX/ztb/KlvsRLHFYCBFV92wgYc7r0gWU+Mqzdj7xyROKoYRFgZd6lC6n/k6cIH2",
	"t/pZ3q3z3J88XUdXqPzurTVKUBQZe7UZSmhPgnE2UJDG31B3OmqnI6XPvY/XjeUrfW9+B6dSYhrTH5aZ",
	"oqF1EljAGsli7jZoW+oogPHtPg6wuRnJoTXOISy/YiGENMXFBYbXERdpkTKUMpRo+/dg6c31+3ZjpQyT",
	"C/tho90ycnTM47UMt50ka+/fF2icfhsZOnhm8sTs3878kEjaQPpG9+Zj457uetMgm8k0hArbqAsyzB6d",
	"PjqdYUWO8TuL2uzeXveZQM5m6KsI84W546XcQn7uVLlUmCUONzUNfM/m2Hze/RubMuN3dNfxGU031PCD",
	"PtcN5hGm9w7a8u4aS9p0ww/6XOOpJ1qnW/8NAAD//2Uiu/HoVAAA",
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
