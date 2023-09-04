package domain

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"stovetime"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name" xml:"itemname"`
			IngredientCount string `json:"ingredient_count" xml:"itemcount"`
			IngredientUnit  string `json:"ingredient_unit" xml:"itemunit"`
		} `json:"ingredients" xml:"ingredients>item"`
	} `json:"cake" xml:"cake"`
}

func (r *Recipes) PrintXML() {
	toXml, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(toXml))
}

func (r *Recipes) PrintJSON() {
	toJson, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(toJson))
}

//
//func (r *Recipes) Compare(r2 *Recipes) ([]string, error) {
//
//}
