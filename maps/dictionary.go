// Sample code to demonstrate the use of maps
package maps

// Dictionary maps definitions to words
type Dictionary map[string]string

const (
	// ErrNotFound means the requested word was not found in the dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists indicates you are trying to add a word that already exists
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	// ErrWordDoesNotExist occure when trying to update a word that is not in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr are `errors` that can happen when interacting with a dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search for the definition of a `word` in the dictionary
//
// Returns the definition of the word, or an `ErrNotFound` error
// if the `word` cannot be found in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add a new definition to the dictionary
func (d Dictionary) Add(word, definition string) error {

	// If the word cannot be found then add it. Otherwise
	// return ErrWordExists error.
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update the definition of an existing word in the dictionary
func (d Dictionary) Update(word, definition string) error {
	// If the word cannot be found then return an error. Otherwise
	// update the definition.
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete removes a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
