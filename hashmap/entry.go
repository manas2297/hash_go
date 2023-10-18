package hashmap

type Entry struct {
	Key   interface{}
	Value interface{}
	Next  *Entry
}

func NewEntry(key interface{}, value interface{}) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
	}
}
