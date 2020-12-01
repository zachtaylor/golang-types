package types

// Map is map[Any]Any
type Map = map[Any]Any

// MapKeys returns a slice of keys for the map
func MapKeys(m Map) Slice {
	keys := make(Slice, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
