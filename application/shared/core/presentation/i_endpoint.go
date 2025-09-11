package presentation

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) (any, error)
}

func GenerateHandlerFunc(e IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := e.HandleRequest(w, r)
		if err != nil {
			log.Printf("Err nil degil")
			RespondWithErrorJson(w, &domain.ApplicationError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		RespondWithJSON(w, data)

		// TODO RESPONDWITHJSON ETC
	}
}

func RespondWithJSON(w http.ResponseWriter, payload any) {
	w.Header().Set("content-type", "application/json")
	data, err := jsonmapper.EncodeObjectToJson(payload)
	if err != nil {
		log.Printf("Something wrong with RespondWithJson Method with: %v payload, %v err", payload, err)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)

}

func RespondWithErrorJson(w http.ResponseWriter, appError domain.IApplicationError) {
	w.Header().Set("content-type", "application/json")
	// TODO: LOGGING
	// TODO: ApiProblem to ProblemDetails
	payload, err := jsonmapper.EncodeObjectToJson(appError)
	if err != nil {
		log.Printf("Something wrong with Respond With Error Json with : %v error struct, %v err", appError, err)
		panic(err)
	}
	w.WriteHeader(appError.GetCode())
	_, _ = w.Write(payload)

}
