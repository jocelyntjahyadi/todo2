package main

import (
	"log"
	"time"

	"github.com/jocelyntjahyadi/todo2/model"
	"github.com/jocelyntjahyadi/todo2/storage"
)

func main() {

	//initialize
	// var memStore = storage.NewMemory()
	var memStore = storage.GetStorage(storage.StorageTypeDatabase)

	// CREATE //
	if err := memStore.Create(model.Todo{
		ID:          1,
		Title:       "First",
		Description: "First Todo",
		CreatedAt:   time.Now(),
	}); err != nil {
		log.Fatal(err)
	}
	log.Println("todo created")

	// DETAIL //
	obj, err := memStore.Detail(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d: %s\n", obj.ID, obj.Title)

	// LIST //
	list, err := memStore.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range list {
		log.Printf("ID : %d , DESC : %s \n", val.ID, val.Description)
	}

}
