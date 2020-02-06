package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/exercise/models"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
)

type CommentRequest struct {
	Comment string `json:"comment"`
}

type CommentResponse struct {
	Id      string `json:"id,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func PostCommentsHandler(service services.ICommentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		request := CommentRequest{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Printf("error parsing request %v", err)
			panic(err)
		}

		// get the query path
		vars := mux.Vars(r)
		org := vars["name"]

		id, err := service.Create(&models.Comment{
			Org:     org,
			Comment: request.Comment,
		})

		response := &CommentResponse{}

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		if id == "" {
			w.WriteHeader(http.StatusServiceUnavailable)
			response.Message = "Something went wrong on creating comment."
		}

		response.Id = id

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
