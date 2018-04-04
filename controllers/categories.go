package controllers

import (
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/categories"
)

var count int

//CreateCategory
func CreateCategory(params categories.CreateCategoryParams) model.Category {
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("New Record for category", params.Category)

	db.Create(&params.Body)

	var category model.Category

	db.Where("category = ?", params.Body.Category).Find(&category)
	return category
}

//UpdateCategory
func UpdateCategory(params categories.UpdateCategoryParams) model.Category {
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Update Record to category", params.Body.Category)

	var category model.Category

	db.Model(&category).Where("id = ?", params.ID).Update("category", params.Body.Category)
	// fmt.Println("UPDATED", category)
	category.ID = params.ID
	return category

}

//DeleteCategory
func DeleteCategory(params categories.DeleteCategoryParams) model.Category {
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var category model.Category
	category.ID = params.ID

	db.Delete(&category)

	return category

}

//DeleteCategories
func DeleteCategories(params categories.DeleteCategoriesParams) model.Category {
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var category model.Category

	db.Delete(&category)

	return category

}

//GetCategory
func GetCategory(params categories.GetCategoryParams) model.Category {
	db := dbpkg.Connect()

	defer db.Close()

	fmt.Println("Get Record ID", params.ID)

	var category model.Category
	category.ID = params.ID

	db.First(&category)

	return category

}

//GetCategories
func GetCategories(params categories.GetCategoriesParams) model.Categories {
	db := dbpkg.Connect()

	defer db.Close()

	fmt.Println("Get All Records")

	var categories model.Categories

	db.Find(&categories)

	return categories

}
