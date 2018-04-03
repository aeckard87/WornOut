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
	//------------------------i

	//----SubCategories----
	api.SubcategoriesCreateSubCategoryByCategoryHandler = subcategories.CreateSubCategoryByCategoryHandlerFunc(func(params subcategories.CreateSubCategoryByCategoryParams) middleware.Responder {
		created := controller.CreateSubCategoryByCategory(params)
		if created.ID == 0 {
			return subcategories.NewCreateSubCategoryByCategoryBadRequest()
		}
		return subcategories.NewCreateSubCategoryByCategoryCreated().WithPayload(&created)
	})
	api.SubcategoriesUpdateSubCategoryHandler = subcategories.UpdateSubCategoryHandlerFunc(func(params subcategories.UpdateSubCategoryParams) middleware.Responder {
		updated := controller.UpdateSubCategory(params)
		if updated.Subcategory != params.Body.Subcategory {
			return subcategories.NewUpdateSubCategoryNotFound()
		}
		return subcategories.NewUpdateSubCategoryCreated().WithPayload(&updated)
	})
	api.SubcategoriesDeleteSubCategoriesHandler = subcategories.DeleteSubCategoriesHandlerFunc(func(params subcategories.DeleteSubCategoriesParams) middleware.Responder {
		_ = controller.DeleteSubCategories(params)
		return subcategories.NewDeleteSubCategoriesOK()
	})

	api.SubcategoriesDeleteSubCategoriesByCategoryHandler = subcategories.DeleteSubCategoriesByCategoryHandlerFunc(func(params subcategories.DeleteSubCategoriesByCategoryParams) middleware.Responder {
		_ = controller.DeleteSubcategoriesByCategory(params)
		return subcategories.NewDeleteSubCategoryOK()
	})
	api.SubcategoriesDeleteSubCategoryHandler = subcategories.DeleteSubCategoryHandlerFunc(func(params subcategories.DeleteSubCategoryParams) middleware.Responder {
		_ = controller.DeleteSubCategory(params)
		return subcategories.NewDeleteSubCategoryOK()
	})
	api.SubcategoriesGetSubCategoriesHandler = subcategories.GetSubCategoriesHandlerFunc(func(params subcategories.GetSubCategoriesParams) middleware.Responder {
		fetched := controller.GetSubcategories(params)
		return subcategories.NewGetSubCategoriesOK().WithPayload(fetched)
	})
	api.SubcategoriesGetSubCategoriesByCategoryHandler = subcategories.GetSubCategoriesByCategoryHandlerFunc(func(params subcategories.GetSubCategoriesByCategoryParams) middleware.Responder {
		fetched := controller.GetSubcategoriesByCategory(params)
		return subcategories.NewGetSubCategoriesByCategoryOK().WithPayload(fetched)
	})
	api.SubcategoriesGetSubCategoryHandler = subcategories.GetSubCategoryHandlerFunc(func(params subcategories.GetSubCategoryParams) middleware.Responder {
		fetched := controller.GetSubCategory(params)
		return subcategories.NewGetSubCategoryOK().WithPayload(&fetched)
	})

	//-----

	//----DETAILS----
	api.DetailsCreateDetailHandler = details.CreateDetailHandlerFunc(func(params details.CreateDetailParams) middleware.Responder {
		created := controller.CreateDetail(params)
		if created.ID == 0 {
			return details.NewCreateDetailBadRequest()
		}
		return details.NewCreateDetailCreated().WithPayload(&created)
	})
	api.DetailsUpdateDetailHandler = details.UpdateDetailHandlerFunc(func(params details.UpdateDetailParams) middleware.Responder {
		updated := controller.UpdateDetail(params)
		if updated.Detail != params.Body.Detail {
			return details.NewUpdateDetailBadRequest()
		}
		return details.NewUpdateDetailCreated().WithPayload(&updated)
	})
	api.DetailsDeleteDetailHandler = details.DeleteDetailHandlerFunc(func(params details.DeleteDetailParams) middleware.Responder {
		_ = controller.DeleteDetail(params)
		return details.NewDeleteDetailOK()
	})
	api.DetailsDeleteDetailsHandler = details.DeleteDetailsHandlerFunc(func(params details.DeleteDetailsParams) middleware.Responder {
		_ = controller.DeleteDetails(params)
		return details.NewDeleteDetailsOK()
	})
	api.DetailsGetDetailHandler = details.GetDetailHandlerFunc(func(params details.GetDetailParams) middleware.Responder {
		fetched := controller.GetDetail(params)
		if fetched.Detail == "" {
			return details.NewGetDetailBadRequest()
		}
		return details.NewGetDetailOK().WithPayload(&fetched)
	})
	api.DetailsGetDetailsHandler = details.GetDetailsHandlerFunc(func(params details.GetDetailsParams) middleware.Responder {
		fetched := controller.GetDetails(params)
		return details.NewGetDetailsOK().WithPayload(fetched)
	})

	//--------

	//----DESCRIPTORS----
	api.DescriptorsCreateDescriptorByDetailHandler = descriptors.CreateDescriptorByDetailHandlerFunc(func(params descriptors.CreateDescriptorByDetailParams) middleware.Responder {
		created := controller.CreateDescriptorByDetail(params)
		if created.ID == 0 {
			return descriptors.NewCreateDescriptorByDetailBadRequest()
		}
		return descriptors.NewCreateDescriptorByDetailCreated().WithPayload(&created)
	})
	api.DescriptorsUpdateDescriptorHandler = descriptors.UpdateDescriptorHandlerFunc(func(params descriptors.UpdateDescriptorParams) middleware.Responder {
		updated := controller.UpdateDescriptor(params)
		if updated.Descriptor != params.Body.Descriptor {
			return descriptors.NewUpdateDescriptorBadRequest()
		}
		return descriptors.NewUpdateDescriptorCreated().WithPayload(&updated)
	})
	api.DescriptorsDeleteDescriptorHandler = descriptors.DeleteDescriptorHandlerFunc(func(params descriptors.DeleteDescriptorParams) middleware.Responder {
		_ = controller.DeleteDescriptor(params)
		return descriptors.NewDeleteDescriptorOK()
	})
	api.DescriptorsDeleteDescriptorsHandler = descriptors.DeleteDescriptorsHandlerFunc(func(params descriptors.DeleteDescriptorsParams) middleware.Responder {
		_ = controller.DeleteDescriptors(params)
		return descriptors.NewDeleteDescriptorsOK()
	})
	api.DescriptorsDeleteDescriptorsByDetailHandler = descriptors.DeleteDescriptorsByDetailHandlerFunc(func(params descriptors.DeleteDescriptorsByDetailParams) middleware.Responder {
		_ = controller.DeleteDescriptorsByDetail(params)
		return descriptors.NewDeleteDescriptorsByDetailOK()
	})
	api.DescriptorsGetDescriptorHandler = descriptors.GetDescriptorHandlerFunc(func(params descriptors.GetDescriptorParams) middleware.Responder {
		fetched := controller.GetDescriptor(params)
		if fetched.Descriptor == "" {
			return descriptors.NewGetDescriptorBadRequest()
		}
		return descriptors.NewGetDescriptorOK().WithPayload(&fetched)
	})
	api.DescriptorsGetDescriptorsHandler = descriptors.GetDescriptorsHandlerFunc(func(params descriptors.GetDescriptorsParams) middleware.Responder {
		fetched := controller.GetDescriptors(params)
		if fetched[0].Descriptor == "" {
			return descriptors.NewGetDescriptorsBadRequest()
		}
		return descriptors.NewGetDescriptorsOK().WithPayload(fetched)
	})
	api.DescriptorsGetDescriptorsByDetailHandler = descriptors.GetDescriptorsByDetailHandlerFunc(func(params descriptors.GetDescriptorsByDetailParams) middleware.Responder {
		fetched := controller.GetDescriptorsByDetail(params)
		if len(fetched) == 0 {
			return descriptors.NewGetDescriptorsByDetailNotFound()
		}
		return descriptors.NewGetDescriptorsByDetailOK().WithPayload(fetched)
		// return descriptors.NewGetDescriptorsByDetailOK().WithPayload(fetched)
	})

	//-------

	//----USERS----
	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
		created := controller.CreateUser(params)
		if created.ID == 0 {
			return users.NewCreateUserBadRequest()
		}
		return users.NewCreateUserCreated().WithPayload(&created)
	})
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams) middleware.Responder {
		updated := controller.UpdateUser(params)
		return users.NewUpdateUserCreated().WithPayload(&updated)
	})
	api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(func(params users.DeleteUserParams) middleware.Responder {
		_ = controller.DeleteUser(params)
		return users.NewDeleteUserOK()
	})
	api.UsersDeleteUsersHandler = users.DeleteUsersHandlerFunc(func(params users.DeleteUsersParams) middleware.Responder {
		_ = controller.DeleteUsers(params)
		return users.NewDeleteUsersOK()
	})
	api.UsersGetUserHandler = users.GetUserHandlerFunc(func(params users.GetUserParams) middleware.Responder {
		fetched := controller.GetUser(params)
		if fetched.Username == "" {
			return users.NewGetUserNotFound()
		}
		return users.NewGetUserOK().WithPayload(&fetched)
	})
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
		fetched := controller.GetUsers(params)
		if len(fetched) == 0 {
			return users.NewGetUsersNotFound()
		}
		return users.NewGetUsersOK().WithPayload(fetched)
	})
	//--------

	api.ItemsCreateItemHandler = items.CreateItemHandlerFunc(func(params items.CreateItemParams) middleware.Responder {
		created := controller.CreateItem(params)
		if created.ID == 0 {
			return items.NewCreateItemBadRequest()
		}
		return items.NewCreateItemCreated().WithPayload(&created)
	})
	api.ItemsUpdateItemHandler = items.UpdateItemHandlerFunc(func(params items.UpdateItemParams) middleware.Responder {
		updated := controller.UpdateItem(params)
		return items.NewUpdateItemCreated().WithPayload(&updated)
	})
	api.ItemsDeleteItemHandler = items.DeleteItemHandlerFunc(func(params items.DeleteItemParams) middleware.Responder {
		_ = controller.DeleteItem(params)
		return items.NewDeleteItemOK()
	})
	api.ItemsDeleteItemsHandler = items.DeleteItemsHandlerFunc(func(params items.DeleteItemsParams) middleware.Responder {
		_ = controller.DeleteItems(params)
		return items.NewDeleteItemsOK()
	})
	api.ItemsDeleteItemsByOwnerHandler = items.DeleteItemsByOwnerHandlerFunc(func(params items.DeleteItemsByOwnerParams) middleware.Responder {
		_ = controller.DeleteItemsByOwner(params)
		return items.NewDeleteItemsByOwnerOK()
	})
	api.ItemsDeleteItemsBySubCategoryHandler = items.DeleteItemsBySubCategoryHandlerFunc(func(params items.DeleteItemsBySubCategoryParams) middleware.Responder {
		_ = controller.DeleteItemsBySubCategory(params)
		return items.NewDeleteItemsBySubCategoryOK()
	})
	api.ItemsGetItemHandler = items.GetItemHandlerFunc(func(params items.GetItemParams) middleware.Responder {
		fetched := controller.GetItem(params)
		if fetched.Name == "" {
			return items.NewGetItemsByOwnerNotFound()
		}
		return items.NewGetItemOK().WithPayload(&fetched)
	})
	api.ItemsGetItemsHandler = items.GetItemsHandlerFunc(func(params items.GetItemsParams) middleware.Responder {
		fetched := controller.GetItems(params)
		if len(fetched) == 0 {
			return items.NewGetItemsNotFound()
		}
		return items.NewGetItemsOK().WithPayload(fetched)
	})
	api.ItemsGetItemsByOwnerHandler = items.GetItemsByOwnerHandlerFunc(func(params items.GetItemsByOwnerParams) middleware.Responder {
		fetched := controller.GetItemsByOwner(params)
		if len(fetched) == 0 {
			return items.NewGetItemsByOwnerNotFound()
		}
		return items.NewGetItemsByOwnerOK().WithPayload(fetched)
	})
	api.ItemsGetItemsBySubCategoryHandler = items.GetItemsBySubCategoryHandlerFunc(func(params items.GetItemsBySubCategoryParams) middleware.Responder {
		fetched := controller.GetItemsBySubCategory(params)
		if len(fetched) == 0 {
			return items.NewGetItemsBySubCategoryNotFound()
		}
		return items.NewGetItemsBySubCategoryOK().WithPayload(fetched)
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
