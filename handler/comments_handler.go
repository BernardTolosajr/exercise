package handler

import (
	"encoding/json"
	"net/http"

	"github.com/exercise/models"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Comment request
//	MemberId - The member Id
//  Comment - The comment
type CommentRequest struct {
	MemberId string `json:"member_id"`
	Comment  string `json:"comment"`
}

type CommentResponse struct {
	Id      string `json:"id,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type DeleteCommentResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type CommentsResponse struct {
	Comments []*services.Comment `json:"comments"`
	Message  string              `json:"message,omitempty"`
}

// TODO move the handler to struct fo better structure
func GetCommentsHandler(service services.ICommentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the query path
		vars := mux.Vars(r)
		org := vars["name"]

		results, err := service.GetAllBy(org)

		response := &CommentsResponse{}

		if err != nil {
			log.Errorf("error on fetching comment %v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
		}

		w.WriteHeader(http.StatusOK)
		response.Comments = results
		json.NewEncoder(w).Encode(response)
	}
}

// Delete Comment Handler handle deleting of comments
func DeleteCommentsHandler(service services.ICommentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// get the query path
		vars := mux.Vars(r)
		org := vars["name"]

		result, err := service.DeleteAll(org)

		response := &DeleteCommentResponse{}

		if err != nil {
			log.Errorf("error on creating comment %v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
		}

		if result == 0 {
			response.Message = "No changes found."
			json.NewEncoder(w).Encode(response)
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Post Comment Handler handle creating new comments
func PostCommentsHandler(service services.ICommentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		request := CommentRequest{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Errorf("error on parsing body %v", err)
			panic(err)
		}

		// get the query path
		vars := mux.Vars(r)
		org := vars["name"]

		id, err := service.Create(&models.Comment{
			Org:      org,
			Comment:  request.Comment,
			MemberId: request.MemberId,
		})

		response := &CommentResponse{}

		if err != nil {
			log.Errorf("error on creating comment %v", err)
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
