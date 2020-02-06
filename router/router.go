package router

import (
	"github.com/exercise/db"
	"github.com/exercise/handler"
	"github.com/exercise/repositories"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
)

// New create instance of mux router with mongoDB
func New(mongoDB *db.MongoDB) *mux.Router {
	r := mux.NewRouter()

	organizationService := services.NewOrganization(&repositories.OrganizationRepository{
		MongoDB: mongoDB,
	})

	commentService := services.NewCommentService(&repositories.CommentsRepository{
		MongoDB: mongoDB,
	})

	r.Handle("/orgs", handler.PostOrganizationHandler(organizationService)).Methods("POST")
	r.Handle("/orgs/{name}/comments", handler.PostCommentsHandler(commentService)).Methods("POST")

	return r
}

// New create instance of mux router with mongoDB
func NewWithService(commentService services.ICommentService) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/orgs/{name}/comments", handler.PostCommentsHandler(commentService)).Methods("POST")

	return r
}
