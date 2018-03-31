// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	controller "github.com/aeckard87/WornOut/controllers"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	// "github.com/go-swagger/go-swagger/exagccmples/todo-list/models"

	graceful "github.com/tylerb/graceful"

	"github.com/aeckard87/WornOut/restapi/operations"
	"github.com/aeckard87/WornOut/restapi/operations/categories"
	"github.com/aeckard87/WornOut/restapi/operations/descriptors"
	"github.com/aeckard87/WornOut/restapi/operations/details"
	"github.com/aeckard87/WornOut/restapi/operations/items"
	"github.com/aeckard87/WornOut/restapi/operations/subcategories"
	"github.com/aeckard87/WornOut/restapi/operations/users"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yml

func configureFlags(api *operations.WornOutAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WornOutAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//----Categories----
	api.CategoriesCreateCategoryHandler = categories.CreateCategoryHandlerFunc(func(params categories.CreateCategoryParams) middleware.Responder {
		created := controller.CreateCategory(params)
		if created.ID == 0 {
			return categories.NewCreateCategoryBadRequest()
		}
		return categories.NewCreateCategoryCreated().WithPayload(&created)
	})
	api.CategoriesUpdateCategoryHandler = categories.UpdateCategoryHandlerFunc(func(params categories.UpdateCategoryParams) middleware.Responder {
		updated := controller.UpdateCategory(params)
		if updated.Category != params.Body.Category {
			return categories.NewUpdateCategoryBadRequest()
		}
		return categories.NewUpdateCategoryCreated().WithPayload(&updated)
	})
	api.CategoriesDeleteCategoryHandler = categories.DeleteCategoryHandlerFunc(func(params categories.DeleteCategoryParams) middleware.Responder {
		_ = controller.DeleteCategory(params)
		return categories.NewDeleteCategoryOK()
	})
	api.CategoriesDeleteCategoriesHandler = categories.DeleteCategoriesHandlerFunc(func(params categories.DeleteCategoriesParams) middleware.Responder {
		_ = controller.DeleteCategories(params)
		return categories.NewDeleteCategoriesOK()
	})
	api.CategoriesGetCategoryHandler = categories.GetCategoryHandlerFunc(func(params categories.GetCategoryParams) middleware.Responder {
		get := controller.GetCategory(params)
		if get.Category == "" {
			return categories.NewGetCategoryNotFound()
		}
		return categories.NewGetCategoryOK().WithPayload(&get)
	})

	api.CategoriesGetCategoriesHandler = categories.GetCategoriesHandlerFunc(func(params categories.GetCategoriesParams) middleware.Responder {
		getAll := controller.GetCategories(params)
		return categories.NewGetCategoriesOK().WithPayload(getAll)
		// return middleware.NotImplemented("operation categories.GetCategories has not yet been implemented")
	})
	//------------------------

	api.DescriptorsCreateDescriptorHandler = descriptors.CreateDescriptorHandlerFunc(func(params descriptors.CreateDescriptorParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.CreateDescriptor has not yet been implemented")
	})
	api.DescriptorsCreateDescriptorByDetailHandler = descriptors.CreateDescriptorByDetailHandlerFunc(func(params descriptors.CreateDescriptorByDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.CreateDescriptorByDetail has not yet been implemented")
	})
	api.DetailsCreateDetailHandler = details.CreateDetailHandlerFunc(func(params details.CreateDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation details.CreateDetail has not yet been implemented")
	})
	api.ItemsCreateItemHandler = items.CreateItemHandlerFunc(func(params items.CreateItemParams) middleware.Responder {
		return middleware.NotImplemented("operation items.CreateItem has not yet been implemented")
	})
	api.SubcategoriesCreateSubCategoryHandler = subcategories.CreateSubCategoryHandlerFunc(func(params subcategories.CreateSubCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.CreateSubCategory has not yet been implemented")
	})
	api.SubcategoriesCreateSubCategoryByCategoryHandler = subcategories.CreateSubCategoryByCategoryHandlerFunc(func(params subcategories.CreateSubCategoryByCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.CreateSubCategoryByCategory has not yet been implemented")
	})
	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.CreateUser has not yet been implemented")
	})
	api.DescriptorsDeleteDescriptorHandler = descriptors.DeleteDescriptorHandlerFunc(func(params descriptors.DeleteDescriptorParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.DeleteDescriptor has not yet been implemented")
	})
	api.DescriptorsDeleteDescriptorsHandler = descriptors.DeleteDescriptorsHandlerFunc(func(params descriptors.DeleteDescriptorsParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.DeleteDescriptors has not yet been implemented")
	})
	api.DescriptorsDeleteDescriptorsByDetailHandler = descriptors.DeleteDescriptorsByDetailHandlerFunc(func(params descriptors.DeleteDescriptorsByDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.DeleteDescriptorsByDetail has not yet been implemented")
	})
	api.DetailsDeleteDetailHandler = details.DeleteDetailHandlerFunc(func(params details.DeleteDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation details.DeleteDetail has not yet been implemented")
	})
	api.DetailsDeleteDetailsHandler = details.DeleteDetailsHandlerFunc(func(params details.DeleteDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation details.DeleteDetails has not yet been implemented")
	})
	api.ItemsDeleteItemHandler = items.DeleteItemHandlerFunc(func(params items.DeleteItemParams) middleware.Responder {
		return middleware.NotImplemented("operation items.DeleteItem has not yet been implemented")
	})
	api.ItemsDeleteItemsHandler = items.DeleteItemsHandlerFunc(func(params items.DeleteItemsParams) middleware.Responder {
		return middleware.NotImplemented("operation items.DeleteItems has not yet been implemented")
	})
	api.SubcategoriesDeleteSubCategoriesHandler = subcategories.DeleteSubCategoriesHandlerFunc(func(params subcategories.DeleteSubCategoriesParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.DeleteSubCategories has not yet been implemented")
	})
	api.SubcategoriesDeleteSubCategoriesByCategoryHandler = subcategories.DeleteSubCategoriesByCategoryHandlerFunc(func(params subcategories.DeleteSubCategoriesByCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.DeleteSubCategoriesByCategory has not yet been implemented")
	})
	api.SubcategoriesDeleteSubCategoryHandler = subcategories.DeleteSubCategoryHandlerFunc(func(params subcategories.DeleteSubCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.DeleteSubCategory has not yet been implemented")
	})
	api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(func(params users.DeleteUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.DeleteUser has not yet been implemented")
	})
	api.UsersDeleteUsersHandler = users.DeleteUsersHandlerFunc(func(params users.DeleteUsersParams) middleware.Responder {
		return middleware.NotImplemented("operation users.DeleteUsers has not yet been implemented")
	})
	api.DescriptorsGetDescriptorHandler = descriptors.GetDescriptorHandlerFunc(func(params descriptors.GetDescriptorParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.GetDescriptor has not yet been implemented")
	})
	api.DescriptorsGetDescriptorsHandler = descriptors.GetDescriptorsHandlerFunc(func(params descriptors.GetDescriptorsParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.GetDescriptors has not yet been implemented")
	})
	api.DescriptorsGetDescriptorsByDetailHandler = descriptors.GetDescriptorsByDetailHandlerFunc(func(params descriptors.GetDescriptorsByDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.GetDescriptorsByDetail has not yet been implemented")
	})
	api.DetailsGetDetailHandler = details.GetDetailHandlerFunc(func(params details.GetDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation details.GetDetail has not yet been implemented")
	})
	api.DetailsGetDetailsHandler = details.GetDetailsHandlerFunc(func(params details.GetDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation details.GetDetails has not yet been implemented")
	})
	api.ItemsGetItemHandler = items.GetItemHandlerFunc(func(params items.GetItemParams) middleware.Responder {
		return middleware.NotImplemented("operation items.GetItem has not yet been implemented")
	})
	api.ItemsGetItemsHandler = items.GetItemsHandlerFunc(func(params items.GetItemsParams) middleware.Responder {
		return middleware.NotImplemented("operation items.GetItems has not yet been implemented")
	})
	api.ItemsGetItemsByOwnerHandler = items.GetItemsByOwnerHandlerFunc(func(params items.GetItemsByOwnerParams) middleware.Responder {
		return middleware.NotImplemented("operation items.GetItemsByOwner has not yet been implemented")
	})
	api.SubcategoriesGetSubCategoriesHandler = subcategories.GetSubCategoriesHandlerFunc(func(params subcategories.GetSubCategoriesParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.GetSubCategories has not yet been implemented")
	})
	api.SubcategoriesGetSubCategoriesByCategoryHandler = subcategories.GetSubCategoriesByCategoryHandlerFunc(func(params subcategories.GetSubCategoriesByCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.GetSubCategoriesByCategory has not yet been implemented")
	})
	api.SubcategoriesGetSubCategoryHandler = subcategories.GetSubCategoryHandlerFunc(func(params subcategories.GetSubCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.GetSubCategory has not yet been implemented")
	})
	api.UsersGetUserHandler = users.GetUserHandlerFunc(func(params users.GetUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUser has not yet been implemented")
	})
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
	})
	api.DescriptorsUpdateDescriptorHandler = descriptors.UpdateDescriptorHandlerFunc(func(params descriptors.UpdateDescriptorParams) middleware.Responder {
		return middleware.NotImplemented("operation descriptors.UpdateDescriptor has not yet been implemented")
	})
	api.DetailsUpdateDetailHandler = details.UpdateDetailHandlerFunc(func(params details.UpdateDetailParams) middleware.Responder {
		return middleware.NotImplemented("operation details.UpdateDetail has not yet been implemented")
	})
	api.ItemsUpdateItemHandler = items.UpdateItemHandlerFunc(func(params items.UpdateItemParams) middleware.Responder {
		return middleware.NotImplemented("operation items.UpdateItem has not yet been implemented")
	})
	api.SubcategoriesUpdateSubCategoryHandler = subcategories.UpdateSubCategoryHandlerFunc(func(params subcategories.UpdateSubCategoryParams) middleware.Responder {
		return middleware.NotImplemented("operation subcategories.UpdateSubCategory has not yet been implemented")
	})
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.UpdateUser has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
