// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"encoding/binary"
	"github.com/attic-labs/noms/go/hash"
)

// Int is a Noms Value wrapper around the primitive int32 type.
type Uint uint64

// Value interface
func (v Uint) Value() Value {
	return v
}

func (v Uint) Equals(other Value) bool {
	return v == other
}

func (v Uint) Less(other Value) bool {
	if v2, ok := other.(Uint); ok {
		return v < v2
	}
	return UintKind < other.Kind()
}

func (v Uint) Hash() hash.Hash {
	return getHash(v)
}

func (v Uint) WalkValues(cb ValueCallback) {
}

func (v Uint) WalkRefs(cb RefCallback) {
}

func (v Uint) typeOf() *Type {
	return UintType
}

func (v Uint) Kind() NomsKind {
	return UintKind
}

func (v Uint) valueReadWriter() ValueReadWriter {
	return nil
}

func (v Uint) writeTo(w nomsWriter) {
	UintKind.writeTo(w)
	w.writeUint(v)
}

func (v Uint) valueBytes() []byte {
	// We know the size of the buffer here so allocate it once.
	// UintKind, int (Varint), exp (Varint)
	buff := make([]byte, 1+2*binary.MaxVarintLen64)
	w := binaryNomsWriter{buff, 0}
	v.writeTo(&w)
	return buff[:w.offset]
}