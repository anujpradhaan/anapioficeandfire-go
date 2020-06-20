package characters

import (
	"encoding/json"
	"fmt"
	"gameofthrones/resttemplate"
	"gameofthrones/urls"
	"gameofthrones/util"
)

type Character struct {
	URL         string   `json:"url"`         // The hypermedia URL of this resource
	Name        string   `json:"name"`        // The name of this character
	Gender      string   `json:"gender"`      //The gender of this character.
	Culture     string   `json:"culture"`     //The culture that this character belongs to.
	Born        string   `json:"born"`        //Textual representation of when and where this character was born.
	Died        string   `json:"died"`        //Textual representation of when and where this character died.
	Titles      []string `json:"titles"`      //TThe titles that this character holds.
	Aliases     []string `json:"aliases"`     //The aliases that this character goes by.
	Father      string   `json:"father"`      //The character resource URL of this character's father.
	Mother      string   `json:"mother"`      //The character resource URL of this character's mother.
	Spouse      string   `json:"spouse"`      //An array of Character resource URLs that has had a POV-chapter in this book.
	Allegiances []string `json:"allegiances"` //An array of House resource URLs that this character is loyal to.
	Books       []string `json:"books"`       //An array of Book resource URLs that this character has been in.
	PovBooks    []string `json:"povBooks"`    //An array of Book resource URLs that this character has had a POV-chapter in.
	TvSeries    []string `json:"tvSeries"`    //An array of names of the seasons of Game of Thrones that this character has been in.
	PlayedBy    []string `json:"playedBy"`    //An array of actor names that has played this character in the TV show Game Of Thrones.

}

type CharacterFilter struct {
	Name    string `tag:"name"`    //Characters with the given name are included in the response.
	Gender  string `tag:"gender"`  //Characters with the given gender are included in the response.
	Culture string `tag:"culture"` //Characters with the given culture are included in the response
	Born    string `tag:"born"`    //Characters that were born this given year are included in the response.
	Died    string `tag:"died"`    //Characters that died this given year are included in the response.
	IsAlive bool   `tag:"isAlive"` //Textual representation of when and where this character died.
}

func FindAll() []Character {
	var characters []Character
	resttemplate.MakeGetCall(urls.Characters, &characters, map[string]string{})
	prettyJSON, _ := json.MarshalIndent(characters, "", "    ")
	fmt.Printf("%s = %s\n", urls.Books, string(prettyJSON))
	return characters
}

func FindById(id int) Character {
	var character Character
	url := fmt.Sprintf("%s/%d", urls.Characters, id)
	resttemplate.MakeGetCall(url, &character, map[string]string{})
	return character
}

func FindByName(nameOfCharacter string) []Character {
	var characters []Character
	queryParamMap := util.GetFilterMap(CharacterFilter{Name: nameOfCharacter})
	resttemplate.MakeGetCall(urls.Characters, &characters, queryParamMap)
	return characters
}

func FindByFilter(characterFilter CharacterFilter) []Character {
	var characters []Character
	queryParamMap := util.GetFilterMap(characterFilter)
	resttemplate.MakeGetCall(urls.Characters, &characters, queryParamMap)
	return characters
}
