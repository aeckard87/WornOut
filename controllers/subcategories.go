package controllers

import (
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/subcategories"
)

// CreateSubCategoryByCategory returns SubCategory
func CreateSubCategoryByCategory(params subcategories.CreateSubCategoryByCategoryParams) model.SubCategory {
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.SubCategory

	// db.Create(&params.Body).Related(&category)
	db.Exec("insert into sub_categories (subcategory,category_id) values (?,?)", params.Body.Subcategory, params.ID)

	db.Where("subcategory = ?", params.Body.Subcategory).Find(&sc)
	return sc
}

// UpdateSubCategory returns Subcategory
func UpdateSubCategory(params subcategories.UpdateSubCategoryParams) model.SubCategory {
	fmt.Println("Update Subcateogry")
	db := dbpkg.Connect()

	defer db.Close()

	var sc model.SubCategory

	db.Model(&sc).Where("id = ?", params.ID).Update("subcategory", params.Body.Subcategory)
	db.Where("id = ?", params.ID).Find(&sc)
	return sc

}

// DeleteSubCategory returns empty SubCategory
func DeleteSubCategory(params subcategories.DeleteSubCategoryParams) model.SubCategory {
	fmt.Println("Delete Subcateogry")
	db := dbpkg.Connect()

	defer db.Close()

	var sc model.SubCategory

	db.Where("id = ?", params.ID).Delete(&sc)

	return sc

}

// DeleteSubcategoriesByCategory returns empty SubCategory given Category.ID
func DeleteSubcategoriesByCategory(params subcategories.DeleteSubCategoriesByCategoryParams) model.SubCategory {
	fmt.Println("Delete Subcateogry by Category")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var sc model.SubCategory

	db.Where("category_id = ?", params.ID).Delete(&sc)

	return sc

}

// DeleteSubCategories returns empty SubCategory
func DeleteSubCategories(params subcategories.DeleteSubCategoriesParams) model.SubCategory {
	fmt.Println("Delete Subcateogries")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var sc model.SubCategory

	db.Delete(&sc)

	return sc

}

// GetSubCategory returns SubCategory given SubCategory.ID
func GetSubCategory(params subcategories.GetSubCategoryParams) model.SubCategory {
	fmt.Println("GetSubCategory")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.SubCategory

	db.Where("id = ?", params.ID).Find(&sc)

	return sc

}

// GetSubcategoriesByCategory returns SubCategories given Category.ID
func GetSubcategoriesByCategory(params subcategories.GetSubCategoriesByCategoryParams) model.SubCategories {
	fmt.Println("GetSubCategoriesByCategory")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.SubCategories

	db.Where("category_id = ?", params.ID).Find(&sc)

	return sc
}

// GetSubcategories returns SubCategories
func GetSubcategories(params subcategories.GetSubCategoriesParams) model.SubCategories {
	fmt.Println("GetSubCategories")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.SubCategories

	db.Find(&sc)

	return sc

}
