package repo

import (
	"github.com/llmons/who-is-the-undercover/internal/base/data"
	"github.com/llmons/who-is-the-undercover/internal/entity"
)

type UndercoverRepo struct {
	data *data.Data
}

func NewUndercoverRepo(data *data.Data) *UndercoverRepo {
	return &UndercoverRepo{
		data: data,
	}
}

func (repo *UndercoverRepo) GetAllWordPairs() ([]*entity.WordPair, error) {
	wordPairList := make([]*entity.WordPair, 0)
	repo.data.DB.Find(&wordPairList)
	return wordPairList, nil
}
