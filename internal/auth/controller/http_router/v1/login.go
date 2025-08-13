package v1

import (
	"fmt"
	"net/http"
	"no_api/internal/auth/dto"

	"github.com/go-chi/render"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	input := dto.Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		IP:       r.RemoteAddr,
	}

	err := input.Validate()
	if err != nil {
		fmt.Println("validate error", err)
		http.Error(w, "validate error", http.StatusBadRequest)
		return
	}

	output, err := h.usecase.Login(r.Context(), input)

	if err != nil {
		fmt.Println("login error", err)
		http.Error(w, "login error", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, output)
}
