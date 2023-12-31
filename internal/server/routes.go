package server

import (
	"go-server/internal/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes - contains all server routes
func (s *server) routes() *httprouter.Router {

	router := httprouter.New()

	router.POST("/login", s.handleLogin)
	router.GET("/timesheets", s.WithAuth(s.handleGetTimesheets))
	router.POST("/timesheet", s.WithAuth(s.handlePostTimesheet))

	// Enable CORS for other methods
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			utils.SetCorsHeaders(w, r)
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	return router
}
