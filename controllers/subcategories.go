package controllers

import (
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/categories"
	"github.com/aeckard87/WornOut/restapi/operations/subcategories"
)

func CreateSubCategory(params subcategories.CreateSubCategoryParams) model.SubCategory {
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("New Record for SubCategory", params.SubCategory)

	db.Create(&params.Body)

	var subcategory model.SubCategory

	db.Where("subcategory = ?", params.Body.Subcategory).Find(&subcategory)
	return subcategory
}

func CreateSubCategoryByCategory(params subcategories.CreateSubCategoryByCategoryParams) model.SubCategory {
	fmt.Println("Create Subcateogry by Category")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("New Record for SubCategory", params.SubCategory)
	// var category model.Category
	var subcategory model.SubCategory

	// category.ID = params.ID

	// fmt.Println("CAT ID:", category.ID)
	fmt.Println("PARMS ID:", params.ID)
	fmt.Println("BODY ID:", params.Body.ID)

	// db.Create(&params.Body).Related(&category)
	db.Exec("insert into sub_categories (subcategory,category_id) values (?,?)", params.Body.Subcategory, params.ID)

	db.Where("subcategory = ?", params.Body.Subcategory).Find(&subcategory)
	return subcategory
}

func UpdateSubCategory(params subcategories.UpdateSubCategoryParams) model.SubCategory {
	fmt.Println("Update Subcateogry")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Update Record to SubCategory", params.Body.SubCategory)

	var subcategory model.SubCategory

	db.Model(&subcategory).Where("id = ?", params.ID).Update("subcategory", params.Body.Subcategory)
	// fmt.Println("UPDATED", SubCategory)
	subcategory.ID = params.ID
	return subcategory

}

func DeleteSubCategory(params subcategories.DeleteSubCategoryParams) model.SubCategory {
	fmt.Println("Delete Subcateogry")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var SubCategory model.SubCategory
	SubCategory.ID = params.ID

	db.Delete(&SubCategory)

	return SubCategory

}

func DeleteSubcategoriesByCategory(params subcategories.DeleteSubCategoriesByCategoryParams) model.SubCategory {
	fmt.Println("Delete Subcateogry by Category")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var SubCategory model.SubCategory

	db.Delete(&SubCategory)

	return SubCategory

}

func DeleteSubCategories(params subcategories.DeleteSubCategoriesParams) model.SubCategory {
	fmt.Println("Delete Subcateogries")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var SubCategory model.SubCategory

	db.Delete(&SubCategory)

	return SubCategory

}
func GetSubCategory(params subcategories.GetSubCategoryParams) model.SubCategory {
	db := dbpkg.Connect()

	defer db.Close()

	fmt.Println("Get Record ID", params.ID)

	var SubCategory model.SubCategory
	SubCategory.ID = params.ID

	db.First(&SubCategory)

	return SubCategory

}

func Getsubcategories(params categories.GetCategoriesParams) model.Categories {
	db := dbpkg.Connect()

	defer db.Close()

	fmt.Println("Get All Records")

	var subcategories model.Categories

	db.Find(&subcategories)

	return subcategories

}
