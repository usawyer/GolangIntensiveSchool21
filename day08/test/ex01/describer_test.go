package test

import (
	"day08/internal/describer"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestUnknownPlant(t *testing.T) {
	plant := describer.UnknownPlant{
		FlowerType: "Lavender",
		LeafType:   "oblong",
		Color:      230,
	}

	expected := "FlowerType:Lavender,\nLeafType:oblong,\nColor(color_scheme=rgb):230\n"
	output := captureOutput(describer.DescribePlant, plant)
	assert.Equal(t, expected, output)
}

func TestPlant(t *testing.T) {
	plant := describer.AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	}

	expected := "FlowerColor:10,\nLeafType:lanceolate,\nHeight(unit=inches):15\n"
	output := captureOutput(describer.DescribePlant, plant)
	assert.Equal(t, expected, output)
}

func TestErrorInput(t *testing.T) {
	expected := "Invalid type to describe plant\n"
	output := captureOutput(describer.DescribePlant, 12)
	assert.Equal(t, expected, output)
}

func captureOutput(f func(interface{}), plant interface{}) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f(plant)
	w.Close()
	os.Stdout = old

	capturedOutput, _ := io.ReadAll(r)
	return string(capturedOutput)
}
