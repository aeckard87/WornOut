package controllers

import (
	"encoding/json"
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/items"
)

type Descriptions struct {
	Descriptors []*model.Descriptors
}

func CreateItem(params items.CreateItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	descriptions, err := json.Marshal(params.Body.Descriptions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(descriptions))
	db.Exec("INSERT INTO items (name,sub_category_id,descriptions,user_id) VALUES (?,?,?,?)", params.Body.Name, params.Body.SubCategoryID, string(descriptions), params.Body.UserID)

	item := GetItemByName(params.Body.Name)

	return item
}

func GetItemByName(name string) model.Item {
	db := dbpkg.Connect()
	defer db.Close()
	var item model.Item

	//This Gets everything BUT descriptions
	db.Where("name = ?", name).Find(&item)
	db.Select("user_id").Table("items").Where("name = ?", name).Scan(&item)
	// db.DB().QueryRow("SELECT * FROM items WHERE name = ?", name).Scan(&item.ID, &item.Name, &item.SubCategoryID, &item.UserID, &item.Descriptions)
	// a, err := json.Marshal(item)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Item:\t", string(a))

	// This properly gets descriptions!

	var blob string
	db.DB().QueryRow("SELECT descriptions FROM items WHERE name = ?", name).Scan(&blob)
	// fmt.Println(blob)
	err := json.Unmarshal([]byte(blob), &item.Descriptions)
	if err != nil {
		fmt.Println(err)
	}

	//Marshal
	// b, err := json.Marshal(item.Descriptions)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Descriptors:\t", string(b))

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

func GetItem(params items.GetItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()
	var item model.Item

	//This Gets everything BUT descriptions
	db.Where("id = ?", params.ID).Find(&item)
	db.Select("user_id").Table("items").Where("id = ?", params.ID).Scan(&item)

	// This properly gets descriptions!
	var blob string
	db.DB().QueryRow("SELECT descriptions FROM items WHERE id = ?", params.ID).Scan(&blob)
	// fmt.Println(blob)
	err := json.Unmarshal([]byte(blob), &item.Descriptions)
	if err != nil {
		fmt.Println(err)
	}

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
