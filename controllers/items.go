package controllers

import (
	"encoding/json"
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/items"
)

func CreateItem(params items.CreateItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	descriptions, err := json.Marshal(params.Body.Descriptions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(descriptions))
	db.Exec("INSERT INTO items (name,sub_category_id,descriptions,user_id) VALUES (?,?,?,?)", params.Body.Name, params.Body.SubCategoryID, string(descriptions), params.Body.UserID)
	// db.Create(params.Body)

	var item model.Item
	// var descriptors model.Descriptors
	var test []byte

	// db.Where("name = ?", params.Body.Name).Find(&item)
	db.Raw("SELECT id, name, sub_category_id,user_id FROM items WHERE name = ?", params.Body.Name).Scan(&item)
	db.Raw("SELECT descriptions FROM items WHERE name = ?", params.Body.Name).Scan(&test)
	fmt.Println(test)
	// b := json.Unmarshal(test, &item.Descriptions)
	// fmt.Println("b", b)

	// item.Descriptions = descriptors

	// b, err := json.Marshal(item)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Item:", string(b))

	return item
}

func UpdateItem(params items.UpdateItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Model(&item).Where("id = ?", params.ID).Updates(params.Body)
	// fmt.Println("UPDATED", item)
	item.ID = params.ID
	return item

}

func DeleteItem(params items.DeleteItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("id=?", params.ID).Delete(&item)

	return item

}

func DeleteItems(params items.DeleteItemsParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Delete(&item)

	return item

}

func DeleteItemsByOwner(params items.DeleteItemsByOwnerParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("owner_id=?", params.ID).Delete(&item)

	return item

}

func DeleteItemsBySubCategory(params items.DeleteItemsBySubCategoryParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("sub_category_id=?", params.ID).Delete(&item)

	return item

}

func GetItem(params items.GetItemParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Items

	db.Where("id=?", params.ID).Find(&item)

	return item

}

func GetItems(params items.GetItemsParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Find(&items)

	return items

}

func GetItemsByOwner(params items.GetItemsByOwnerParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Where("owner_id=?", params.ID).Find(&items)

	return items

}

func GetItemsBySubCategory(params items.GetItemsBySubCategoryParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Where("sub_category_id=?", params.ID).Find(&items)

	return items

}
