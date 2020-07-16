package main

import(
	"github.com/jocelyntjahyadi/todo2/model"
	"github.com/jocelyntjahyadi/todo2/storage"
	"log"
	"time"
)

func main(){

	//inisialisasi
	// var memStore = storage.NewMemory()
	var memStore = storage.GetStorage(storage.StorageTypeMock)

	//create
	if err := memStore.Create(model.Todo{
		ID : 1,
		Title : "First",
		Description : "First Todo",
		CreatedAt: time.Now(),
	}); err != nil{ //kalau error log fatal.
		log.Fatal(err)
	}
	log.Println("todo created")

	// detail
	obj, err := memStore.Detail(1) //ini di deklarasi lagi karna err di atas di deklarasi dalam err aja.
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("%d: %s\n",obj.ID,obj.Title)

	// // MOCK
	// var memMock = storage.Mock()
	// obj, err = memMock.Detail(3) //ini di deklarasi lagi karna err di atas di deklarasi dalam err aja.
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// log.Printf("%d: %s\n",obj.ID,obj.Title)

	list, err := memStore.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range list {
		log.Printf("ID : %d , DESC : %s \n", val.ID, val.Description)
	}



	
}