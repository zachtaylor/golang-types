package types

// Hasher is hashing interface
type Hasher interface {
	Hash(string) string
}
