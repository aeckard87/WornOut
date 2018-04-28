package controllers

import (
	"fmt"
	"log"

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
	var descriptors []model.Descriptor

	// descriptions, err := json.Marshal(params.Body.Descriptions)
	descriptors = params.Body.Descriptions
	// err := json.Unmarshal([]byte(params.Body.Descriptions), &descriptors)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(descriptions))
	db.Exec("INSERT INTO items (name,sub_category_id,user_id) VALUES (?,?,?)", params.Body.Name, params.Body.SubCategoryID, params.Body.UserID)
	// , string(descriptions)

	item := GetItemByName(params.Body.Name)

	for index, _ := range descriptors {
		fmt.Println(descriptors[index].Descriptor)
		db.Exec("INSERT INTO items2descriptors (item_id,descriptor_id) VALUES (?,?)", item.ID, descriptors[index].ID)
	}

	item = GetItemByName(params.Body.Name)

	return item
}

// GetItemByName returns Item given Item.Name
func GetItemByName(name string) model.Item {
	db := dbpkg.Connect()
	defer db.Close()
	var item model.Item

	//This Gets everything BUT descriptions
	db.Select("*").Where("name = ?", name).Find(&item)

	// Gets Descriptions
	rows, err := db.Raw("SELECT descriptors.* FROM  descriptors JOIN items2descriptors ON descriptors.id = items2descriptors.descriptor_id WHERE items2descriptors.item_id = ?;", fmt.Sprint(item.ID)).Rows()

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var descriptions []model.Descriptor
	for rows.Next() {
		var descriptor model.Descriptor

		if err := rows.Scan(&descriptor.ID, &descriptor.Descriptor, &descriptor.DetailID); err != nil {
			log.Fatal(err)
		}
		descriptions = append(descriptions, descriptor)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	item.Descriptions = descriptions

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

	params.Body.ID = params.ID

	// db.Where("id = ?", params.ID).Table("items").Updates(params.Body.Name)
	db.Model(&item).Where("id = ?", params.ID).Update("name", params.Body.Name)
	db.Model(&item).Where("id = ?", params.ID).Update("user_id", params.Body.UserID)
	db.Model(&item).Where("id = ?", params.ID).Update("sub_category_id", params.Body.SubCategoryID)

	//This Gets everything BUT descriptions
	db.Select("*").Where("id = ?", params.ID).Find(&item)

	//rm all descriptors then add them back *yea I know... how horrible of me.
	db.Exec("DELETE FROM items2descriptors WHERE item_id = ?;", params.ID)

	descriptors := params.Body.Descriptions
	for index := range descriptors {
		fmt.Println(descriptors[index].Descriptor)
		db.Exec("INSERT INTO items2descriptors (item_id,descriptor_id) VALUES (?,?)", params.ID, descriptors[index].ID)
	}

	item.Descriptions = getItemsDescriptionByID(params.ID)

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
	db.Select("*").Where("id = ?", params.ID).Find(&item)

	// Gets Descriptions
	item.Descriptions = getItemsDescriptionByID(params.ID)

	return item

}

// GetItems returns Items
func GetItems(params items.GetItemsParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Select("*").Table("items").Scan(&items)
	for _, item := range items {

		item.Descriptions = getItemsDescriptionByID(item.ID)
	}

	return items

}

// GetItemsByOwner returns Items given User.ID in params
func GetItemsByOwner(params items.GetItemsByOwnerParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Select("*").Table("items").Where("user_id=?", params.ID).Scan(&items)
	for _, item := range items {
		item.Descriptions = getItemsDescriptionByID(item.ID)
	}

	return items

}

// GetItemsBySubCategory returns Items given SubCategory.ID in params
func GetItemsBySubCategory(params items.GetItemsBySubCategoryParams) model.Items {
	db := dbpkg.Connect()
	defer db.Close()

	var items model.Items

	db.Select("*").Table("items").Where("sub_category_id=?", params.ID).Scan(&items)
	for _, item := range items {
		item.Descriptions = getItemsDescriptionByID(item.ID)
	}

	return items

}

func getItemsDescriptionByID(id int64) []model.Descriptor {
	db := dbpkg.Connect()
	defer db.Close()
	// Gets Descriptions

	var descriptions []model.Descriptor
	rows, err := db.Raw("SELECT descriptors.* FROM  descriptors JOIN items2descriptors ON descriptors.id = items2descriptors.descriptor_id WHERE items2descriptors.item_id = ?;", id).Rows()

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var descriptor model.Descriptor

		if err := rows.Scan(&descriptor.ID, &descriptor.Descriptor, &descriptor.DetailID); err != nil {
			log.Fatal(err)
		}
		descriptions = append(descriptions, descriptor)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return descriptions
}
