package controls

import (
	"encoding/json"
	"net/http"

	controller "github.com/aeckard87/goServe/controllers"
	model "github.com/aeckard87/goServe/models"
	"github.com/gorilla/mux"
)

func CreateDetail(w http.ResponseWriter, r *http.Request) {
	var detail model.Detail
	json.Unmarshal(GetBodyBytes(r), &detail)
	vars := mux.Vars(r)

	//should have an err response too
	detail = controller.CreateDetail(detail, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(detail)
	// fmt.Fprintf(w, StringifyDet(detail))
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	var detail model.Detail
	json.Unmarshal(GetBodyBytes(r), &detail)
	vars := mux.Vars(r)

	//should have an err response too
	detail = controller.GetDetail(detail, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detail)
	// fmt.Fprintf(w, StringifyDet(detail))
}

func GetDetails(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	details := controller.GetDetails()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
	// fmt.Fprintf(w, StringifyDets(details))
}

func UpdateDetail(w http.ResponseWriter, r *http.Request) {
	var detail model.Detail
	json.Unmarshal(GetBodyBytes(r), &detail)
	vars := mux.Vars(r)

	//should have an err response too
	detail = controller.UpdateDetail(detail, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detail)
	// fmt.Fprintf(w, StringifyDet(detail))
}

func DeleteDetail(w http.ResponseWriter, r *http.Request) {
	var detail model.Detail
	json.Unmarshal(GetBodyBytes(r), &detail)
	vars := mux.Vars(r)

	//should have an err response too
	detail = controller.DeleteDetail(detail, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(detail)
	// fmt.Fprintf(w, StringifyDet(detail))
}

func DeleteDetails(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	details := controller.DeleteDetails()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
	// fmt.Fprintf(w, StringifyDets(details))
}
