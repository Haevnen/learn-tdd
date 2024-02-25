package main

// var ErrNotFound = errors.New("could not find the word you were looking for")
const ErrNotFound = DictionaryErr("could not find the word you were looking for")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// a map value is a pointer to a runtime.hmap structure
// Not a reference type
// We can perform read in nil map. But write will make a runtime panic
// So, You should never initialize an empty map variable
// E.g., var m map[string]string -> m := map[string]string
type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if value, ok := d[key]; ok {
		return value, nil
	}

	return "", ErrNotFound
}
