package repo

import (
	"github.com/llmons/who-is-the-undercover/internal/base/data"
	"github.com/llmons/who-is-the-undercover/internal/entity"
	"github.com/llmons/who-is-the-undercover/internal/service"
)

type UndercoverRepo struct {
	data *data.Data
}

func NewUndercoverRepo(data *data.Data) service.UndercoverRepo {
	return &UndercoverRepo{
		data: data,
	}
}

func (repo *UndercoverRepo) GetAllWordPairs() ([]*entity.WordPair, error) {
	wordPairList := make([]*entity.WordPair, 0)
	err := repo.data.DB.Find(&wordPairList)
	if err != nil {
		return nil, err
	}
	return wordPairList, nil
}
