package types

import "strconv"

// ToStringInt casts int to string
func ToStringInt(int int) string { return ToStringI64(int64(int)) }

// ToStringI64 casts int64 to string
func ToStringI64(int64 int64) string { return strconv.FormatInt(int64, 10) }
