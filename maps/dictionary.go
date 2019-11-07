package maps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	definition, found := d[word]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}
