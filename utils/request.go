package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}
