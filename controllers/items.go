package controllers

import (
	"encoding/json"
	"fmt"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/items"
)

// Descriptions struct type
type Descriptions struct {
	Descriptors []*model.Descriptors
}

// CreateItem returns Item
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

// GetItemByName returns Item given Item.Name
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

// UpdateItem returns Item
func UpdateItem(params items.UpdateItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Model(&item).Where("id = ?", params.ID).Updates(params.Body)
	// fmt.Println("UPDATED", item)
	item.ID = params.ID
	return item

}

// DeleteItem returns empty Item
func DeleteItem(params items.DeleteItemParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("id=?", params.ID).Delete(&item)

	return item

}

// DeleteItems returns empty Item
func DeleteItems(params items.DeleteItemsParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Delete(&item)

	return item

}

// DeleteItemsByOwner returns Itm given User.ID
func DeleteItemsByOwner(params items.DeleteItemsByOwnerParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("user_id=?", params.ID).Delete(&item)

	return item

}

// DeleteItemsBySubCategory returns Item given SubCategory.ID
func DeleteItemsBySubCategory(params items.DeleteItemsBySubCategoryParams) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("sub_category_id=?", params.ID).Delete(&item)

	return item

}

// GetItem returns Item given Item.ID
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

// GetItems returns Items
func GetItems(params items.GetItemsParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Find(&items)
	for _, item := range items {
		db.Select("user_id").Table("items").Where("id = ?", item.ID).Scan(&item)

		// This properly gets descriptions!
		var blob string
		db.DB().QueryRow("SELECT descriptions FROM items WHERE id = ?", item.ID).Scan(&blob)
		// fmt.Println(blob)
		err := json.Unmarshal([]byte(blob), &item.Descriptions)
		if err != nil {
			fmt.Println(err)
		}
	}

	return items

}

// GetItemsByOwner returns Items given User.ID in params
func GetItemsByOwner(params items.GetItemsByOwnerParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Where("user_id=?", params.ID).Find(&items)
	for _, item := range items {
		db.Select("user_id").Table("items").Where("id = ?", item.ID).Scan(&item)

		// This properly gets descriptions!
		var blob string
		db.DB().QueryRow("SELECT descriptions FROM items WHERE id = ?", item.ID).Scan(&blob)
		// fmt.Println(blob)
		err := json.Unmarshal([]byte(blob), &item.Descriptions)
		if err != nil {
			fmt.Println(err)
		}
	}

	return items

}

// GetItemsBySubCategory returns Items given SubCategory.ID in params
func GetItemsBySubCategory(params items.GetItemsBySubCategoryParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Where("sub_category_id=?", params.ID).Find(&items)
	for _, item := range items {
		db.Select("user_id").Table("items").Where("id = ?", item.ID).Scan(&item)

		// This properly gets descriptions!
		var blob string
		db.DB().QueryRow("SELECT descriptions FROM items WHERE id = ?", item.ID).Scan(&blob)
		// fmt.Println(blob)
		err := json.Unmarshal([]byte(blob), &item.Descriptions)
		if err != nil {
			fmt.Println(err)
		}
	}

	return items

}
