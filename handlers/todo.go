package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmgchess/gotodo/models"
	"github.com/cmgchess/gotodo/storage"
	"github.com/cmgchess/gotodo/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	store storage.Storage
}

func NewTodoHandler(store storage.Storage) *TodoHandler {
	return &TodoHandler{store: store}
}

func (h *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := h.store.GetTodos()
	utils.JSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}

	todo, err := h.store.GetTodoByID(i)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) AddTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	todo := h.store.AddTodo(todoRequest)
	utils.JSON(w, http.StatusCreated, todo)
}

func (h *TodoHandler) EnableTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	todo, err := h.store.ChangeEnableStatus(i, true)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DisableTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	todo, err := h.store.ChangeEnableStatus(i, false)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
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

	todo, err := h.store.UpdateTodo(i, todoRequest)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	if err := h.store.DeleteTodo(i); err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Todo deleted successfully"))
}
