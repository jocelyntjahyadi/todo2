package storage

import(
	"errors"
	"github.com/jocelyntjahyadi/todo2/model"
)

//membuat objek namanya memoery
type memory struct{
	store map[int32]model.Todo
}

//ini optional, tapi kita butuh initiate mapnya.
func newMemory() memory{ //ini constructor
	return memory{
		store:make(map[int32]model.Todo),
	}
}

//Signature functionnya harus sama dengan yang di storage interface.
func (o memory) Create(obj model.Todo)error{
	o.store[obj.ID] = obj //memasukkan data ke map. o adalah receiver Memeory, ada atribut store. Masukkan ke store dengan idnya key.
	return nil
}

//
func (o memory) Detail(id int32)(model.Todo,error){

	//OK NOTATION
	// kalau objeknya ada nialinya, dan ok nilainya true (kalau ga ketemu ok false)
	if obj, ok := o.store[id]; ok{ //untuk ngecek apakah id tertentu exist dalam sebuah array atau map.
		return obj, nil
	}

	// ATAU

	// if obj, ok := o.store[id]; 
	// if ok{ 
	// 	return obj, nil
	// }

    return model.Todo{},errors.New("todo tidak ditemukan")
}

func (o memory) Delete(id int32) error{
	delete(o.store, id)
	return nil
}

func (o memory) List()([]model.Todo, error){
	result := []model.Todo{}
	for _, v := range o.store {
		result = append(result, v)
	}
	return result, nil
}

