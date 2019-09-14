// Sample code to demonstrate the use of maps
package maps

type Dictionary map[string]string

func (d Dictionary) Search(dictionary map[string]string, word string) string {
	return d[word]
}
