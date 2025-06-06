package http

import (
	"api/internal/application"
	"encoding/json"
	"net/http"

	"github.com/jeffotoni/quick"
)

type ItemHandler struct {
	CreateUC *application.CreateItemUseCase
	GetUC    *application.GetItemUseCase
}

func NewItemHandler(createUC *application.CreateItemUseCase, getUC *application.GetItemUseCase) *ItemHandler {
	return &ItemHandler{CreateUC: createUC, GetUC: getUC}
}

func (h *ItemHandler) RegisterRoutes(app *quick.App) {
	app.Post("/items", h.CreateItem)
	app.Get("/items/:id", h.GetItemByID)
}

type createItemRequest struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (h *ItemHandler) CreateItem(c *quick.Ctx) error {
	var req createItemRequest
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "invalid request"})
	}

	item, err := h.CreateUC.Execute(req.Name, req.Value)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(item)
}

func (h *ItemHandler) GetItemByID(c *quick.Ctx) error {
	id := c.Param("id")
	item, err := h.GetUC.Execute(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]string{"error": "item not found"})
	}

	return c.JSON(item)
}
