package storage

import (
	"log"
	"time"

	"github.com/jocelyntjahyadi/todo2/model"

	"database/sql"

	// ini buat manggil postgres
	_ "github.com/lib/pq"
)

type database struct{}

func connection() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", "host=postgres port=5432 user=postgres "+
		"password=password database=db sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func (o database) Create(obj model.Todo) error {

	db, _ := connection()
	_, err := db.Exec("CREATE TABLE todo2 (id int, title text, description text, createdat timestamp) ")

	// _, err := db.Exec("INSERT INTO todo2 (id, title, description, createdat) VALUES "+
	// 	"($1,$2,$3,$4);",
	// 	obj.ID,
	// 	obj.Title,
	// 	obj.Description,
	// 	time.Now())

	return err
}

//
func (o database) Detail(id int32) (model.Todo, error) {

	db, _ := connection()

	rows, err := db.Query("SELECT id, title, description, createdat FROM todo2 WHERE id = $1;", id)
	if err != nil {
		return model.Todo{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var st model.Todo
		err := rows.Scan(&st.ID, &st.Title, &st.Description, &st.CreatedAt)

		// err := rows.Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone)
		if err != nil {
			return model.Todo{}, err
		}
		return model.Todo{ID: st.ID, Title: st.Title, Description: st.Description, CreatedAt: time.Now()}, nil
	}
	return model.Todo{}, nil

}

func (o database) Delete(id int32) error {

	db, _ := connection()

	_, err := db.Exec("DELETE from todo2 WHERE id = $1",
		id)
	//$ -> cara ngedefine variable di postgreSQL
	return err //kalau g
}

func (o database) List() ([]model.Todo, error) {

	db, _ := connection()

	rows, err := db.Query("SELECT id, title, description, createdat FROM todo2 LIMIT 10;")
	// rows, err := db.Query("SELECT id,name,address,phone FROM customers LIMIT 10;")
	// COALESCE supaya kalau birthdate kosong, return stringnya null (soalnya null isi kolomnya database error kalau dipanggil disini)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//defer -> yang harus dipanggil di akhir (biar ga lupa)

	res := make([]model.Todo, 0) //inisialisasi variable slice, dengan size 0

	// res := []model.Todo

	for rows.Next() {
		var st model.Todo
		err := rows.Scan(&st.ID, &st.Title, &st.Description, &st.CreatedAt)
		// err := rows.Scan(&cus.ID, &cus.Name, &cus.Address, &cus.Phone)
		if err != nil {
			return nil, err
		}
		res = append(res, st)
	}
	return res, nil
}
