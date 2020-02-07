package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/exercise/models"
	"github.com/exercise/services"
)

// Member request
//	Login - The Organization username
//  Admin - The login of the user who will manage this organization
// 	ProfileName - The organization's display name.
type OrganizationRequest struct {
	Login       string `json:"login"`
	ProfileName string `json:"profile_name"`
	Admin       string `json:"admin"`
}

type OrganizationResponse struct {
	Id      string `json:"id,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// TODO move the handler to struct fo better structure
// OrganizationHandler handles the creation of new organization
func PostOrganizationHandler(service services.IOrganizationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		request := OrganizationRequest{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Printf("error parsing request %v", err)
			panic(err)
		}

		id, err := service.Create(&models.Organization{
			Login:       request.Login,
			ProfileName: request.ProfileName,
			Admin:       request.Admin,
		})

		response := &OrganizationResponse{}

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		if id == "" {
			w.WriteHeader(http.StatusServiceUnavailable)
			response.Message = "Something went wrong."
		}

		response.Id = id

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
