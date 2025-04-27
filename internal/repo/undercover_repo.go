package repo

import (
	"github.com/llmons/havefun/internal/base/data"
	"github.com/llmons/havefun/internal/entity"
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
