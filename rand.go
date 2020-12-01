package types

import "math/rand"

// Rand = rand.Rand
type Rand = rand.Rand

// Seeder = rand.Source
type Seeder = rand.Source

// NewRand = rand.New
func NewRand(seed Seeder) *Rand { return rand.New(seed) }

// NewSeeder = rand.NewSource
func NewSeeder(seed int64) Seeder { return rand.NewSource(seed) }
