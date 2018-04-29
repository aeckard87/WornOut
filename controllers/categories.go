package controllers

import (
	"strconv"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	// "github.com/aeckard87/test/client/categories"
)

// CreateCategory returns created result of type Category.
func CreateCategory(category model.Category) model.Category {
	db := dbpkg.Connect()
	defer db.Close()

	db.Create(&category)

	db.Where("category = ?", category.Category).Find(&category)
	return category
}

// UpdateCategory returns updated result of type Category.
func UpdateCategory(category model.Category, id string) model.Category {
	db := dbpkg.Connect()
	defer db.Close()

	db.Model(&category).Where("id = ?", id).Update("category", category)
	category.ID, _ = strconv.ParseInt(id, 10, 64)
	return category

}

// DeleteCategory returns empty result of type Category.
func DeleteCategory(id string) model.Category {
	db := dbpkg.Connect()
	defer db.Close()

	var category model.Category
	category.ID, _ = strconv.ParseInt(id, 10, 64)

	db.Delete(&category)

	return category

}

// DeleteCategories returns empty result of type Categories
func DeleteCategories() []model.Category {
	db := dbpkg.Connect()
	defer db.Close()

	var categories []model.Category
	db.Delete(&categories)

	return categories

}

// GetCategory returns a Category.
func GetCategory(id string) model.Category {
	db := dbpkg.Connect()
	defer db.Close()

	var category model.Category
	category.ID, _ = strconv.ParseInt(id, 10, 64)

	db.First(&category)

	return category

}

// GetCategories returns Categories.
func GetCategories() []model.Category {
	db := dbpkg.Connect()
	defer db.Close()

	var categories []model.Category

	db.Find(&categories)

	return categories

}
