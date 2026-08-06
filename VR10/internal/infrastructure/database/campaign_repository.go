package database

import "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) error {
	// Implementation for saving campaign
	r.campaigns = append(r.campaigns, *campaign)
	return nil
}

func (r *CampaignRepository) Get() []campaign.Campaign {
	return r.campaigns
}
