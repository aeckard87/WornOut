package controllers

import (
	"strconv"

	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
)

// CreateUser returns User
func CreateUser(user model.User) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	db.Create(&user)

	db.Where("username = ?", user.Username).Find(&user)
	return user
}

// UpdateUser returns User
func UpdateUser(user model.User, id string) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	db.Model(&user).Where("id = ?", id).Updates(user)
	return user

}

// DeleteUser returns empty User
func DeleteUser(id string) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user model.User
	user.ID, _ = strconv.ParseInt(id, 10, 64)

	db.Delete(&user)

	return user

}

// DeleteUsers returns empty User
func DeleteUsers() []model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user []model.User
	db.Delete(&user)

	return user

}

// GetUser returns User given User.ID
func GetUser(id string) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user model.User
	user.ID, _ = strconv.ParseInt(id, 10, 64)

	db.Find(&user)

	return user

}

// GetUsers returns Users
func GetUsers() []model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var users []model.User

	db.Find(&users)

	return users

}
