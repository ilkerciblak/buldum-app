package account_db

import "github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"

func ContactInformationDBModelToModel(a AccountContactInformation) *model.ContactInformation {
	return &model.ContactInformation{
		Id:          a.ID,
		UserID:      a.ProfileID.UUID,
		Type:        model.ContactInformationTypeFromString(a.ContactInformationType.String),
		Publicity:   a.IsPublic.Bool,
		ContactInfo: a.ContactInformation.String,
		IsArchived:  a.IsArchived.Bool,
		CreatedAt:   a.CreatedAt.Time,
		UpdatedAt:   a.UpdatedAt.Time,
		DeletedAt:   a.DeletedAt.Time,
	}
}

func ContactInformationDbModelListToModelList(a []AccountContactInformation) []*model.ContactInformation {
	res := make([]*model.ContactInformation, len(a))

	for i, m := range a {
		res[i] = ContactInformationDBModelToModel(m)
	}

	return res
}
