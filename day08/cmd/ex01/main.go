package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) {
	v := reflect.ValueOf(plant)
	t := v.Type()

	if t.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			printField(v.Field(i), t.Field(i))
			if i != v.NumField()-1 {
				fmt.Print(",")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Invalid type to describe plant")
	}
}

func printField(f reflect.Value, t reflect.StructField) {
	tag := t.Tag.Get("color_scheme")
	if tag == "" {
		tag = t.Tag.Get("unit")
	}

	if tag == "" {
		fmt.Printf("%s:%v", t.Name, f)
	} else {
		fmt.Printf("%s(%s=%s):%v", t.Name, strings.Split(string(t.Tag), ":")[0], tag, f)
	}
}

func main() {
	plant := UnknownPlant{
		FlowerType: "Lavender",
		LeafType:   "oblong",
		Color:      230,
	}

	plant2 := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	}

	describePlant(plant)
	fmt.Println()
	describePlant(plant2)
	fmt.Println()
	describePlant(12)
}
