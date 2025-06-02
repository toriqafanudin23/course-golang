package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"setup-database/entity"

	"github.com/joho/godotenv"

	// "time"

	_ "github.com/lib/pq"
)

var psqlInfo string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func main() {
	// student := entity.Student{Id: 5, Name: "Dhiya Rohadatul Aisy", Email: "aisy@gmail.com", Address: "Wonosobo", BirthDate: time.Date(1999, 06, 16, 0, 0, 0, 0, time.Local), Gender: "F"}

	// addStudent(student)
	// updateStudent(student)
	// deleteStudent("4")
	students := getAllStudent()
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	// }
	// fmt.Println(getStudentById(2))
	fmt.Println(students)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func addStudent(student entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO mst_student (id, name, email, address, birth_date, gender) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data!")
	}
}

func updateStudent(student entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE mst_student SET name = $2, email = $3, address = $4, birth_date = $5, gender = $6 WHERE id = $1;"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Update Data!")
	}
}

func deleteStudent(id string) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM mst_student WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Delete Data!")
	}
}

func getAllStudent() []entity.Student {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func scanStudent(rows *sql.Rows) []entity.Student {
	students := []entity.Student{}
	var err error

	for rows.Next() {
		student := entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)
		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return students
}

func getStudentById(id int) entity.Student {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_student WHERE id = $1"

	student := entity.Student{}
	err = db.QueryRow(sqlStatement, id).Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)
	if err != nil {
		panic(err)
	}

	return student
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully connected!")
	}
	return db
}
