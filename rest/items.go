package controls

import (
	"encoding/json"
	"net/http"

	controller "github.com/aeckard87/goServe/controllers"
	model "github.com/aeckard87/goServe/models"
	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	json.Unmarshal(GetBodyBytes(r), &item)

	//should have an err response too
	item = controller.CreateItem(item)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
	// fmt.Fprintf(w, StringifyItem(item))
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	json.Unmarshal(GetBodyBytes(r), &item)
	vars := mux.Vars(r)

	//should have an err response too
	item = controller.GetItem(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
	// fmt.Fprintf(w, StringifyItem(item))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	items := controller.GetItems()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, StringifyItems(items))
}

func GetItemsByOwner(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	vars := mux.Vars(r)
	items := controller.GetItemsByOwner(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, StringifyItems(items))
}

func GetItemsBySubCategory(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	vars := mux.Vars(r)
	items := controller.GetItemsBySubCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, StringifyItems(items))
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	json.Unmarshal(GetBodyBytes(r), &item)
	vars := mux.Vars(r)

	//should have an err response too
	item = controller.UpdateItem(item, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
	// fmt.Fprintf(w, StringifyItem(item))
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	json.Unmarshal(GetBodyBytes(r), &item)
	vars := mux.Vars(r)

	//should have an err response too
	item = controller.DeleteItem(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
	// fmt.Fprintf(w, StringifyItem(item))
}

func DeleteItems(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	items := controller.DeleteItems()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, StringifyItems(items))
}

func DeleteItemsByOwner(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	vars := mux.Vars(r)
	items := controller.DeleteItemsByOwner(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, StringifyItems(items))
}

func DeleteItemsBySubCategory(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	vars := mux.Vars(r)
	items := controller.DeleteItemsBySubCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, StringifyItems(items))
}
