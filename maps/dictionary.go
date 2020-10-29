package main

import "errors"

// NOTE: Map is passed by reference already.
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	word, ok := d[word]
	if !ok {
		return "", errors.New("not found")
	}
	return word, nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return errors.New("exists")
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}
	d[word] = newDefinition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}
	delete(d, word)
	return nil
}
