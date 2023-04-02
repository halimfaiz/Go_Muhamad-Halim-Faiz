package database

import (
	"Praktikum/config"
	"Praktikum/models"
)

var (
	books []models.Book
	book  models.Book
)

func GetBooks() (interface{}, error) {
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	if e := config.DB.Find(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func Createbook(book *models.Book) (*models.Book, error) {
	if e := config.DB.Save(book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBookById(id int) (interface{}, error) {
	if e := config.DB.Delete(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func UpdateBookById(id int, updatedBook *models.Book) (*models.Book, error) {
	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}

	book.Judul = updatedBook.Judul
	book.Penulis = updatedBook.Penulis
	book.Penerbit = updatedBook.Penerbit

	if e := config.DB.Save(&book).Error; e != nil {
		return nil, e
	}
	return updatedBook, nil
}
