package books

import (
	"encoding/json"
	"fmt"
	"gameofthrones/resttemplate"
	"gameofthrones/urls"
	"gameofthrones/util"
)

type Book struct {
	URL           string   `json:"url"`           //The hypermedia URL of this resource
	Name          string   `json:"name"`          //The name of this book
	Isbn          string   `json:"isbn"`          //The International Standard Book Number (ISBN-13) that uniquely identifies this book.
	Authors       []string `json:"authors"`       //An array of names of the authors that wrote this book.
	NumberOfPages int      `json:"numberOfPages"` //The number of pages in this book.
	Publisher     string   `json:"publisher"`     //The company that published this book.
	Country       string   `json:"country"`       //The country that this book was published in
	MediaType     string   `json:"mediaType"`     //The type of media this book was released in.
	Released      string   `json:"released"`      //The date (ISO 8601) when this book was released.
	Characters    []string `json:"characters"`    //An array of Character resource URLs that has been in this book.
	PovCharacters []string `json:"povCharacters"` //An array of Character resource URLs that has had a POV-chapter in this book.
}

type BookFilter struct {
	Name            string `tag:"name"`
	FromReleaseDate string `tag:"fromReleaseDate"`
	ToReleaseDate   string `tag:"toReleaseDate"`
}

func FindAll() []Book {
	var books []Book
	resttemplate.MakeGetCall(urls.Books, &books, map[string]string{})
	prettyJSON, _ := json.MarshalIndent(books, "", "    ")
	fmt.Printf("%s = %s\n", urls.Books, string(prettyJSON))
	return books
}

func FindById(id int) Book {
	var book Book
	url := fmt.Sprintf("%s/%d", urls.Books, id)
	resttemplate.MakeGetCall(url, &book, map[string]string{})
	return book
}

func FindByName(nameofBook string) []Book {
	var books []Book
	queryParamMap := util.GetFilterMap(BookFilter{Name: nameofBook})
	resttemplate.MakeGetCall(urls.Books, &books, queryParamMap)
	return books
}

func FindByFilter(bookFilter BookFilter) []Book {
	var books []Book
	queryParamMap := util.GetFilterMap(bookFilter)
	resttemplate.MakeGetCall(urls.Books, &books, queryParamMap)
	return books
}
