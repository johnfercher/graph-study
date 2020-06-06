package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	models2 "github.com/johnfercher/graph-study/russiandoll/models"
	"math/rand"
	"net/http"
	"strconv"
)

type RussianDollController struct {
}

func (b RussianDollController) Bind(router *mux.Router) {
	router.HandleFunc("/api/russian-doll", b.GetAll).Methods("GET")
	router.HandleFunc("/api/russian-doll/{id}", b.GetById).Methods("GET")
	router.HandleFunc("/api/russian-doll", b.Create).Methods("POST")
	router.HandleFunc("/api/russian-doll/{id}", b.Update).Methods("PUT")
	router.HandleFunc("/api/russian-doll/{id}", b.Delete).Methods("DELETE")
}

func (b RussianDollController) Delete(writer http.ResponseWriter, request *http.Request) {
	var russianDolls []models2.RussianDoll

	russianDolls = append(russianDolls, models2.RussianDoll{Id: "1", Name: "russianDoll1", ParentId: ""})
	russianDolls = append(russianDolls, models2.RussianDoll{Id: "2", Name: "russianDoll2", ParentId: "1"})

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range russianDolls {
		if item.Id == params["id"] {
			russianDolls = append(russianDolls[:index], russianDolls[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(nil)
}

func (b RussianDollController) Update(writer http.ResponseWriter, request *http.Request) {
	var russianDolls []models2.RussianDoll

	russianDolls = append(russianDolls, models2.RussianDoll{Id: "1", Name: "russianDoll1", ParentId: ""})
	russianDolls = append(russianDolls, models2.RussianDoll{Id: "2", Name: "russianDoll2", ParentId: "1"})

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range russianDolls {
		if item.Id == params["id"] {
			var book models2.RussianDoll
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.Id = params["id"]
			russianDolls[index] = book
			json.NewEncoder(writer).Encode(book)
			return
		}
	}

	json.NewEncoder(writer).Encode(&models2.RussianDoll{})
}

func (b RussianDollController) Create(writer http.ResponseWriter, request *http.Request) {
	var russianDolls []models2.RussianDoll

	russianDolls = append(russianDolls, models2.RussianDoll{Id: "1", Name: "russianDoll1", ParentId: ""})
	russianDolls = append(russianDolls, models2.RussianDoll{Id: "2", Name: "russianDoll2", ParentId: "1"})

	writer.Header().Set("Content-Type", "application/json")
	var russianDoll models2.RussianDoll
	_ = json.NewDecoder(request.Body).Decode(&russianDoll)
	russianDoll.Id = strconv.Itoa(rand.Intn(1000000))
	russianDolls = append(russianDolls, russianDoll)
	json.NewEncoder(writer).Encode(russianDoll)
}

func (b RussianDollController) GetById(writer http.ResponseWriter, request *http.Request) {
	var russianDolls []models2.RussianDoll

	russianDolls = append(russianDolls, models2.RussianDoll{Id: "1", Name: "russianDoll1", ParentId: ""})
	russianDolls = append(russianDolls, models2.RussianDoll{Id: "2", Name: "russianDoll2", ParentId: "1"})

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range russianDolls {
		if item.Id == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&models2.RussianDoll{})
}

func (c RussianDollController) GetAll(writer http.ResponseWriter, request *http.Request) {
	var russianDolls []models2.RussianDoll

	russianDolls = append(russianDolls, models2.RussianDoll{Id: "1", Name: "russianDoll1", ParentId: ""})
	russianDolls = append(russianDolls, models2.RussianDoll{Id: "2", Name: "russianDoll2", ParentId: "1"})

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(russianDolls)
}
