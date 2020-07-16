package storage

import (
	"errors"
	"time"

	"github.com/jocelyntjahyadi/todo2/model"
)

type mock struct{}

func (o mock) Create(obj model.Todo) error {
	// o.store[obj.ID] = obj //memasukkan data ke map. o adalah receiver Memeory, ada atribut store. Masukkan ke store dengan idnya key.
	return nil
}

//
func (o mock) Detail(id int32) (model.Todo, error) {

	if id == 1 {
		return model.Todo{ID: 1, Title: "Mock Title", Description: "mocked description", CreatedAt: time.Now()}, nil
	}
	return model.Todo{}, errors.New("Todo tidak ditemukan")
}

func (o mock) Delete(id int32) error {

	return nil
}

func (o mock) List() ([]model.Todo, error) {

	result := []model.Todo{}
	result = append(result, model.Todo{
		ID:          1,
		Title:       "Mock Title",
		Description: "Mock Description",
		CreatedAt:   time.Now(),
	}, model.Todo{
		ID:          2,
		Title:       "Mock Title",
		Description: "Second mock",
		CreatedAt:   time.Now(),
	})
	return result, nil

}
