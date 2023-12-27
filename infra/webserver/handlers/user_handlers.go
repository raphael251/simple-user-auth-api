package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/raphael251/simple-user-auth-api/infra/database"
	"github.com/raphael251/simple-user-auth-api/internal/dto"
	"github.com/raphael251/simple-user-auth-api/internal/entity"
	"github.com/raphael251/simple-user-auth-api/pkg/utils"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user dto.CreateUserInput

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := entity.NewUser(user.Email, user.Password)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMessage := utils.Error{Message: "bad request"}
			json.NewEncoder(w).Encode(errMessage)
			return
		}

		foundUser, _ := database.FindUserByEmail(db, user.Email)

		if foundUser != nil {
			w.WriteHeader(http.StatusBadRequest)
			errMessage := utils.Error{Message: "this email is already being used"}
			json.NewEncoder(w).Encode(errMessage)
		}

		err = database.InsertUser(db, u)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
