package object

import (
	"fmt"
	"hash/fnv"
	"strings"
)

// EX: map[Object]Object, we create map["key"] = "value". The problem arises when accessing the value
// map ["key"] will not get the value because "key" represents the seperate *object.String, not old one
// So the possible solutions would be check the index value through each key in GOs internal map object, or use
// hash as key during storing and retrieveing, we are using the second option

// HashKey holds the hash for key
type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKeyFunc() HashKey {

	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKeyFunc() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HasheyFunc() HashKey {

	// For string we are generating and using the hash of the string
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair is object for holding the key and value
type HashPair struct {
	Key   Object
	Value Object
}

// Hash maps hashKey to the map represented in the guest langauge
type Hash struct {
	Pairs map[HashKey]HashPair // We are using this so that we can get key and value easily
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	return fmt.Sprintf(`{%s}`, strings.Join(pairs, ", "))
}
