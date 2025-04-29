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

func (repo *UndercoverRepo) GetRandomWordPair() (*entity.WordPair, error) {
	wordPair := &entity.WordPair{}
	has, err := repo.data.DB.OrderBy("rand() desc").Limit(1).Get(wordPair)
	if !has || err != nil {
		return nil, err
	}
	return wordPair, nil
}
