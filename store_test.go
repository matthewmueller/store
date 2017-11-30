package store_test

import (
	"testing"

	"github.com/matthewmueller/store"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	db, err := store.New("store_test")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.Put("a", 1)
	var v int
	if err := db.Get("a", &v); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, v)

	if err := db.Delete("a"); err != nil {
		t.Fatal(err)
	}

	if err := db.Get("a", &v); err != store.ErrNotFound {
		t.Fatal(err)
	}
}

func TestPersistence(t *testing.T) {
	db, err := store.New("store_test")
	if err != nil {
		t.Fatal(err)
	}

	db.Put("b", 1)

	var v int
	if err := db.Get("b", &v); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, v)

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	db, err = store.New("store_test")
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Get("b", &v); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, v)

	if err := db.Delete("b"); err != nil {
		t.Fatal(err)
	}
}
