package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cmgchess/gotodo/models"
	"github.com/cmgchess/gotodo/storage"
	"github.com/cmgchess/gotodo/utils"
	"github.com/go-playground/validator/v10"
)

type TodoHandler struct {
	store storage.Storage
}

func NewTodoHandler(store storage.Storage) *TodoHandler {
	return &TodoHandler{store: store}
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todos, err := h.store.GetTodos(ctx)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	utils.JSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	i, err := utils.ParseIDFromRequest(r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}

	todo, err := h.store.GetTodoByID(ctx, i)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var todoRequest models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todoRequest); err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}
	if err := utils.ValidateStruct(todoRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.Error(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	todo, err := h.store.AddTodo(ctx, todoRequest)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	utils.JSON(w, http.StatusCreated, todo)
}

func (h *TodoHandler) EnableTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	i, err := utils.ParseIDFromRequest(r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	todo, err := h.store.ChangeEnableStatus(ctx, i, true)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DisableTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	i, err := utils.ParseIDFromRequest(r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	todo, err := h.store.ChangeEnableStatus(ctx, i, false)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	i, err := utils.ParseIDFromRequest(r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	var todoRequest models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todoRequest); err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

	if err := utils.ValidateStruct(todoRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.Error(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	todo, err := h.store.UpdateTodo(ctx, i, todoRequest)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	i, err := utils.ParseIDFromRequest(r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	if err := h.store.DeleteTodo(ctx, i); err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
