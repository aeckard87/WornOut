package controllers

import (
	"fmt"
	"strconv"

	dbpkg "github.com/aeckard87/goServe/db"
	model "github.com/aeckard87/goServe/models"
)

// CreateDetail returns Detail
func CreateDetail(detail model.Detail, id string) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	db.Create(&detail)

	db.Where("detail = ?", detail.Detail).Find(&detail)
	return detail
}

// UpdateDetail returns updated result of type Detail.
func UpdateDetail(detail model.Detail, id string) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	db.Model(&detail).Where("id = ?", id).Update("detail", detail)
	detail.ID, _ = strconv.ParseInt(id, 10, 64)
	return detail

}

// DeleteDetail returns empty Detail
func DeleteDetail(detail model.Detail, id string) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	detail.ID, _ = strconv.ParseInt(id, 10, 64)

	db.Delete(&detail)

	return detail
}

// DeleteDetails returns empty result of type Details
func DeleteDetails() []model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	var details []model.Detail
	db.Delete(&details)

	return details
}

// GetDetail returns a Detail.
func GetDetail(detail model.Detail, id string) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()
	fmt.Println("ID: " + id)
	detail.ID, _ = strconv.ParseInt(id, 10, 64)

	db.First(&detail)

	return detail

}

// GetDetails returns Details.
func GetDetails() []model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	var details []model.Detail

	db.Find(&details)

	return details

}
