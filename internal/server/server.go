package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sldt-test-task/internal/models"
	"sldt-test-task/internal/service"
)

type Server struct {
	Router  *mux.Router
	Service *service.VerificationService
}

func NewServer(r *mux.Router, service *service.VerificationService) *Server {
	return &Server{Router: r, Service: service}
}

func (s *Server) RunServer() {
	s.registerRoutes()
	port := ":8080"

	log.Printf("start HTTP server at %s port", port)
	log.Fatal(http.ListenAndServe(port, s.Router))
}

func (s *Server) VerifyCreditCards(w http.ResponseWriter, r *http.Request) {

	var creditCard models.CreditCard

	err := json.NewDecoder(r.Body).Decode(&creditCard)
	if err != nil {
		log.Printf("Unable to get credit card info from request: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.Service.Verify(creditCard)
	if err != nil {
		log.Printf("Invalid credit card info from request: %v\n", err)
		s.sendErrorResponse(err, w)
		return
	}

	s.sendResponse(w)
}
