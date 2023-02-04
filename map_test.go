package smapper

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMapBySlice(t *testing.T) {
	type user struct {
		id     int64
		name   string
		salary float32
		age    int8
	}

	users := []user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: 35},
		{id: 46327846528375, name: "user2", salary: 152.3, age: 37},
		{id: 46327846528376, name: "user1", salary: 165.5, age: 25},
		{id: 46327846528377, name: "user3", salary: 110.7, age: 23},
	}

	t.Run("nil struct slice", func(t *testing.T) {
		var emptyUsers []user
		m := MapBySlice[user, string](emptyUsers, unsafe.Offsetof(user{}.name))
		require.Equal(t, (map[string]user)(nil), m)
	})

	t.Run("nil slice", func(t *testing.T) {
		var ids []int64
		m := MapBySlice[int64, int64](ids, 0)
		require.Equal(t, (map[int64]int64)(nil), m)
	})

	t.Run("struct slice first field", func(t *testing.T) {
		m := MapBySlice[user, int64](users, unsafe.Offsetof(user{}.id))
		require.Equal(t, map[int64]user{
			46327846528374: {id: 46327846528374, name: "user1", salary: 102.1, age: 35},
			46327846528375: {id: 46327846528375, name: "user2", salary: 152.3, age: 37},
			46327846528376: {id: 46327846528376, name: "user1", salary: 165.5, age: 25},
			46327846528377: {id: 46327846528377, name: "user3", salary: 110.7, age: 23},
		}, m)
	})

	t.Run("struct slice second field", func(t *testing.T) {
		m := MapBySlice[user, string](users, unsafe.Offsetof(user{}.name))
		require.Equal(t, map[string]user{
			"user1": {id: 46327846528376, name: "user1", salary: 165.5, age: 25},
			"user2": {id: 46327846528375, name: "user2", salary: 152.3, age: 37},
			"user3": {id: 46327846528377, name: "user3", salary: 110.7, age: 23},
		}, m)
	})
}

func TestMapBySliceP(t *testing.T) {
	type user struct {
		id     int64
		name   string
		salary float32
		age    int8
	}

	users := []*user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: 35},
		{id: 46327846528375, name: "user2", salary: 152.3, age: 37},
		{id: 46327846528376, name: "user1", salary: 165.5, age: 25},
		{id: 46327846528377, name: "user3", salary: 110.7, age: 23},
	}

	t.Run("nil struct slice", func(t *testing.T) {
		var emptyUsers []*user
		m := MapBySliceP[user, string](emptyUsers, unsafe.Offsetof(user{}.name))
		require.Equal(t, (map[string]*user)(nil), m)
	})

	t.Run("nil slice", func(t *testing.T) {
		var ids []*int64
		m := MapBySliceP[int64, int64](ids, 0)
		require.Equal(t, (map[int64]*int64)(nil), m)
	})

	t.Run("struct slice first field", func(t *testing.T) {
		m := MapBySliceP[user, int64](users, unsafe.Offsetof(user{}.id))
		require.Equal(t, map[int64]*user{
			46327846528374: {id: 46327846528374, name: "user1", salary: 102.1, age: 35},
			46327846528375: {id: 46327846528375, name: "user2", salary: 152.3, age: 37},
			46327846528376: {id: 46327846528376, name: "user1", salary: 165.5, age: 25},
			46327846528377: {id: 46327846528377, name: "user3", salary: 110.7, age: 23},
		}, m)
	})

	t.Run("struct slice fourth field", func(t *testing.T) {
		m := MapBySliceP[user, int8](users, unsafe.Offsetof(user{}.age))
		require.Equal(t, map[int8]*user{
			35: {id: 46327846528374, name: "user1", salary: 102.1, age: 35},
			37: {id: 46327846528375, name: "user2", salary: 152.3, age: 37},
			25: {id: 46327846528376, name: "user1", salary: 165.5, age: 25},
			23: {id: 46327846528377, name: "user3", salary: 110.7, age: 23},
		}, m)
	})
}

func TestSetBySlice(t *testing.T) {
	type user struct {
		id     int64
		name   string
		salary float32
		age    int8
	}

	users := []user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: 35},
		{id: 46327846528375, name: "user2", salary: 152.3, age: 37},
		{id: 46327846528376, name: "user1", salary: 165.5, age: 25},
		{id: 46327846528377, name: "user3", salary: 110.7, age: 23},
	}

	t.Run("nil struct slice", func(t *testing.T) {
		var emptyUsers []user
		m := SetBySlice[user, string](emptyUsers, unsafe.Offsetof(user{}.name))
		require.Equal(t, (map[string]struct{})(nil), m)
	})

	t.Run("nil slice", func(t *testing.T) {
		var ids []int64
		m := SetBySlice[int64, int64](ids, 0)
		require.Equal(t, (map[int64]struct{})(nil), m)
	})

	t.Run("slice", func(t *testing.T) {
		ids := []int64{1, 4, 5, 6, 7, 7, 6}
		m := SetBySlice[int64, int64](ids, 0)
		require.Equal(t, map[int64]struct{}{
			1: {},
			4: {},
			5: {},
			6: {},
			7: {},
		}, m)
	})

	t.Run("slice of poiners", func(t *testing.T) {
		a, b, c := ref(1), ref(2), ref(2)
		ids := []*int{a, b, c, nil, nil}
		m := SetBySlice[*int, *int](ids, 0)
		require.Equal(t, map[*int]struct{}{
			a:   {},
			b:   {},
			c:   {},
			nil: {},
		}, m)
	})

	t.Run("struct slice first field", func(t *testing.T) {
		m := SetBySlice[user, int64](users, unsafe.Offsetof(user{}.id))
		require.Equal(t, map[int64]struct{}{
			46327846528374: {},
			46327846528375: {},
			46327846528376: {},
			46327846528377: {},
		}, m)
	})

	t.Run("struct slice second field", func(t *testing.T) {
		m := SetBySlice[user, float32](users, unsafe.Offsetof(user{}.salary))
		require.Equal(t, map[float32]struct{}{
			102.1: {},
			152.3: {},
			165.5: {},
			110.7: {},
		}, m)
	})
}

func TestSetBySliceP(t *testing.T) {
	type user struct {
		id     int64
		name   string
		salary float32
		age    int8
	}

	users := []*user{
		{id: 46327846528374, name: "user1", salary: 102.1, age: 35},
		{id: 46327846528375, name: "user2", salary: 152.3, age: 37},
		nil,
		{id: 46327846528376, name: "user1", salary: 165.5, age: 25},
		nil,
		{id: 46327846528377, name: "user3", salary: 110.7, age: 23},
	}

	t.Run("nil struct slice", func(t *testing.T) {
		var emptyUsers []*user
		m := SetBySliceP[user, string](emptyUsers, unsafe.Offsetof(user{}.name))
		require.Equal(t, (map[string]struct{})(nil), m)
	})

	t.Run("nil slice", func(t *testing.T) {
		var ids []*int64
		m := SetBySliceP[int64, *int64](ids, 0)
		require.Equal(t, (map[*int64]struct{})(nil), m)
	})

	t.Run("slice of poiners", func(t *testing.T) {
		a, b, c := ref(1), ref(2), ref(2)
		ids := []*int{a, b, c, nil, nil}
		m := SetBySliceP[int, int](ids, 0)
		fmt.Print(m)
		require.Equal(t, map[int]struct{}{
			1: {},
			2: {},
		}, m)
	})

	t.Run("struct slice first field", func(t *testing.T) {
		m := SetBySliceP[user, int64](users, unsafe.Offsetof(user{}.id))
		require.Equal(t, map[int64]struct{}{
			46327846528374: {},
			46327846528375: {},
			46327846528376: {},
			46327846528377: {},
		}, m)
	})

	t.Run("struct slice second field", func(t *testing.T) {
		m := SetBySliceP[user, float32](users, unsafe.Offsetof(user{}.salary))
		require.Equal(t, map[float32]struct{}{
			102.1: {},
			152.3: {},
			165.5: {},
			110.7: {},
		}, m)
	})
}
