package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blue-script/study_go/postgres"
	"github.com/jackc/pgx/v5"
)

type EmployeeHandler struct {
	con *pgx.Conn
}

func NewEmployeeHandler(con *pgx.Conn) *EmployeeHandler {
	return &EmployeeHandler{con: con}
}

func (h *EmployeeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getEmployees(w, r)
	case http.MethodPost:
		h.createEmployee(w, r)
	case http.MethodDelete:
		h.deleteEmployees(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *EmployeeHandler) getEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := postgres.SelectRows(r.Context(), h.con)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Println(employees)

	writeJSON(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) createEmployee(w http.ResponseWriter, r *http.Request) {
	employee := postgres.EmployeeModel{}

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := postgres.InsertRow(r.Context(), h.con, employee); err != nil {
		http.Error(w, "failed to insert", http.StatusInternalServerError)
		return
	}
}

func (h *EmployeeHandler) deleteEmployees(w http.ResponseWriter, r *http.Request) {
	var ids []int

	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := postgres.DeleteEmployees(r.Context(), h.con, ids); err != nil {
		http.Error(w, "failed to delete", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
