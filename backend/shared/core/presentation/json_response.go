package corepresentation

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

func RespondWithJSON(w http.ResponseWriter, payload any) {
	w.Header().Set("content-type", "application/json")

	if payload != nil {
		data, err := jsonmapper.EncodeObjectToJson(payload)
		if err != nil {
			log.Printf("Something wrong with RespondWithJson Method with: %v payload, %v err", payload, err)
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}

}

func RespondWithErrorJson(w http.ResponseWriter, appError coredomain.IApplicationError) {
	w.Header().Set("content-type", "application/problem")
	// TODO: LOGGING
	// TODO: ApiProblem to ProblemDetails
	payload, err := jsonmapper.EncodeObjectToJson(appError)

	if err != nil {
		log.Printf("Something wrong with Respond With Error Json with : %v error struct, %v err", appError, err)
		RespondWithErrorJson(w, &coredomain.InternalServerError)

	}

	w.WriteHeader(appError.GetCode())
	_, _ = w.Write(payload)

}

func RespondWithProblemDetails(w http.ResponseWriter, appError error) {
	w.Header().Set("Content-Type", "application/problem+json")

	problemDetails := ToProblemDetails(appError)

	payload, err := jsonmapper.EncodeObjectToJson(problemDetails)
	if err != nil {
		log.Printf("Something Wrong With RespondWithProblemDetails func: %v problemdetails, %v err", problemDetails, err)
		RespondWithErrorJson(w, &coredomain.InternalServerError)
	}

	w.WriteHeader(problemDetails.Status)
	_, _ = w.Write(payload)

}
