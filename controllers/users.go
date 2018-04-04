package controllers

import (
	dbpkg "github.com/aeckard87/WornOut/db"
	model "github.com/aeckard87/WornOut/models"
	"github.com/aeckard87/WornOut/restapi/operations/users"
)

// CreateUser returns User
func CreateUser(params users.CreateUserParams) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	db.Create(&params.Body)

	var user model.User

	db.Where("username = ?", params.Body.Username).Find(&user)
	return user
}

// UpdateUser returns User
func UpdateUser(params users.UpdateUserParams) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user model.User

	db.Model(&user).Where("id = ?", params.ID).Updates(params.Body)
	// fmt.Println("UPDATED", user)
	user.ID = params.ID
	return user

}

// DeleteUser returns empty User
func DeleteUser(params users.DeleteUserParams) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user model.User
	user.ID = params.ID

	db.Delete(&user)

	return user

}

// DeleteUsers returns empty User
func DeleteUsers(params users.DeleteUsersParams) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user model.User

	db.Delete(&user)

	return user

}

// GetUser returns User given User.ID
func GetUser(params users.GetUserParams) model.User {
	db := dbpkg.Connect()
	defer db.Close()

	var user model.User
	user.ID = params.ID

	db.First(&user)

	return user

}

// GetUsers returns Users
func GetUsers(params users.GetUsersParams) model.Users {
	db := dbpkg.Connect()
	defer db.Close()

	var users model.Users

	db.Find(&users)

	return users

}
