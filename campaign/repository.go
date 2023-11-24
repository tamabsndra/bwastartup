package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	Delete(campaign Campaign) error
	CreateImage(campaignImages CampaignImages) (CampaignImages, error)
	MarkAllImagesNonPrimary(campaignID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("User").Preload("CampaignImages").First(&campaign, ID).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Delete(campaign Campaign) error {
	err := r.db.Delete(&campaign).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateImage(campaignImages CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&campaignImages).Error
	if err != nil {
		return campaignImages, err
	}

	return campaignImages, nil
}

func (r *repository) MarkAllImagesNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImages{}).Where("campaign_id", campaignID).Update("is_primary", 0).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
