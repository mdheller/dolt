// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"github.com/attic-labs/noms/go/hash"
)

var NullValue Null
var NullHash = getHash(NullValue)

// IsNull returns true if the value is nil, or if the value is of kind NULLKind
func IsNull(val Value) bool {
	return val == nil || val.Kind() == NullKind
}

// Int is a Noms Value wrapper around the primitive int32 type.
type Null byte

// Value interface
func (v Null) Value() Value {
	return v
}

func (v Null) Equals(other Value) bool {
	return other.Kind() == NullKind
}

func (v Null) Less(other Value) bool {
	return NullKind < other.Kind()
}

func (v Null) Hash() hash.Hash {
	return NullHash
}

func (v Null) WalkValues(cb ValueCallback) {
}

func (v Null) WalkRefs(cb RefCallback) {
}

func (v Null) typeOf() *Type {
	return NullType
}

func (v Null) Kind() NomsKind {
	return NullKind
}

func (v Null) valueReadWriter() ValueReadWriter {
	return nil
}

func (v Null) writeTo(w nomsWriter) {
	NullKind.writeTo(w)
}

func (v Null) valueBytes() []byte {
	buff := make([]byte, 1)
	w := binaryNomsWriter{buff, 0}
	v.writeTo(&w)
	return buff
}