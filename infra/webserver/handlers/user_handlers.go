package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/raphael251/simple-user-auth-api/infra/database"
	"github.com/raphael251/simple-user-auth-api/internal/dto"
	"github.com/raphael251/simple-user-auth-api/internal/entity"
	"github.com/raphael251/simple-user-auth-api/pkg/utils"
)

// Create user godoc
// @Summary		 Create user
// Description Create user
// @Tags			 users
// @Accept		 json
// @Produce		 json
// @Param			 request 	body 		 dto.CreateUserInput true "user request"
// @Success		 201
// @Failure		 500		 	{object} utils.Error
// @Router		 /users	[post]
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

		w.WriteHeader(http.StatusCreated)
	}
}

// HandleLogin godoc
// @Summary		 Login
// Description Handles user login, returning a JWT access token
// @Tags			 users
// @Accept		 json
// @Produce		 json
// @Param			 request 			body 			dto.UserLoginInput		true	"user credentials"
// @Success		 200 					{object}	dto.UserLoginOutput
// @Failure		 401					{object}	utils.Error
// @Failure		 500		 			{object}	utils.Error
// @Router		 /users/login	[post]
func HandleLogin(db *sql.DB, jwt *jwtauth.JWTAuth, jwtExpiresIn int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user dto.UserLoginInput

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		foundUser, err := database.FindUserByEmail(db, user.Email)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			errMessage := utils.Error{Message: "Invalid e-mail or password"}
			json.NewEncoder(w).Encode(errMessage)
			return
		}

		isValid := foundUser.ValidatePassword(user.Password)

		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			errMessage := utils.Error{Message: "Invalid e-mail or password"}
			json.NewEncoder(w).Encode(errMessage)
			return
		}

		_, tokenString, _ := jwt.Encode(map[string]interface{}{
			"sub": foundUser.ID.String(),
			"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
		})

		accessToken := dto.UserLoginOutput{AccessToken: tokenString}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accessToken)
	}
}
