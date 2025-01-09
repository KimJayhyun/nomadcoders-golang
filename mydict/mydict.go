package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("could not find the word you were looking for")
	errWordExist = errors.New("word is already exists")
	errCantUpdate = errors.New("can not update non-exists word")
	errCantDelete = errors.New("can not delete non-exsists word")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, word)
	}
	return nil
}