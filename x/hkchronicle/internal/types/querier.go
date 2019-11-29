package types

import "strings"

// Query Result Payload for a resolve query
type QueryResEventResolve struct {
	Value string `json:"value"`
	// Time  int64	 `json:"time"`
}

// implement fmt.Stringer
func (r QueryResEventResolve) String() string {
	return r.Value
}

// Query Result Payload for a names query
type QueryResAllEvents []string

// implement fmt.Stringer
func (n QueryResAllEvents) String() string {
	return strings.Join(n[:], "\n")
}
