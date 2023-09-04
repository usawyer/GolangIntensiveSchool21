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

func (r *Recipes) Compare(r2 *Recipes) []string {
	var result []string

	// added or removed cakes
	for _, newCake := range r2.Cakes {
		if _, equalCake := FindCake(r, newCake.Name); !equalCake {
			str := fmt.Sprintf("ADDED cake \"%s\"", newCake.Name)
			result = append(result, str)
		}
	}
	for _, oldCake := range r.Cakes {
		if _, equalCake := FindCake(r2, oldCake.Name); !equalCake {
			str := fmt.Sprintf("REMOVED cake \"%s\"", oldCake.Name)
			result = append(result, str)
		}
	}
	// time changed
	for _, oldCake := range r.Cakes {
		if newCake, equalCake := FindCake(r2, oldCake.Name); equalCake {
			if newCake.Time != oldCake.Time {
				str := fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"",
					newCake.Name, newCake.Time, oldCake.Time)
				result = append(result, str)
			}
			// added or removed ingredients
			for _, newIngredient := range newCake.Ingredients {
				if _, equalIngredient := FindIngredient(&oldCake, newIngredient.IngredientName); !equalIngredient {
					str := fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"",
						newIngredient.IngredientName, newCake.Name)
					result = append(result, str)
				}
			}
			for _, oldIngredient := range oldCake.Ingredients {
				if _, equalIngredient := FindIngredient(newCake, oldIngredient.IngredientName); !equalIngredient {
					str := fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"",
						oldIngredient.IngredientName, oldCake.Name)
					result = append(result, str)
				}
			}
			// count or unit changed
			for _, oldIngredient := range oldCake.Ingredients {
				if newIngredient, equalIngredient := FindIngredient(newCake, oldIngredient.IngredientName); equalIngredient {
					if newIngredient.IngredientCount != "" && oldIngredient.IngredientCount == "" {
						str := fmt.Sprintf("ADDED unit count \"%s\" for ingredient \"%s\" for cake \"%s\"",
							newIngredient.IngredientUnit, newIngredient.IngredientName, newCake.Name)
						result = append(result, str)
					} else if newIngredient.IngredientCount == "" && oldIngredient.IngredientCount != "" {
						str := fmt.Sprintf("REMOVED unit count \"%s\" for ingredient \"%s\" for cake \"%s\"",
							oldIngredient.IngredientCount, newIngredient.IngredientName, newCake.Name)
						result = append(result, str)
					} else if newIngredient.IngredientCount != oldIngredient.IngredientCount {
						str := fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"",
							newIngredient.IngredientName, newCake.Name, newIngredient.IngredientCount, oldIngredient.IngredientCount)
						result = append(result, str)
					}
					if newIngredient.IngredientUnit != "" && oldIngredient.IngredientUnit == "" {
						str := fmt.Sprintf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"",
							newIngredient.IngredientUnit, newIngredient.IngredientName, newCake.Name)
						result = append(result, str)
					} else if newIngredient.IngredientUnit == "" && oldIngredient.IngredientUnit != "" {
						str := fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"",
							oldIngredient.IngredientUnit, newIngredient.IngredientName, newCake.Name)
						result = append(result, str)
					} else if newIngredient.IngredientUnit != oldIngredient.IngredientUnit {
						str := fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"",
							newIngredient.IngredientName, newCake.Name, newIngredient.IngredientUnit, oldIngredient.IngredientUnit)
						result = append(result, str)
					}
				}
			}
		}
	}
	return result
}

func FindCake(recipes *Recipes, nameToFind string) (*Cake, bool) {
	for _, cake := range recipes.Cakes {
		if cake.Name == nameToFind {
			return &cake, true
		}
	}
	return nil, false
}

func FindIngredient(cake *Cake, nameToFind string) (*Ingredients, bool) {
	for _, ingredient := range cake.Ingredients {
		if ingredient.IngredientName == nameToFind {
			return &ingredient, true
		}
	}
	return nil, false
}
