package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrWordNotFound error = errors.New("word not in dictionary")
var ErrWordExists error = errors.New("word already exists")

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}
