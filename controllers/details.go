package controllers

import (
	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/details"
)

// CreateDetail returns Detail
func CreateDetail(params details.CreateDetailParams) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	db.Create(&params.Body)

	var d model.Detail

	db.Where("detail = ?", params.Body.Detail).Find(&d)
	return d
}

// UpdateDetail returns Detail
func UpdateDetail(params details.UpdateDetailParams) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	var d model.Detail

	db.Model(&d).Where("id = ?", params.ID).Update("detail", params.Body.Detail)
	//cascade items values

	d.ID = params.ID
	return d

}

// DeleteDetail returns empty Detail
func DeleteDetail(params details.DeleteDetailParams) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	var d model.Detail
	d.ID = params.ID

	db.Delete(&d)

	return d

}

// DeleteDetails returns empty Detail
func DeleteDetails(params details.DeleteDetailsParams) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	var d model.Detail

	db.Delete(&d)

	return d

}

// GetDetail returns Detail given Detail.ID
func GetDetail(params details.GetDetailParams) model.Detail {
	db := dbpkg.Connect()
	defer db.Close()

	var d model.Detail
	d.ID = params.ID

	db.First(&d)

	return d

}

// GetDetails returns Details
func GetDetails(params details.GetDetailsParams) model.Details {
	db := dbpkg.Connect()
	defer db.Close()

	var d model.Details

	db.Find(&d)

	return d

}
