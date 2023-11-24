package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	DeleteCampaign(input GetCampaignDetailInput, userID int) error
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImages, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) DeleteCampaign(input GetCampaignDetailInput, userID int) error {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return err
	}

	if campaign.UserID != userID {
		return errors.New("not an owner of the campaign")
	}

	err = s.repository.Delete(campaign)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.Description = inputData.Description
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return campaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImages, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImages{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImages{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImages{}, err
		}
	}

	campaignImages := CampaignImages{}
	campaignImages.CampaignID = input.CampaignID
	campaignImages.IsPrimary = isPrimary
	campaignImages.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImages)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
