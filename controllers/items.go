package controllers

import (
	"fmt"
	"log"
	"strconv"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
)

// CreateItem returns Item
func CreateItem(item model.Item) model.Item {
	db := dbpkg.Connect()
	defer db.Close()
	var descriptors []model.Descriptor

	// descriptions, err := json.Marshal(params.Body.Descriptions)
	descriptors = item.Descriptions
	// err := json.Unmarshal([]byte(params.Body.Descriptions), &descriptors)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(descriptions))
	db.Exec("INSERT INTO items (name,sub_category_id,user_id) VALUES (?,?,?)", item.Name, item.SubCategoryID, item.UserID)
	// , string(descriptions)

	item = GetItemByName(item.Name)

	for index, _ := range descriptors {
		fmt.Println(descriptors[index].Descriptor)
		db.Exec("INSERT INTO items2descriptors (item_id,descriptor_id) VALUES (?,?)", item.ID, descriptors[index].ID)
	}

	item = GetItemByName(item.Name)

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
func UpdateItem(item model.Item, id string) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	item.ID, _ = strconv.ParseInt(id, 10, 64)
	// db.Where("id = ?", params.ID).Table("items").Updates(params.Body.Name)
	db.Model(&item).Where("id = ?", id).Update("name", item.Name)
	db.Model(&item).Where("id = ?", id).Update("user_id", item.UserID)
	db.Model(&item).Where("id = ?", id).Update("sub_category_id", item.SubCategoryID)

	//This Gets everything BUT descriptions
	db.Select("*").Where("id = ?", id).Find(&item)

	//rm all descriptors then add them back *yea I know... how horrible of me.
	db.Exec("DELETE FROM items2descriptors WHERE item_id = ?;", id)

	descriptors := item.Descriptions
	for index := range descriptors {
		fmt.Println(descriptors[index].Descriptor)
		db.Exec("INSERT INTO items2descriptors (item_id,descriptor_id) VALUES (?,?)", id, descriptors[index].ID)
	}

	item.Descriptions = getItemsDescriptionByID(item.ID)

	return item
}

// DeleteItem returns empty Item
func DeleteItem(id string) model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var item model.Item

	db.Where("id=?", id).Delete(&item)

	return item
}

// DeleteItems returns empty Item
func DeleteItems() []model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var items []model.Item

	db.Delete(&items)

	return items

}

// DeleteItemsByOwner returns Itm given User.ID
func DeleteItemsByOwner(id string) []model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var items []model.Item

	db.Where("user_id=?", id).Delete(&items)

	return items

}

// DeleteItemsBySubCategory returns Item given SubCategory.ID
func DeleteItemsBySubCategory(id string) []model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var items []model.Item

	db.Where("sub_category_id=?", id).Delete(&items)

	return items

}

// GetItem returns Item given Item.ID
func GetItem(id string) model.Item {
	db := dbpkg.Connect()
	defer db.Close()
	var item model.Item

	//This Gets everything BUT descriptions
	db.Select("*").Where("id = ?", id).Find(&item)

	// Gets Descriptions
	item.Descriptions = getItemsDescriptionByID(item.ID)

	return item
}

// GetItems returns Items
func GetItems() []model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var items []model.Item

	db.Select("*").Table("items").Scan(&items)
	for _, item := range items {

		item.Descriptions = getItemsDescriptionByID(item.ID)
	}

	return items

}

// GetItemsByOwner returns Items given User.ID in params
func GetItemsByOwner(id string) []model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var items []model.Item

	db.Select("*").Table("items").Where("user_id=?", id).Scan(&items)
	for _, item := range items {
		item.Descriptions = getItemsDescriptionByID(item.ID)
	}

	return items

}

// GetItemsBySubCategory returns Items given SubCategory.ID in params
func GetItemsBySubCategory(id string) []model.Item {
	db := dbpkg.Connect()
	defer db.Close()

	var items []model.Item

	db.Select("*").Table("items").Where("sub_category_id=?", id).Scan(&items)
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
