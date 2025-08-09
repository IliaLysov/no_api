package v1

import (
	"fmt"
	"net/http"
	"no_api/internal/auth/dto"

	"github.com/go-chi/render"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateUserInput{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err := input.Validate()
	if err != nil {
		fmt.Println("validate error", err)
		http.Error(w, "validate error", http.StatusBadRequest)
		return
	}

	output, err := h.usecase.CreateUser(r.Context(), input)
	if err != nil {
		fmt.Println("error to create user", err)
		http.Error(w, "error to create user", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, output)
}
