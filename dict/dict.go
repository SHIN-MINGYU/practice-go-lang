package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFount = errors.New("Not Found")
var errCantUpdate = errors.New("Cant Update non-existing word")
var errWordExist = errors.New("Word is Exist")

// type also can have method 

// func can return more than 2 variables
// Search for a word
func (d Dictionary) Search(word string) (string,error){
	value, exists := d[word]
	// the type of map return the 2 type of value, that is the value of map, and boolean what confirm the value is exist

	if exists{
		return value, nil
	}
	return "", errNotFount
}

// Add a word in dictonary
func (d Dictionary) Add(word, def string) error{
	_, err := d.Search(word)
	if err == errNotFount{
		d[word] = def
		return nil
	}
	return errWordExist
} 

//Update a word in dictionary
func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)

	switch err{
	case nil:
		d[word] = definition
	case errNotFount:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string){
	delete(d, word)
}