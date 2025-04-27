package service

import (
	"github.com/llmons/havefun/internal/entity"
)

type UndercoverRepo interface {
	GetAllWordPairs() ([]*entity.WordPair, error)
}

type UndercoverService struct {
	repo UndercoverRepo
}

func NewUndercoverService(repo UndercoverRepo) *UndercoverService {
	return &UndercoverService{
		repo: repo,
	}
}

func (s *UndercoverService) GetAllWordPairs() ([]*entity.WordPair, error) {
	wordPairList, err := s.repo.GetAllWordPairs()
	if err != nil {
		return nil, err
	}
	return wordPairList, nil
}
