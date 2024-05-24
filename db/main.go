package db

import "github.com/kisinga/go-htmx-tictactoe/model"

type inMemoryDB struct {
	Games   *map[string]*model.Board
	Clients *map[string][]chan *model.Cell
}

func (db inMemoryDB) Create() {

}

func (db inMemoryDB) Read() {

}

func (db inMemoryDB) Update() {

}

func (db inMemoryDB) Delete() {

}

type DB interface {
	Create()
	Read()
	Update()
	Delete()
}

func NewDB() DB {
	return &inMemoryDB{}
}
