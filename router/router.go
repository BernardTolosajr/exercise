package router

import (
	"fmt"
	"net/http"

	"github.com/exercise/db"
	"github.com/exercise/handler"
	"github.com/exercise/repositories"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
)

// New create instance of mux router with mongoDB
func New(mongoDB *db.MongoDB) *mux.Router {
	r := mux.NewRouter()

	// TODO move this to factory
	organizationService := services.NewOrganization(&repositories.OrganizationRepository{
		MongoDB: mongoDB,
	})

	commentService := services.NewCommentService(&repositories.CommentsRepository{
		MongoDB: mongoDB,
	})

	memberService := services.NewMemberService(&repositories.MembersRepository{
		MongoDB: mongoDB,
	})

	// organization handler
	r.Handle("/orgs", handler.PostOrganizationHandler(organizationService)).Methods("POST")
	r.Handle("/orgs/{name}/comments", handler.GetCommentsHandler(commentService)).Methods("GET")
	r.Handle("/orgs/{name}/comments", handler.PostCommentsHandler(commentService)).Methods("POST")
	r.Handle("/orgs/{name}/comments", handler.DeleteCommentsHandler(commentService)).Methods("DELETE")
	// members handler
	r.Handle("/orgs/{name}/members", handler.PostMemberHandler(memberService)).Methods("POST")
	r.Handle("/orgs/{name}/members", handler.GetMembersHandler(memberService)).Methods("GET")
	// ping
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "pong")
	})

	return r
}
