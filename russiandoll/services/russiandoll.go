package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/johnfercher/graph-study/russiandoll/models"
	"github.com/johnfercher/graph-study/russiandoll/repositories"
	"net/http"
)

type RussianDollService struct {
	repository *repositories.RussianDollRepository
}

func NewRussianDollService(repository *repositories.RussianDollRepository) *RussianDollService {
	return &RussianDollService{
		repository: repository,
	}
}

func (self *RussianDollService) RemoveFromRussianDoll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.repository.RemoveFromRussianDoll(ctx, params["parent_id"], params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}

func (self *RussianDollService) AddToRussianDoll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.repository.AddToRussianDoll(ctx, params["parent_id"], params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}

func (self *RussianDollService) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	russianDoll := &models.RussianDoll{}
	err := json.NewDecoder(request.Body).Decode(russianDoll)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	params := mux.Vars(request)
	russianDoll.Id = params["id"]

	err = self.repository.Update(ctx, russianDoll)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(russianDoll)
}

func (self *RussianDollService) Create(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	russianDoll := &models.RussianDoll{}
	err := json.NewDecoder(request.Body).Decode(russianDoll)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	id, _ := uuid.NewRandom()
	russianDoll.Id = id.String()

	err = self.repository.Create(ctx, russianDoll)
	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(russianDoll)
}

func (self *RussianDollService) GetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	russianDoll, err := self.repository.GetById(ctx, params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(russianDoll)
}

func (self *RussianDollService) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	ctx := context.TODO()

	russianDolls, err := self.repository.GetAll(ctx)

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(russianDolls)
}

func (self *RussianDollService) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	params := mux.Vars(request)

	err := self.repository.Delete(ctx, params["id"])

	if err != nil {
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(nil)
}
