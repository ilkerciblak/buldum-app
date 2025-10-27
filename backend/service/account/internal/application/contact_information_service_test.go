package application_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/shared/logging"
	"github.com/stretchr/testify/assert"
)

var tingList = map[uuid.UUID]*model.ContactInformation{
	uuid.MustParse("11111111-1111-1111-1111-111111111111"): {
		Id:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:      uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		Type:        model.Email,
		ContactInfo: "test@example.com",
		IsArchived:  false,
		Publicity:   true,
	},
	uuid.MustParse("22222222-2222-2222-2222-222222222222"): {
		Id:          uuid.MustParse("22222222-2222-2222-2222-222222222222"),
		UserID:      uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		Type:        model.PhoneNumber,
		ContactInfo: "+1234567890",
		IsArchived:  false,
		Publicity:   false,
	},
	uuid.MustParse("33333333-3333-3333-3333-333333333333"): {
		Id:          uuid.MustParse("33333333-3333-3333-3333-333333333333"),
		UserID:      uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de"),
		Type:        model.Email,
		ContactInfo: "test2@example.com",
		IsArchived:  true,
		Publicity:   true,
	},
	uuid.MustParse("44444444-4444-4444-4444-444444444444"): {
		Id:          uuid.MustParse("44444444-4444-4444-4444-444444444444"),
		UserID:      uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338"),
		Type:        model.PhoneNumber,
		ContactInfo: "+0987654321",
		IsArchived:  true,
		Publicity:   false,
	},
}

var contactInformationService = application.NewContactInformationService(
	&mockRepository{},
	logging.NewSlogger(
		logging.LoggerOptions{
			MinLevel:    logging.DEBUG,
			JsonLogging: true,
			LoggingRate: 1,
		},
	),
)

type mockRepository struct {
}

func (m *mockRepository) GetAllByAccountId(accountId uuid.UUID, ctx context.Context) []*model.ContactInformation {
	var result []*model.ContactInformation
	for _, ci := range tingList {
		if ci.UserID == accountId {
			result = append(result, ci)
		}
	}
	return result
}
func (m *mockRepository) GetByAccountAndType(accountId uuid.UUID, t model.ContactInformationType, ctx context.Context) (*model.ContactInformation, error) {
	for _, ci := range tingList {
		if ci.UserID == accountId && ci.Type == t {
			return ci, nil
		}
	}
	return nil, nil
}
func (m *mockRepository) CreateCI(model model.ContactInformation, ctx context.Context) error {
	tingList[model.Id] = &model
	return nil
}
func (m *mockRepository) UpdateCI(id uuid.UUID, updated model.ContactInformation, ctx context.Context) error {
	if _, exists := tingList[id]; !exists {
		return nil
	}
	tingList[id] = &updated
	return nil
}
func (m *mockRepository) ArchiveCI(id uuid.UUID, ctx context.Context) error {
	if ci, exists := tingList[id]; exists {
		ci.IsArchived = true
	}
	return nil
}

func TestApplicationLayer__GetContactInformation(t *testing.T) {

	// generate cases and test then O(n)
	accounts := []uuid.UUID{
		uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de"),
		uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338"),
	}
	for _, account := range accounts {
		t.Run(account.String(), func(t *testing.T) {
			contactInfo, _ := contactInformationService.GetSingleByAccountByType(account, model.Email.String(), context.Background())
			assert.NotNil(t, contactInfo)
		})
	}

}
func TestApplicationLayer__GetContactInformations(t *testing.T) {

	// generate cases and test then O(n)
	accounts := []uuid.UUID{
		uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de"),
		uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338"),
	}
	for _, account := range accounts {
		t.Run(account.String(), func(t *testing.T) {
			contactInfos := contactInformationService.GetAllByAccount(account, context.Background())
			assert.NotNil(t, contactInfos)
		})
	}
}
func TestApplicationLayer__CreateContactInformation(t *testing.T) {
	// generate cases and test then O(n)
	accounts := []uuid.UUID{
		uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de"),
		uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338"),
	}
	for _, account := range accounts {
		t.Run(account.String(), func(t *testing.T) {
			newCI := dto.ContactInformationCreateDTO{

				AccountID:   account,
				Type:        model.Email.String(),
				ContactInfo: "test@example.com",
				Publicity:   false,
			}
			err := contactInformationService.Create(newCI, context.Background())
			assert.NoError(t, err)
		})
	}
}
func TestApplicationLayer__UpdateContactInformation(t *testing.T) {
	// generate cases and test then O(n)
	accounts := []uuid.UUID{
		uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de"),
		uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338"),
	}
	for _, account := range accounts {
		t.Run(account.String(), func(t *testing.T) {
			updatedDto := dto.ContactInformationUpdateDTO{

				AccountID:   account,
				ContactInfo: "updated@example.com",
			}

			err := contactInformationService.Update(account, updatedDto, context.Background())
			assert.NoError(t, err)
		})
	}
}
func TestApplicationLayer__ArhiveContactInformation(t *testing.T) {
	// generate cases and test then O(n)
	accounts := []uuid.UUID{
		uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5"),
		uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de"),
		uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338"),
	}
	for _, account := range accounts {
		t.Run(account.String(), func(t *testing.T) {
			var ciID uuid.UUID
			for _, ci := range tingList {
				if ci.UserID == account {
					ciID = ci.Id
					break
				}
			}
			err := contactInformationService.Archive(ciID, context.Background())
			assert.NoError(t, err)
		})
	}
}
