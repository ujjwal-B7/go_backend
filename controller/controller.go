package controller

import (
	"backend/model"
	"backend/utils"
	"encoding/json"
	"io"
	"net/http"
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
		utils.JSON(w, http.StatusInsufficientStorage, err)
		return
	}

	utils.JSON(w, http.StatusOK, "Futsal registered.")
}
