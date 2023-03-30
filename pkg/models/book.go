package models

import(
	"github.com/jinzhu/gorm"
	"github.com/amit/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct(
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"name"`
	Author string string `json:"author"`
	Publication string `json:"publication"`
)


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}



func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.where("ID=?",Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.where("ID=?",Id).Delete(book)
	return book
}





