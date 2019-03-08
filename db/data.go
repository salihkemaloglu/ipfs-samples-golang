package db

import (
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

// Item a Items by its id
type Item struct {
	Name        string `bson:"name" json:"name"`
	Value       string `bson:"value" json:"value"`
	Description string `bson:"description" json:"description"`
	ItemID      string `bson:"-" json:"-"`
	Count       int    `bson:"count" json:"count"`
}

// LsLink a Items by its id
type LsLink struct {
	Hash string
	Name string
	Size uint64
	Type int
}

//User string
type User struct {
	Username string `bson:"author_id"`
	Password string `bson:"content"`
	Title    string `bson:"title"`
}

// Insert a Items by its id
func (r Item) Insert() (string, error) {

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(r.Name))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return "nil", nil
	}

	return cid, nil

}

// GetAll a Items by its id
func (r Item) GetAll() {

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.List(r.Value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)

	}
	fmt.Println("len", len(cid))
	for index := 0; index < len(cid); index++ {
		fmt.Fprintf(os.Stderr, "error: %v", cid[index])
	}

}
