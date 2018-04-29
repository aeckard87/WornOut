package controllers

import (
	"fmt"

	dbpkg "github.com/aeckard87/goServe/db"
	model "github.com/aeckard87/goServe/models"
)

// CreateSubCategoryByCategory returns SubCategory
func CreateSubCategoryByCategory(subcategory model.SubCategory, id string) model.SubCategory {
	db := dbpkg.Connect()
	defer db.Close()

	// db.Create(&params.Body).Related(&category)
	db.Exec("insert into sub_categories (subcategory,category_id) values (?,?)", subcategory.Subcategory, id)

	db.Where("subcategory = ?", subcategory.Subcategory).Find(&subcategory)
	return subcategory
}

// UpdateSubCategory returns Subcategory
func UpdateSubCategory(subcategory model.SubCategory, id string) model.SubCategory {
	fmt.Println("Update Subcateogry")
	db := dbpkg.Connect()
	defer db.Close()

	db.Model(&subcategory).Where("id = ?", id).Update("subcategory", subcategory)
	db.Where("id = ?", id).Find(&subcategory)
	return subcategory

}

// DeleteSubCategory returns empty SubCategory
func DeleteSubCategory(id string) model.SubCategory {
	fmt.Println("Delete Subcateogry")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.SubCategory

	db.Where("id = ?", id).Delete(&sc)

	return sc

}

// DeleteSubCategoriesByCategory returns empty SubCategory given Category.ID
func DeleteSubCategoriesByCategory(id string) []model.SubCategory {
	fmt.Println("Delete Subcateogry by Category")
	db := dbpkg.Connect()
	defer db.Close()

	var sc []model.SubCategory

	db.Where("category_id = ?", id).Delete(&sc)

	return sc

}

// DeleteSubCategories returns empty SubCategory
func DeleteSubCategories() []model.SubCategory {
	fmt.Println("Delete Subcateogries")
	db := dbpkg.Connect()
	defer db.Close()

	var sc []model.SubCategory

	db.Delete(&sc)

	return sc

}

// GetSubCategory returns SubCategory given SubCategory.ID
func GetSubCategory(id string) model.SubCategory {
	fmt.Println("GetSubCategory")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.SubCategory

	db.Where("id = ?", id).Find(&sc)

	return sc

}

// GetSubCategoriesByCategory returns SubCategories given Category.ID
func GetSubCategoriesByCategory(id string) []model.SubCategory {
	fmt.Println("GetSubCategoriesByCategory")
	db := dbpkg.Connect()
	defer db.Close()

	var sc []model.SubCategory

	db.Where("category_id = ?", id).Find(&sc)

	return sc
}

// GetSubCategories returns SubCategories
func GetSubCategories() []model.SubCategory {
	fmt.Println("GetSubCategories")
	db := dbpkg.Connect()
	defer db.Close()

	var sc []model.SubCategory

	db.Find(&sc)

	return sc

}
