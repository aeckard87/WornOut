package controllers

import (
	"fmt"

	dbpkg "github.com/aeckard87/goServe/db"
	model "github.com/aeckard87/goServe/models"
)

// CreateDescriptorByDetail returns type Descriptor.
func CreateDescriptorByDetail(descriptor model.Descriptor, id string) model.Descriptor {
	db := dbpkg.Connect()
	defer db.Close()

	// db.Create(&params.Body).Related(&detail)
	db.Exec("insert into descriptors (descriptor,detail_id) values (?,?)", descriptor.Descriptor, id)

	db.Where("descriptor = ?", descriptor.Descriptor).Find(&descriptor)
	return descriptor
}

// UpdateDescriptor returns type Descriptor.
func UpdateDescriptor(descriptor model.Descriptor, id string) model.Descriptor {
	fmt.Println("Update Descriptor")
	db := dbpkg.Connect()
	defer db.Close()

	db.Model(&descriptor).Where("id = ?", id).Update("descriptor", descriptor)
	db.Where("id = ?", id).Find(&descriptor)
	return descriptor

}

// DeleteDescriptor returns empty Descriptor.
func DeleteDescriptor(id string) model.Descriptor {
	fmt.Println("Delete Descriptor")
	db := dbpkg.Connect()
	defer db.Close()

	var descriptor model.Descriptor

	db.Where("id = ?", id).Delete(&descriptor)

	return descriptor
}

// DeleteDescriptorsByDetail returns empty Descriptor given Detail.ID.
func DeleteDescriptorsByDetail(id string) []model.Descriptor {
	fmt.Println("Delete Descriptors by Detail")
	db := dbpkg.Connect()
	defer db.Close()

	var descriptors []model.Descriptor

	db.Where("detail_id = ?", id).Delete(&descriptors)

	return descriptors

}

// DeleteDescriptors returns empty Descriptor.
func DeleteDescriptors() []model.Descriptor {
	fmt.Println("Delete Descriptors")
	db := dbpkg.Connect()
	defer db.Close()

	var descriptor []model.Descriptor

	db.Delete(&descriptor)

	return descriptor
}

// GetDescriptor returns Descriptor given Descriptor.ID.
func GetDescriptor(id string) model.Descriptor {
	fmt.Println("GetDescriptor")
	db := dbpkg.Connect()
	defer db.Close()

	var descriptor model.Descriptor

	db.Where("id = ?", id).Find(&descriptor)

	return descriptor

}

// GetDescriptorsByDetail returns Descriptors given Detail.ID.
func GetDescriptorsByDetail(id string) []model.Descriptor {
	fmt.Println("GetDescriptorsByDetail")
	db := dbpkg.Connect()
	defer db.Close()

	var descriptors []model.Descriptor

	db.Where("detail_id = ?", id).Find(&descriptors)

	return descriptors
}

// GetDescriptors returns Descriptors.
func GetDescriptors() []model.Descriptor {
	fmt.Println("GetDescriptors")
	db := dbpkg.Connect()
	defer db.Close()

	var descriptors []model.Descriptor

	db.Find(&descriptors)

	return descriptors

}
