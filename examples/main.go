package main

import (
	"fmt"
	"unsafe"

	"github.com/nastvood/smapper"
)

type user struct {
	id     int64
	name   string
	salary float32
	age    int8
}

func (u *user) String() string {
	if u == nil {
		return ""
	}

	return fmt.Sprintf("%+v", *u)
}

var (
	users = []user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: 35},
		{id: 46327846528375, name: "user2", salary: 152.3, age: 37},
		{id: 46327846528376, name: "user1", salary: 165.5, age: 25},
		{id: 46327846528377, name: "user3", salary: 110.7, age: 23},
	}

	pusers = []*user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: 35},
		nil,
		{id: 46327846528375, name: "user2", salary: 152.3, age: 37},
		{id: 46327846528376, name: "user1", salary: 165.5, age: 25},
		nil,
		{id: 46327846528377, name: "user3", salary: 110.7, age: 23},
	}
)

func mapBySlice() {
	fmt.Printf("%+v\n\n", smapper.MapBySlice[user, string](users, unsafe.Offsetof(user{}.name)))
}

func mapBySliceP() {
	fmt.Printf("%s\n\n", smapper.MapBySliceP[user, string](pusers, unsafe.Offsetof(user{}.name)))
}

func setBySlice() {
	fmt.Printf("%s\n\n", smapper.SetBySlice[user, string](users, unsafe.Offsetof(user{}.name)))
}

func setBySliceP() {
	fmt.Printf("%+v\n\n", smapper.SetBySliceP[user, int64](pusers, unsafe.Offsetof(user{}.id)))
}

func sliceByStructs() {
	fmt.Printf("%#v\n\n", smapper.SliceByStructs[user, int64](users, unsafe.Offsetof(user{}.id)))
}

func sliceByStructsP() {
	fmt.Printf("%#v\n\n", smapper.SliceByStructsP[user, string](pusers, unsafe.Offsetof(user{}.name)))
}

func main() {
	mapBySlice()

	mapBySliceP()

	setBySlice()

	setBySliceP()

	sliceByStructs()

	sliceByStructsP()
}
