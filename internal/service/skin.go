package service

import (
	"database/sql"
	"log/slog"
	"skin-monkey/internal/entity"
	repository "skin-monkey/internal/repository/postgres"
)

type SkinService struct {
	repository repository.Skin
	log        *slog.Logger
	bot        *BotService
}

func NewSkinService(repo repository.Skin, log *slog.Logger, bot *BotService) *SkinService {
	return &SkinService{
		repository: repo,
		log:        log,
		bot:        bot,
	}
}

func (s *SkinService) CreateSkin(skin *entity.Skin) error {
	if err := skin.Validate(); err != nil {
		return err
	}

	err := s.repository.CreateSkin(skin)
	if err != nil {
		return err
	}

	if skin.New {

	}

	return nil
}

func (s *SkinService) GetSkin(id string) (*entity.Skin, error) {
	return s.repository.GetSkin(id)
}

func (s *SkinService) UpdateSkin(skin *entity.Skin) error {
	existingSkin, err := s.GetSkin(skin.Id)
	if err != nil {
		return err
	}

	existingSkin.Name = skin.Name

	if err = existingSkin.Validate(); err != nil {
		s.log.Error(err.Error())
		return err
	}

	if err = s.repository.UpdateSkin(existingSkin); err != nil {
		return err
	}

	return nil
}

func (s *SkinService) DeleteSkin(id string) error {
	_, err := s.GetSkin(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return repository.ErrSkinNotFound
		}

		return err
	}

	return s.repository.DeleteSkin(id)
}

func (s *SkinService) GetSkinFilter() (*entity.Skin, error) {
	return s.repository.GetSkinsFilter()
}
