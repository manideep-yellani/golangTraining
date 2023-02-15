package store

import (
	"database/sql"

	"github.com/awesomeProject/internal/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db}
}

func Post(st models.Student) models.Student {
	db, _ := sql.Open("mysql", "root:1@tcp(3036)")

	db.Exec("insert into student(name,age,phone) values (?,?,?)", &st.Name, &st.Age, &st.Phone)
	//if err!=nil{
	//	return nil,"ERROR WHILE INSERTING"
	//}

	rows, _ := db.Query("select id from student where name=? and age=? and phone=? ", &st.Name, &st.Age, &st.Phone)
	var id int
	for rows.Next() {
		rows.Scan(&id)
	}
	st.Id = id
	return st

}
