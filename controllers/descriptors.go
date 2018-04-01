package controllers

import (
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/descriptors"
)

func CreateDescriptorByDetail(params descriptors.CreateDescriptorByDetailParams) model.Descriptor {
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.Descriptor

	// db.Create(&params.Body).Related(&detail)
	db.Exec("insert into descriptors (descriptor,detail_id) values (?,?)", params.Body.Descriptor, params.ID)

	db.Where("descriptor = ?", params.Body.Descriptor).Find(&sc)
	return sc
}

func UpdateDescriptor(params descriptors.UpdateDescriptorParams) model.Descriptor {
	fmt.Println("Update Descriptor")
	db := dbpkg.Connect()

	defer db.Close()

	var sc model.Descriptor

	db.Model(&sc).Where("id = ?", params.ID).Update("descriptor", params.Body.Descriptor)
	db.Where("id = ?", params.ID).Find(&sc)
	return sc

}

func DeleteDescriptor(params descriptors.DeleteDescriptorParams) model.Descriptor {
	fmt.Println("Delete Descriptor")
	db := dbpkg.Connect()

	defer db.Close()

	var sc model.Descriptor

	db.Where("id = ?", params.ID).Delete(&sc)

	return sc

}

func DeleteDescriptorsByDetail(params descriptors.DeleteDescriptorsByDetailParams) model.Descriptor {
	fmt.Println("Delete Descriptors by Detail")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var sc model.Descriptor

	db.Where("detail_id = ?", params.ID).Delete(&sc)

	return sc

}

func DeleteDescriptors(params descriptors.DeleteDescriptorsParams) model.Descriptor {
	fmt.Println("Delete Descriptors")
	db := dbpkg.Connect()

	defer db.Close()

	// fmt.Println("Delete Record ID", params.ID)

	var sc model.Descriptor

	db.Delete(&sc)

	return sc

}
func GetDescriptor(params descriptors.GetDescriptorParams) model.Descriptor {
	fmt.Println("GetDescriptor")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.Descriptor

	db.Where("id = ?", params.ID).Find(&sc)

	return sc

}

func GetDescriptorsByDetail(params descriptors.GetDescriptorsByDetailParams) model.Descriptors {
	fmt.Println("GetDescriptorsByDetail")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.Descriptors

	db.Where("detail_id = ?", params.ID).Find(&sc)

	return sc
}

func GetDescriptors(params descriptors.GetDescriptorsParams) model.Descriptors {
	fmt.Println("GetDescriptors")
	db := dbpkg.Connect()
	defer db.Close()

	var sc model.Descriptors

	db.Find(&sc)

	return sc

}
