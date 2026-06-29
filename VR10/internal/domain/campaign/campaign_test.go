package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "Body HI!"
	contacts = []string{"email1@e.com", "emaile2@e.com"}
	fake     = faker.New()
)

func TestNewCampaign(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(t, campaign.Name, "Campaign x")
	assert.Equal(t, campaign.Content, content)
	assert.Equal(t, len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_ID_IsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}

func Test_NewCampaign_Created_IsNotNill(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(t, campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal(t, "name is required", err.Error())
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("Name", content, contacts)

	assert.Equal(t, "name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal(t, "name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, contacts)

	assert.Equal("createdby is invalid", err.Error())
}
