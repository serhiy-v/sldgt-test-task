package server

import (
	"encoding/json"
	"net/http"
	"sldt-test-task/internal/models"
)

func (s *Server) registerRoutes() {
	s.Router.Use(SetHeaders)
	s.Router.HandleFunc("/verify", s.VerifyCreditCards).Methods("POST")
}

func SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) sendResponse(w http.ResponseWriter) {
	resp := models.ServiceResponse{
		Valid: true,
	}

	json.NewEncoder(w).Encode(resp)

	w.WriteHeader(http.StatusOK)
}

func (s *Server) sendErrorResponse(err error, w http.ResponseWriter) {
	switch err.Error() {
	case models.CardExpError:
		svcErr := models.ServiceError{
			Code:    01,
			Message: err.Error(),
		}

		resp := models.ServiceErrorResponse{
			Valid: false,
			Error: svcErr,
		}

		json.NewEncoder(w).Encode(resp)
	case models.InvCardNumError:
		svcErr := models.ServiceError{
			Code:    02,
			Message: err.Error(),
		}

		resp := models.ServiceErrorResponse{
			Valid: false,
			Error: svcErr,
		}

		json.NewEncoder(w).Encode(resp)
	case models.InvExpMonthError:
		svcErr := models.ServiceError{
			Code:    03,
			Message: err.Error(),
		}

		resp := models.ServiceErrorResponse{
			Valid: false,
			Error: svcErr,
		}

		json.NewEncoder(w).Encode(resp)

		w.WriteHeader(http.StatusOK)
	}

}
