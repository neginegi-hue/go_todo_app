package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/neginegi-hue/go_todo_app/entity"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validater *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrRespponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := ru.Validater.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrRespponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := ru.Validater.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrRespponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
	}

	u, err := ru.Service.RegisterUser(ctx, b.Name, b.Password, b.Role)
	if err != nil {
		RespondJSON(ctx, w, &ErrRespponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
	}

	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
