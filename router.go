package main

import (
	"encoding/json"
	"fmt"
	"gameofthrones/characters"
)

func main() {
	// root := models.Root{}
	// resttemplate.MakeGetCall(urls.Root, &root)
	// //fmt.Printf("%+v\n", root)
	// prettyJSON, _ := json.MarshalIndent(root, "", "    ")
	// fmt.Printf("%s = %s\n", urls.Root, string(prettyJSON))

	// books := books.FindAll()
	// prettyJSON, _ := json.MarshalIndent(books, "", "    ")
	// fmt.Printf("%s\n", string(prettyJSON))

	//book := books.FindByName("A Game of Thrones")
	// book := books.FindById(1)

	// //book := books.FindByReleaseDates("1996-08-0", "")
	// prettyJSON, _ := json.MarshalIndent(book, "", "    ")
	// fmt.Printf("%s\n", string(prettyJSON))

	character := characters.FindById(1)

	//book := books.FindByReleaseDates("1996-08-0", "")
	prettyJSON, _ := json.MarshalIndent(character, "", "    ")
	fmt.Printf("%s\n", string(prettyJSON))

	// character := models.Character{}
	// url = fmt.Sprintf("%s/%d", urls.Characters, 1)
	// resttemplate.MakeGetCall(url, &character)
	// prettyJSON, _ = json.MarshalIndent(character, "", "    ")
	// fmt.Printf("%s = %s\n", url, string(prettyJSON))

	// house := models.House{}
	// url = fmt.Sprintf("%s/%d", urls.Houses, 1)
	// resttemplate.MakeGetCall(url, &house)
	// prettyJSON, _ = json.MarshalIndent(house, "", "    ")
	// fmt.Printf("%s = %s\n", url, string(prettyJSON))
}
