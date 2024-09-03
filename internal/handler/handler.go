package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/pikachu0310/hackathon24spring02-data-server/internal/generate"
	"github.com/pikachu0310/hackathon24spring02-data-server/internal/repository"
	"github.com/pikachu0310/hackathon24spring02-data-server/openapi/models"
	"net/http"
	"sync"
	"time"
)

type Handler struct {
	repo       *repository.Repository
	stockMutex sync.Mutex
	stock      []models.Item
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

// stockItems maintains the stock of items
func (h *Handler) StockItems() {
	for {
		h.stockMutex.Lock()
		if len(h.stock) < 10 {
			item, err := generate.CreateItem()
			if err != nil {
				fmt.Println("Error generating item:", err)
			} else {
				h.stock = append(h.stock, *item)
			}
		}
		h.stockMutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// CreateItem handles the item creation request
func (h *Handler) CreateItem(ctx echo.Context) error {
	h.stockMutex.Lock()
	if len(h.stock) == 0 {
		h.stockMutex.Unlock()
		return ctx.JSON(http.StatusInternalServerError, "No items in stock")
	}

	item := h.stock[0]
	h.stock = h.stock[1:]
	h.stockMutex.Unlock()

	return ctx.JSON(http.StatusOK, item)
}

func (h *Handler) GetItem(ctx echo.Context, itemId openapi_types.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) CombineItems(ctx echo.Context) error {
	combineItemsRequest := new(models.CombineItemsRequest)
	if err := ctx.Bind(combineItemsRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request")
	}

	item, err := generate.CombineItem(combineItemsRequest.Item1Name, combineItemsRequest.Item1Description, combineItemsRequest.Item2Name, combineItemsRequest.Item2Description)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error combining items: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, *item)
}

func (h *Handler) PingServer(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "pong")
}

func (h *Handler) Test(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) MergeItemToMech(ctx echo.Context) error {
	mergeRequest := new(models.MergeRequest)
	if err := ctx.Bind(mergeRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	mech, err := generate.MergeItemToMech(&mergeRequest.Item, &mergeRequest.Mech)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error merging item to mech: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, *mech)
}
