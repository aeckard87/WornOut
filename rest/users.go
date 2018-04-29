package controls

import (
	"encoding/json"
	"net/http"

	controller "github.com/aeckard87/WornOut/controllers"
	model "github.com/aeckard87/WornOut/models"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.Unmarshal(GetBodyBytes(r), &user)

	//should have an err response too
	user = controller.CreateUser(user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	// fmt.Fprintf(w, StringifyUser(user))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.Unmarshal(GetBodyBytes(r), &user)
	vars := mux.Vars(r)

	//should have an err response too
	user = controller.GetUser(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	// fmt.Fprintf(w, StringifyUser(user))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	users := controller.GetUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	// fmt.Fprintf(w, StringifyUsers(users))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.Unmarshal(GetBodyBytes(r), &user)
	vars := mux.Vars(r)

	//should have an err response too
	user = controller.UpdateUser(user, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	// fmt.Fprintf(w, StringifyUser(user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.Unmarshal(GetBodyBytes(r), &user)
	vars := mux.Vars(r)

	//should have an err response too
	user = controller.DeleteUser(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	// fmt.Fprintf(w, StringifyUser(user))
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	users := controller.DeleteUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	// fmt.Fprintf(w, StringifyUsers(users))
}
