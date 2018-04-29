package controls

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/aeckard87/wornOut/v1/",
		Index,
	},

	Route{
		"CreateCategory",
		"POST",
		"/aeckard87/wornOut/v1/categories",
		CreateCategory,
	},

	Route{
		"DeleteCategories",
		"DELETE",
		"/aeckard87/wornOut/v1/categories",
		DeleteCategories,
	},

	Route{
		"DeleteCategory",
		"DELETE",
		"/aeckard87/wornOut/v1/categories/{id}",
		DeleteCategory,
	},

	Route{
		"GetCategories",
		"GET",
		"/aeckard87/wornOut/v1/categories",
		GetCategories,
	},

	Route{
		"GetCategory",
		"GET",
		"/aeckard87/wornOut/v1/categories/{id}",
		GetCategory,
	},

	Route{
		"UpdateCategory",
		"PUT",
		"/aeckard87/wornOut/v1/categories/{id}",
		UpdateCategory,
	},

	Route{
		"CreateDescriptorByDetail",
		"POST",
		"/aeckard87/wornOut/v1/details/{id}/descriptors",
		CreateDescriptorByDetail,
	},

	Route{
		"DeleteDescriptor",
		"DELETE",
		"/aeckard87/wornOut/v1/descriptors/{id}",
		DeleteDescriptor,
	},

	Route{
		"DeleteDescriptors",
		"DELETE",
		"/aeckard87/wornOut/v1/descriptors",
		DeleteDescriptors,
	},

	Route{
		"DeleteDescriptorsByDetail",
		"DELETE",
		"/aeckard87/wornOut/v1/details/{id}/descriptors",
		DeleteDescriptorsByDetail,
	},

	Route{
		"GetDescriptor",
		"GET",
		"/aeckard87/wornOut/v1/descriptors/{id}",
		GetDescriptor,
	},

	Route{
		"GetDescriptors",
		"GET",
		"/aeckard87/wornOut/v1/descriptors",
		GetDescriptors,
	},

	Route{
		"GetDescriptorsByDetail",
		"GET",
		"/aeckard87/wornOut/v1/details/{id}/descriptors",
		GetDescriptorsByDetail,
	},

	Route{
		"UpdateDescriptor",
		"PUT",
		"/aeckard87/wornOut/v1/descriptors/{id}",
		UpdateDescriptor,
	},

	Route{
		"CreateDetail",
		"POST",
		"/aeckard87/wornOut/v1/details",
		CreateDetail,
	},

	Route{
		"DeleteDetail",
		"DELETE",
		"/aeckard87/wornOut/v1/details/{id}",
		DeleteDetail,
	},

	Route{
		"DeleteDetails",
		"DELETE",
		"/aeckard87/wornOut/v1/details",
		DeleteDetails,
	},

	Route{
		"GetDetail",
		"GET",
		"/aeckard87/wornOut/v1/details/{id}",
		GetDetail,
	},

	Route{
		"GetDetails",
		"GET",
		"/aeckard87/wornOut/v1/details",
		GetDetails,
	},

	Route{
		"UpdateDetail",
		"PUT",
		"/aeckard87/wornOut/v1/details/{id}",
		UpdateDetail,
	},

	Route{
		"CreateItem",
		"POST",
		"/aeckard87/wornOut/v1/users/{id}/items",
		CreateItem,
	},

	Route{
		"DeleteItem",
		"DELETE",
		"/aeckard87/wornOut/v1/items/{id}",
		DeleteItem,
	},

	Route{
		"DeleteItems",
		"DELETE",
		"/aeckard87/wornOut/v1/items",
		DeleteItems,
	},

	Route{
		"DeleteItemsByOwner",
		"DELETE",
		"/aeckard87/wornOut/v1/users/{id}/items",
		DeleteItemsByOwner,
	},

	Route{
		"DeleteItemsBySubCategory",
		"DELETE",
		"/aeckard87/wornOut/v1/subcategories/{id}/items",
		DeleteItemsBySubCategory,
	},

	Route{
		"GetItem",
		"GET",
		"/aeckard87/wornOut/v1/items/{id}",
		GetItem,
	},

	Route{
		"GetItems",
		"GET",
		"/aeckard87/wornOut/v1/items",
		GetItems,
	},

	Route{
		"GetItemsByOwner",
		"GET",
		"/aeckard87/wornOut/v1/users/{id}/items",
		GetItemsByOwner,
	},

	Route{
		"GetItemsBySubCategory",
		"GET",
		"/aeckard87/wornOut/v1/subcategories/{id}/items",
		GetItemsBySubCategory,
	},

	Route{
		"UpdateItem",
		"PUT",
		"/aeckard87/wornOut/v1/items/{id}",
		UpdateItem,
	},

	Route{
		"CreateSubCategoryByCategory",
		"POST",
		"/aeckard87/wornOut/v1/categories/{id}/subcategories",
		CreateSubCategoryByCategory,
	},

	Route{
		"DeleteSubCategories",
		"DELETE",
		"/aeckard87/wornOut/v1/subcategories",
		DeleteSubCategories,
	},

	Route{
		"DeleteSubCategoriesByCategory",
		"DELETE",
		"/aeckard87/wornOut/v1/categories/{id}/subcategories",
		DeleteSubCategoriesByCategory,
	},

	Route{
		"DeleteSubCategory",
		"DELETE",
		"/aeckard87/wornOut/v1/subcategories/{id}",
		DeleteSubCategory,
	},

	Route{
		"GetSubCategories",
		"GET",
		"/aeckard87/wornOut/v1/subcategories",
		GetSubCategories,
	},

	Route{
		"GetSubCategoriesByCategory",
		"GET",
		"/aeckard87/wornOut/v1/categories/{id}/subcategories",
		GetSubCategoriesByCategory,
	},

	Route{
		"GetSubCategory",
		"GET",
		"/aeckard87/wornOut/v1/subcategories/{id}",
		GetSubCategory,
	},

	Route{
		"UpdateSubCategory",
		"PUT",
		"/aeckard87/wornOut/v1/subcategories/{id}",
		UpdateSubCategory,
	},

	Route{
		"CreateUser",
		"POST",
		"/aeckard87/wornOut/v1/users",
		CreateUser,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/aeckard87/wornOut/v1/users/{id}",
		DeleteUser,
	},

	Route{
		"DeleteUsers",
		"DELETE",
		"/aeckard87/wornOut/v1/users",
		DeleteUsers,
	},

	Route{
		"GetUser",
		"GET",
		"/aeckard87/wornOut/v1/users/{id}",
		GetUser,
	},

	Route{
		"GetUsers",
		"GET",
		"/aeckard87/wornOut/v1/users",
		GetUsers,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/aeckard87/wornOut/v1/users/{id}",
		UpdateUser,
	},
}
