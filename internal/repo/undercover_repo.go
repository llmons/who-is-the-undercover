package repo

import "github.com/llmons/havefun/internal/base/data"

type UndercoverRepo struct {
	data.Data
}

func NewUndercoverRepo(data data.Data) *UndercoverRepo {
	return &UndercoverRepo{
		Data: data,
	}
}

func(repo *UndercoverRepo) GetAllWords() ([]*Un, error) {
	var entities []*WITUEntity
	err := repo.DB.Find(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}