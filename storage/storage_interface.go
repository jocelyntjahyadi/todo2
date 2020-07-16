package storage

import(
	"github.com/jocelyntjahyadi/todo2/model"
)

//Create untuk membuat todo baru, disimpan dalam storagenya. dia menerima model Todo, dan return error, kalau tidak ada berarti return nil.
//Detail todonya diambil. berdasarkan id. dia return model todo atau error.
//Delete 
type Storage interface{
	Create(obj model.Todo) error //menyimpan ke dalam databse. insert
	Detail(id int32) (model.Todo, error) //mengambil objek todo berdasarkan id
	Delete(id int32) error // me remove todo berdasarkan id

	List()([]model.Todo, error)
}