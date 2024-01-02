package server

import (
	"github.com/SheepJenga/crescendo/utils"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusInternalServerError, "This is an error")
}
