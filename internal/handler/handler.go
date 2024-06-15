package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/pikachu0310/hackathon24spring02-data-server/internal/generate"
	"github.com/pikachu0310/hackathon24spring02-data-server/internal/repository"
	"net/http"
)

type Handler struct {
	repo *repository.Repository
}

func (h *Handler) CreateItem(ctx echo.Context) error {
	item, err := generate.CreateItem()
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) GetItem(ctx echo.Context, itemId openapi_types.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) CombineItems(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetItemWithParameters(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PingServer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "pong")
}

func (h *Handler) Test(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
