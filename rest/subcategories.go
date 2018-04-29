package controls

import (
	"encoding/json"
	"net/http"

	controller "github.com/aeckard87/goServe/controllers"
	model "github.com/aeckard87/goServe/models"
	"github.com/gorilla/mux"
)

func CreateSubCategoryByCategory(w http.ResponseWriter, r *http.Request) {
	var subcategory model.SubCategory
	json.Unmarshal(GetBodyBytes(r), &subcategory)
	vars := mux.Vars(r)

	//should have an err response too
	subcategory = controller.CreateSubCategoryByCategory(subcategory, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subcategory)
	// fmt.Fprintf(w, StringifySubCat(subcategory))
}

func GetSubCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	subcategory := controller.GetSubCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategory)
	// fmt.Fprintf(w, StringifySubCat(subcategory))
}

func GetSubCategories(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	subcategories := controller.GetSubCategories()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategories)
	// fmt.Fprintf(w, StringifySubCats(subcategories))
}

func GetSubCategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	subcategories := controller.GetSubCategoriesByCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategories)
	// fmt.Fprintf(w, StringifySubCats(subcategories))
}

func UpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	var subcategory model.SubCategory
	json.Unmarshal(GetBodyBytes(r), &subcategory)
	vars := mux.Vars(r)

	//should have an err response too
	subcategory = controller.UpdateSubCategory(subcategory, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategory)
	// fmt.Fprintf(w, StringifySubCat(subcategory))
}

func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	subcategory := controller.DeleteSubCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategory)
	// fmt.Fprintf(w, StringifySubCat(subcategory))
}

func DeleteSubCategories(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	subcategories := controller.DeleteSubCategories()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategories)
	// fmt.Fprintf(w, StringifySubCats(subcategories))
}

func DeleteSubCategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	subcategories := controller.DeleteSubCategoriesByCategory(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategories)
	// fmt.Fprintf(w, StringifySubCats(subcategories))
}
