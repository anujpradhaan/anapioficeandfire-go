package houses

import (
	"encoding/json"
	"fmt"
	"gameofthrones/resttemplate"
	"gameofthrones/urls"
	"gameofthrones/util"
)

type House struct {
	URL              string   `json:"url"`              //The hypermedia URL of this resource
	Name             string   `json:"name"`             //The name of this house
	Region           string   `json:"region"`           //The region that this house resides in.
	CoatOfArms       string   `json:"coatOfArms"`       //Text describing the coat of arms of this house.
	Words            string   `json:"words"`            //The words of this house.
	Titles           []string `json:"titles"`           //The titles that this house holds.
	Seats            []string `json:"seats"`            //The seats that this house holds.
	CurrentLord      string   `json:"currentLord"`      //The Character resource URL of this house's current lord.
	Heir             string   `json:"heir"`             //The Character resource URL of this house's heir.
	Overlord         string   `json:"overlord"`         //The House resource URL that this house answers to.
	Founded          string   `json:"founded"`          //The year that this house was founded.
	Founder          string   `json:"founder"`          //The Character resource URL that founded this house.
	DiedOut          string   `json:"diedOut"`          //The year that this house died out.
	AncestralWeapons []string `json:"ancestralWeapons"` //An array of names of the noteworthy weapons that this house owns.
	CadetBranches    []string `json:"cadetBranches"`    //An array of House resource URLs that was founded from this house.
	SwornMembers     []string `json:"swornMembers"`     //An array of Character resource URLs that are sworn to this house.
}

type HouseFilter struct {
	Name                string `tag:"url"`                 //Houses with the given name are included in the response
	Region              string `tag:"region"`              //Houses that belong in the given region are included in the response.
	Words               string `tag:"words"`               //Houses that has the given words are included in the response.
	HasWords            bool   `tag:"hasWords"`            //Houses that have words (or not) are included in the response.
	HasTitles           bool   `tag:"hasTitles"`           //Houses that have titles (or not) are included in the response.
	HasSeats            bool   `tag:"hasSeats"`            //Houses that have seats (or not) are included in the response.
	HasDiedOut          bool   `tag:"hasDiedOut"`          //Houses that are extinct are included in the response.
	HasAncestralWeapons bool   `tag:"hasAncestralWeapons"` //Houses that have ancestral weapons (or not) are included in the response.
}

func FindAll() []House {
	var houses []House
	resttemplate.MakeGetCall(urls.Houses, &houses, map[string]string{})
	prettyJSON, _ := json.MarshalIndent(houses, "", "    ")
	fmt.Printf("%s = %s\n", urls.Books, string(prettyJSON))
	return houses
}

func FindById(id int) House {
	var house House
	url := fmt.Sprintf("%s/%d", urls.Houses, id)
	resttemplate.MakeGetCall(url, &house, map[string]string{})
	return house
}

func FindByName(nameOfHouse string) []House {
	var houses []House
	queryParamMap := util.GetFilterMap(HouseFilter{Name: nameOfHouse})
	resttemplate.MakeGetCall(urls.Houses, &houses, queryParamMap)
	return houses
}

func FindByFilter(houseFilter HouseFilter) []House {
	var houses []House
	queryParamMap := util.GetFilterMap(houseFilter)
	resttemplate.MakeGetCall(urls.Houses, &houses, queryParamMap)
	return houses
}
