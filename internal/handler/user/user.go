package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	service serviceUser
}

func NewUserController(service serviceUser) *Controller {
	return &Controller{service}
}

func (c *Controller) SetIsActive(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
		return
	}

	user, err := c.service.SetIsActive(r.Context(), id, "")

}
