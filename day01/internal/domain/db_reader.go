package domain

type DBReader interface {
	Read() (*Recipes, error)
}
