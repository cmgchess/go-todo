package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cmgchess/gotodo/models"
	"github.com/cmgchess/gotodo/storage"
	"github.com/cmgchess/gotodo/utils"
	"github.com/gorilla/mux"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := storage.GetTodos()
	utils.JSON(w, http.StatusOK, todos)
}

func GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}

	todo, err := storage.GetTodoByID(i)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, todo)
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoRequest models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todoRequest); err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}
	if todoRequest.Name == "" {
		utils.Error(w, http.StatusBadRequest, errors.New("name is required"))
		return
	}

	todo := storage.AddTodo(todoRequest)
	utils.JSON(w, http.StatusCreated, todo)
}

func EnableTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	todo, err := storage.ChangeEnableStatus(i, true)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func DisableTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	todo, err := storage.ChangeEnableStatus(i, false)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
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
	if todoRequest.Name == "" {
		utils.Error(w, http.StatusBadRequest, errors.New("name is required"))
		return
	}

	todo, err := storage.UpdateTodo(i, todoRequest)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, todo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid ID"))
		return
	}
	if err := storage.DeleteTodo(i); err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return

	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Todo deleted successfully"))
}
