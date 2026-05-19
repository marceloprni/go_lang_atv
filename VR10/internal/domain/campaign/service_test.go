package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {

	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Teste y",
		Content: "Body",
		Emails:  []string{"teste1@teste.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(t, id)
	assert.Nil(t, err)
}

func Test_Save_Campaign(t *testing.T) {

	service := Service{}
	repositoryMock := new(repositoryMock)
	newCampaign := contract.NewCampaign{
		Name:    "Teste y",
		Content: "Body",
		Emails:  []string{"teste1@teste.com"},
	}

	repositoryMock.On("Save", mock.MatchedBy(
		func(campaign *Campaign) bool {
			if campaign.Name != newCampaign.Name ||
				campaign.Content != newCampaign.Content ||
				len(campaign.Content) != len(newCampaign.Emails) {
				return false
			}
			return true
		},
	)).Return(nil)

	service = Service{Repository: repositoryMock}
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
