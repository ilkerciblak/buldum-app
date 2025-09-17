package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type CreateAccountEndPoint struct{}

func (c CreateAccountEndPoint) Path() string {
	return "/account"
}

func (c CreateAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {

	com, err := jsonmapper.DecodeRequestBody[command.CreateAccountCommand](r)
	if err != nil {
		return corepresentation.ApiResult[any]{}, coredomain.InternalServerError.WithMessage(err.Error())
	}

	if err := com.Handler(&MockAccountRepository{}, r.Context()); err != nil {
		return corepresentation.ApiResult[any]{}, err
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusCreated,
	}, nil

}

type MockAccountRepository struct {
}

func (m MockAccountRepository) GetById(userId uuid.UUID) (*model.Profile, error) {
	return model.NewProfile("ilkerciblak", "url"), nil
}
func (m MockAccountRepository) GetAll() ([]*model.Profile, error) {
	return []*model.Profile{}, nil
}
func (m MockAccountRepository) Create(p *model.Profile) error {
	return nil
}
func (m MockAccountRepository) Update(userId uuid.UUID, p *model.Profile) error {
	return nil
}
func (m MockAccountRepository) Delete(userId uuid.UUID) error {
	return nil
}
