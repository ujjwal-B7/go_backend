package controller

import (
	"backend/model"
	"backend/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) RegisterFutsal(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	futsal := model.Futsal{}

	json.Unmarshal(body, &futsal)

	err = server.UseCase.SaveFutsal(futsal)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, "Futsal registered.")
}

func (server *Server) GetAllFutsals(w http.ResponseWriter, r *http.Request) {

	allFutsals, err := server.UseCase.GetAllFutsals()

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, allFutsals)
}

func (server *Server) UpdateFutsal(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	futsal := model.Futsal{}

	json.Unmarshal(body, &futsal)

	updatedFutsal, err := server.UseCase.UpdateFutsal(id, futsal)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, updatedFutsal)
}

func (server *Server) UpdateFutsalFields(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	futsal := model.Futsal{}

	json.Unmarshal(body, &futsal)

	updatedFutsal, err := server.UseCase.UpdateFutsalFields(id, futsal)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, updatedFutsal)
}

func (server *Server) DeleteFutsal(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err = server.UseCase.DeleteFutsal(id)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, "Futsal deleted successfully.")

}
