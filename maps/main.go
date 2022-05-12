package main


type Dictionary map[string]string

const (
	ErrWordNotFound = DictionaryError("word not in dictionary")
	ErrWordExists   = DictionaryError("word already exists")
	ErrDoesNotExist = DictionaryError("could not update definition as the word does not exist")
	ErrCouldNotDelete = DictionaryError("could not delete as word does not exist")
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrCouldNotDelete
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}

