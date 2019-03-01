package db

import (
	"fmt"
)

// Item a Items by its id
type Item struct {
	Name        string `bson:"name" json:"name"`
	Value       string `bson:"value" json:"value"`
	Description string `bson:"description" json:"description"`
	ItemID      string `bson:"-" json:"-"`
	Count       int    `bson:"count" json:"count"`
}

// FindByID a Items by its id
func (r Item) FindByID() ([]byte, error) {

	fmt.Println(r)

	return nil, nil
}
