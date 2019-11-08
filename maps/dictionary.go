package maps

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word exists in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	definition, found := d[word]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
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

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
