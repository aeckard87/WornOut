package controls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	model "github.com/aeckard87/WornOut/models"
)

func GetBodyBytes(r *http.Request) []byte {
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	return bodyBytes
}

// Category Helpers
func StringifyCat(category model.Category) string {
	//Marshal
	b, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func StringifyCats(categories []model.Category) string {
	//Marshal
	b, err := json.Marshal(categories)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

//SubCategory Helpers
func StringifySubCat(subcategory model.SubCategory) string {
	//Marshal
	b, err := json.Marshal(subcategory)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func StringifySubCats(subcategories []model.SubCategory) string {
	//Marshal
	b, err := json.Marshal(subcategories)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

//Details Helpers
func StringifyDet(detail model.Detail) string {
	//Marshal
	b, err := json.Marshal(detail)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func StringifyDets(details []model.Detail) string {
	//Marshal
	b, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

//Descriptors Helpers
func StringifyDesc(descriptor model.Descriptor) string {
	//Marshal
	b, err := json.Marshal(descriptor)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func StringifyDescs(descriptors []model.Descriptor) string {
	//Marshal
	b, err := json.Marshal(descriptors)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

//Users Helpers
func StringifyUser(user model.User) string {
	//Marshal
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func StringifyUsers(users []model.User) string {
	//Marshal
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

//Items Helpers
func StringifyItem(item model.Item) string {
	//Marshal
	b, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
func StringifyItems(items []model.Item) string {
	//Marshal
	b, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
