package maps

type Dictionary map[string]string

type Search interface {
	Dictionary() string
}

func (d Dictionary) Search(word string) string {
	return d[word]
}
