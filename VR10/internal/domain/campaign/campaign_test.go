package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "Body"
	contacts = []string{"email1@e.com", "emaile2@e.com"}
)

func TestNewCampaign(t *testing.T) {

	campaign, _ := NewCampaign(name, content, contacts)

	//if campaign.ID != "1" {
	//	t.Error("expected 1", campaign.ID)
	//} else if campaign.Name != name {
	//	t.Error("expected correct name", campaign.Name)
	//} else if campaign.Content != content {
	//	t.Error("expected correct content", campaign.Content)
	//} else if len(campaign.Contacts) != len(contacts) {
	//	t.Error("expected correct content", campaign.Content)
	//}

	assert.Equal(t, campaign.Name, "Campaign x")
	assert.Equal(t, campaign.Content, content)
	assert.Equal(t, len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_ID_IsNotNill(t *testing.T) {
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(t, campaign.ID)
}

func Test_NewCampaign_Created_IsNotNill(t *testing.T) {
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(t, campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {

	_, err := NewCampaign("", content, contacts)

	assert.Equal(t, "name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {

	_, err := NewCampaign(name, "", contacts)

	assert.Equal(t, "content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {

	_, err := NewCampaign(name, content, []string{})

	assert.Equal(t, "contacts is required", err.Error())
}
