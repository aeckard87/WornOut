package controls

import (
	"encoding/json"
	"net/http"

	controller "github.com/aeckard87/WornOut/controllers"
	model "github.com/aeckard87/WornOut/models"
	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	json.Unmarshal(GetBodyBytes(r), &category)

	//should have an err response too
	category = controller.CreateCategory(category)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//fmt.Fprintf(w, StringifyCat(category))
	json.NewEncoder(w).Encode(category)

}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	json.Unmarshal(GetBodyBytes(r), &category)
	vars := mux.Vars(r)

	//should have an err response too
	category = controller.GetCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
	// fmt.Fprintf(w, StringifyCat(category))
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	categories := controller.GetCategories()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
	// fmt.Fprintf(w, StringifyCats(categories))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	json.Unmarshal(GetBodyBytes(r), &category)
	vars := mux.Vars(r)

	//should have an err response too
	category = controller.UpdateCategory(category, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte)
	json.NewEncoder(w).Encode(category)
	// fmt.Fprintf(w, StringifyCat(category))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	json.Unmarshal(GetBodyBytes(r), &category)
	vars := mux.Vars(r)

	//should have an err response too
	category = controller.DeleteCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
	// fmt.Fprintf(w, StringifyCat(category))
}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	categories := controller.DeleteCategories()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
	// fmt.Fprintf(w, StringifyCats(categories))
}
