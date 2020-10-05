package router

import (
	"github.com/gorilla/mux"
	"github.com/lulumel0n/ham-exam-prep-go/server/middleware"
)

// Router for routing
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newq", middleware.GetQuestion).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getq", middleware.GetQuestionV2).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/exam", middleware.GetSimExam).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/gettitles", middleware.GetTitles).Methods("GET", "OPTIONS")

	router.HandleFunc("/image", middleware.ReturnImage).Methods("GET", "OPTIONS")
	return router
}
