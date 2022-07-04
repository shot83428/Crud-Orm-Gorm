package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var server = "localhost"
var user = "sa"
var password = "sa"
var port = 1433
var database = "Booker"

var db *gorm.DB

func main() {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)
	var err error
	db, err = gorm.Open("mssql", connectionString)

	if err != nil {
		log.Fatal("Failed to create connection pool. Error: " + err.Error())
	}
	Read()
	Addtional("Kevin", "0056")
	//update("a", "gg")
	//Delete(4)

}

type Booked struct {
	Id    int64  `gorm:"column:Id"`
	Name  string `gorm:"column:Name"`
	Train string `gorm:"column:Train"`
}

func Read() {
	Books := make([]Booked, 0)

	//db.Find(&Books)
	var val int
	db.Table("TestBook.Booked").Count(&val)
	fmt.Println(val)

	db.Table("TestBook.Booked").Find(&Books)

	for _, Book := range Books {
		db.Model(Book)
		fmt.Println("Name:", Book.Name, "Train:", Book.Train)
	}

}

func Addtional(name string, train string) {
	// err := db.Table("TestBook.Booked").Exec("SET IDENTITY_INSERT TestBook.Booked ON;").Error
	// if err != nil {
	// 	fmt.Println(err)
	// }
	newBook := Booked{Name: name, Train: train}
	err := db.Table("TestBook.Booked").Select("Name", "Train").Create(&newBook).Error
	if err != nil {
		fmt.Println(err)
	}
}

func Delete(val int) {
	err := db.Table("TestBook.Booked").Where("Id = ?", val).Delete(&Booked{}).Error
	if err != nil {
		fmt.Println(err)
	}
}
func update(name string, train string) {
	err := db.Table("TestBook.Booked").Where("Name = ?", name).Update("Train", train).Error
	if err != nil {
		fmt.Println(err)
	}
}
