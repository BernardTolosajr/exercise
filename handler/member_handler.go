package handler

import (
	"encoding/json"
	"net/http"

	"github.com/exercise/models"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Member request
//	Login - The member username
//  AvatarUrl - The member avatar url
type MemberRequest struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

type MemberResponse struct {
	Id      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type MembersResponse struct {
	Members []*services.MemberView `json:"members"`
	Message string                 `json:"message,omitempty"`
}

// TODO move the handler to struct fo better structure
func PostMemberHandler(service services.IMembersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := MemberRequest{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Errorf("error on parsing body %v", err)
			panic(err)
		}

		vars := mux.Vars(r)
		org := vars["name"]

		id, err := service.Create(&models.Member{
			Org:       org,
			AvatarUrl: request.AvatarUrl,
			Login:     request.Login,
		})

		w.Header().Set("Content-Type", "application/json")

		response := &MemberResponse{}

		if err != nil {
			log.Errorf("error on creating member %v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		if id == "" {
			w.WriteHeader(http.StatusServiceUnavailable)
			response.Message = "Something went wrong on creating member."
		}

		response.Id = id
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetMembersHandler(service services.IMembersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the query path
		vars := mux.Vars(r)
		org := vars["name"]

		results, err := service.GetAllBy(org)

		response := &MembersResponse{}

		if err != nil {
			log.Errorf("error on fetching members %v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
		}

		w.WriteHeader(http.StatusOK)
		response.Members = results
		json.NewEncoder(w).Encode(response)
	}
}
