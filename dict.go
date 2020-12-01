package types

import "sort"

// Dict is map[string]Any
type Dict = map[string]Any

// DictKeys returns a sorted slice of string keys for the dict
func DictKeys(dict Dict) []string {
	keys, i := make([]string, len(dict)), 0
	for k := range dict {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
