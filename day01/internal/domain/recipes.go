package domain

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Ingredients struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit" xml:"itemunit"`
}

func (r *Recipes) PrintXML() {
	toXML, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(toXML))
}

func (r *Recipes) PrintJSON() {
	toJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(toJSON))
}

func (r *Recipes) Compare(r2 *Recipes) ([]string, error) {
	var result []string

	// added or removed cakes
	for _, newCake := range r2.Cakes {
		if _, equalCakes := FindCake(r, newCake.Name); !equalCakes {
			str := fmt.Sprintf("ADDED cake \"%s\"", newCake.Name)
			result = append(result, str)
		}
	}
	for _, oldCake := range r.Cakes {
		if _, equalCakes := FindCake(r2, oldCake.Name); !equalCakes {
			str := fmt.Sprintf("REMOVED cake \"%s\"", oldCake.Name)
			result = append(result, str)
		}
	}

	// time changed
	for _, oldCake := range r.Cakes {
		if newCake, equalCakes := FindCake(r2, oldCake.Name); equalCakes {
			if newCake.Time != oldCake.Time {
				str := fmt.Sprintf("CHANGED cooking time for cake \"%s\"  -  \"%s\"  instead of \"%s\"",
					newCake.Name, newCake.Time, oldCake.Time)
				result = append(result, str)
			}
			for 
		}
	}

	return result, nil
}

func FindCake(recipes *Recipes, nameToFind string) (*Cake, bool) {
	for _, cake := range recipes.Cakes {
		if cake.Name == nameToFind {
			return &cake, true
		}
	}
	return nil, false
}
