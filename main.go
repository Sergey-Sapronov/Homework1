package main

import (
	"education/list/storages/list"
	"education/list/storages/model"
	"education/list/storages/slice"
)

func main() {
	var s model.Storage
	s = slice.NewSlice()
	s.Add(1)
	s.Add(5)
	s.Delete(2)
	s.Add(0)
	s.Print()
	s.Delete(3)
	s.SortDecrease()
	s.Print()
	l := list.NewList()
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Print()
	l.Add(1)
	l.SortDecrease()
	l.Print()
}
