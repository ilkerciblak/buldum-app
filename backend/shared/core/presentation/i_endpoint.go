package presentation

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) (any, domain.IApplicationError)
}

func GenerateHandlerFunc(e IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := e.HandleRequest(w, r)
		if err != nil {

			// RespondWithErrorJson(w, &domain.ApplicationException{
			// 	Message: err.Error(),
			// 	Code:    400,
			// })

			RespondWithProblemDetails(w, err)
			return
		}
		RespondWithJSON(w, data)

		// TODO RESPONDWITHJSON ETC
	}
}

func RespondWithJSON(w http.ResponseWriter, payload any) {
	w.Header().Set("content-type", "application/json")

	if payload != nil {
		data, err := jsonmapper.EncodeObjectToJson(payload)
		if err != nil {
			log.Printf("Something wrong with RespondWithJson Method with: %v payload, %v err", payload, err)
			panic(err)
		}
		_, _ = w.Write(data)
	}

	w.WriteHeader(http.StatusOK)

}

func RespondWithErrorJson(w http.ResponseWriter, appError domain.IApplicationError) {
	w.Header().Set("content-type", "application/problem")
	// TODO: LOGGING
	// TODO: ApiProblem to ProblemDetails
	payload, err := jsonmapper.EncodeObjectToJson(appError)

	if err != nil {
		log.Printf("Something wrong with Respond With Error Json with : %v error struct, %v err", appError, err)
		RespondWithErrorJson(w, &domain.InternalServerError)

	}

	w.WriteHeader(appError.GetCode())
	_, _ = w.Write(payload)

}

func RespondWithProblemDetails(w http.ResponseWriter, appError domain.IApplicationError) {
	w.Header().Set("Content-Type", "application/problem+json")

	problemDetails := ToProblemDetails(appError)

	payload, err := jsonmapper.EncodeObjectToJson(problemDetails)
	if err != nil {
		log.Printf("Something Wrong With RRespondWithProblemDetails func: %v problemdetails, %v err", problemDetails, err)
		RespondWithErrorJson(w, &domain.InternalServerError)
	}

	w.WriteHeader(problemDetails.Status)
	_, _ = w.Write(payload)

}
