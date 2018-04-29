package controls

import (
	"encoding/json"
	"net/http"

	controller "github.com/aeckard87/goServe/controllers"
	model "github.com/aeckard87/goServe/models"
	"github.com/gorilla/mux"
)

func CreateDescriptorByDetail(w http.ResponseWriter, r *http.Request) {
	var descriptor model.Descriptor
	json.Unmarshal(GetBodyBytes(r), &descriptor)
	vars := mux.Vars(r)

	//should have an err response too
	descriptor = controller.CreateDescriptorByDetail(descriptor, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(descriptor)
	// fmt.Fprintf(w, StringifyDesc(descriptor))
}

func GetDescriptor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	descriptor := controller.GetDescriptor(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptor)
	// fmt.Fprintf(w, StringifyDesc(descriptor))
}

func GetDescriptors(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	descriptors := controller.GetDescriptors()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptors)
	// fmt.Fprintf(w, StringifyDescs(descriptors))
}

func GetDescriptorsByDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	descriptors := controller.GetDescriptorsByDetail(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptors)
	// fmt.Fprintf(w, StringifyDescs(descriptors))
}

func UpdateDescriptor(w http.ResponseWriter, r *http.Request) {
	var descriptor model.Descriptor
	json.Unmarshal(GetBodyBytes(r), &descriptor)
	vars := mux.Vars(r)

	//should have an err response too
	descriptor = controller.UpdateDescriptor(descriptor, vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptor)
	// fmt.Fprintf(w, StringifyDesc(descriptor))
}

func DeleteDescriptor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	descriptor := controller.DeleteDescriptor(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptor)
	// fmt.Fprintf(w, StringifyDesc(descriptor))
}

func DeleteDescriptors(w http.ResponseWriter, r *http.Request) {
	//should have an err response too
	descriptors := controller.DeleteDescriptors()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptors)
	// fmt.Fprintf(w, StringifyDescs(descriptors))
}

func DeleteDescriptorsByDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//should have an err response too
	descriptors := controller.DeleteDescriptorsByDetail(vars["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(descriptors)
	// fmt.Fprintf(w, StringifyDescs(descriptors))
}
