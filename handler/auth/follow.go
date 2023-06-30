package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RipulHandoo/goChat/db"
	"github.com/RipulHandoo/goChat/db/database"
	"github.com/RipulHandoo/goChat/utils"
)


func FollowUser(w http.ResponseWriter, req *http.Request, user database.User) {
	toFollowIDParam := req.URL.Query().Get("id")
    if toFollowIDParam == "" {
        utils.ResponseWithError(w, http.StatusBadRequest, fmt.Errorf("Can not get the id"))
        return
    }

    toFollowID, err := strconv.ParseInt(toFollowIDParam, 10, 64)
    if err != nil {
        utils.ResponseWithError(w, http.StatusBadRequest, err)
        return
    }
	apiConfig := db.DbClient
	userFollowTuple, followerUpdateErr := apiConfig.FollowUser(req.Context(), database.FollowUserParams{
		FollowingID: toFollowID,
		FollowerID:  user.ID,
	})
	if followerUpdateErr != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, followerUpdateErr)
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userFollowTuple)
}

func UnFollowUser(w http.ResponseWriter, req *http.Request, user database.User) {
	toFollowIDParam := req.URL.Query().Get("id")
    if toFollowIDParam == "" {
        utils.ResponseWithError(w, http.StatusBadRequest, fmt.Errorf("Can not get the id"))
        return
    }

    toFollowID, err := strconv.ParseInt(toFollowIDParam, 10, 64)
    if err != nil {
        utils.ResponseWithError(w, http.StatusBadRequest, err)
        return
    }
	apiConfig := db.DbClient
	userFollowTuple, followerUpdateErr := apiConfig.UnfollowUser(req.Context(), database.UnfollowUserParams{
		FollowingID: toFollowID,
		FollowerID:  user.ID,
	})
	if followerUpdateErr != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, followerUpdateErr)
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, userFollowTuple)
}