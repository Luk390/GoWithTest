package main

import (
	"errors"
	"go/types"
)

type Dictionary map[string]string

const (
	ErrWordNotFound = DictionaryError("word not in dictionary")
	ErrWordExists   = DictionaryError("word already exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}


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
