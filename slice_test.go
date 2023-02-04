package smapper

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestSliceByStructs(t *testing.T) {
	type user struct {
		id     int64
		name   string
		salary float32
		age    *uint8
	}

	users := []user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: ref[uint8](35)},
		{id: 46327846528375, name: "user2", salary: 152.3, age: ref[uint8](37)},
		{id: 46327846528376, name: "user1", salary: 165.5, age: ref[uint8](25)},
		{id: 46327846528377, name: "user3", salary: 110.7, age: ref[uint8](23)},
	}

	t.Run("nil", func(t *testing.T) {
		s := SliceByStructs[user, int64]([]user{}, unsafe.Offsetof(user{}.id))
		require.ElementsMatch(t, nil, s)
	})

	t.Run("struct slice first field", func(t *testing.T) {
		s := SliceByStructs[user, int64](users, unsafe.Offsetof(user{}.id))
		require.ElementsMatch(t, []int64{46327846528374, 46327846528375, 46327846528376, 46327846528377}, s)
	})

	t.Run("struct slice second field", func(t *testing.T) {
		s := SliceByStructs[user, string](users, unsafe.Offsetof(user{}.name))
		require.ElementsMatch(t, []string{"user1", "user2", "user1", "user3"}, s)
	})

	t.Run("struct slice fourth field", func(t *testing.T) {
		s := SliceByStructs[user, *uint8](users, unsafe.Offsetof(user{}.age))
		require.ElementsMatch(t, []*uint8{ref[uint8](35), ref[uint8](37), ref[uint8](25), ref[uint8](23)}, s)
	})
}

func TestSliceByStructsP(t *testing.T) {
	type user struct {
		id     int64
		name   string
		salary float32
		age    *uint8
	}

	users := []*user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: ref[uint8](35)},
		nil,
		{id: 46327846528375, name: "user2", salary: 152.3, age: ref[uint8](37)},
		{id: 46327846528376, name: "user1", salary: 165.5, age: ref[uint8](25)},
		nil,
		{id: 46327846528377, name: "user3", salary: 110.7, age: ref[uint8](23)},
	}

	t.Run("nil", func(t *testing.T) {
		s := SliceByStructsP[user, int64]([]*user{}, unsafe.Offsetof(user{}.id))
		require.ElementsMatch(t, nil, s)
	})

	t.Run("struct slice first field", func(t *testing.T) {
		s := SliceByStructsP[user, int64](users, unsafe.Offsetof(user{}.id))
		require.ElementsMatch(t, []int64{46327846528374, 46327846528375, 46327846528376, 46327846528377}, s)
	})

	t.Run("struct slice second field", func(t *testing.T) {
		s := SliceByStructsP[user, string](users, unsafe.Offsetof(user{}.name))
		require.ElementsMatch(t, []string{"user1", "user2", "user1", "user3"}, s)
	})

	t.Run("struct slice fourth field", func(t *testing.T) {
		s := SliceByStructsP[user, *uint8](users, unsafe.Offsetof(user{}.age))
		require.ElementsMatch(t, []*uint8{ref[uint8](35), ref[uint8](37), ref[uint8](25), ref[uint8](23)}, s)
	})
}
