// Copyright (C) 2019 Kevin L. Mitchell <klmitch@mit.edu>
//
// Licensed under the Apache License, Version 2.0 (the "License"); you
// may not use this file except in compliance with the License.  You
// may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.  See the License for the specific language governing
// permissions and limitations under the License.

package store

import (
	"github.com/klmitch/mathom/object"
)

// ObjIter is an interface for iterating over objects in the database.
type ObjIter interface {
	// Item retrieves the current item from the iterator.  This
	// includes the object key, the object data, and the object
	// metadata.
	Item() ([]byte, []byte, object.ObjMeta)

	// Next advances the iterator to the next item.
	Next()
}

// PtrIter is an interface for iterating over pointers in the
// database.
type PtrIter interface {
	// Item retrieves the current item from the iterator.  This
	// includes the item name and the key to which it points.
	Item() (string, []byte)

	// Next advances the iterator to the next item.
	Next()
}

// Store is an interface for all possible mathom storage systems.  It
// has distinct operations for objects in the store--typically keyed
// by their checksum--and pointers, which allow human-readable names
// to be associated with a given object key.  Mathom storage systems
// can also be wrapped by objects which modify it in some fashion,
// such as encrypting the data in the store.
type Store interface {
	// ObjExists tests if an object with the specified key exists.
	ObjExists(key []byte) bool

	// ObjCreate creates an object with the specified key and
	// data.  Additional metadata describes the type of the
	// object.
	ObjCreate(key, data []byte, meta object.ObjMeta) error

	// ObjGet gets an object with the specified key.  The data and
	// metadata are returned.
	ObjGet(key []byte) ([]byte, object.ObjMeta, error)

	// ObjDelete deletes an object with the specified key.
	ObjDelete(key []byte) error

	// IterObjs iterates over all objects in the store.  It is not
	// safe to modify the store during iteration.
	IterObjs() ObjIter

	// PtrSet creates or updates a named pointer, mapping it to
	// the specified key.
	PtrSet(name string, key []byte) error

	// PtrGet gets the named pointer and return the key it points
	// to.
	PtrGet(name string) ([]byte, error)

	// PtrDelete deletes a pointer.
	PtrDelete(name string) error

	// IterPtrs iterates over all pointers in the store.  It is
	// not safe to modify the store during iteration.
	IterPtrs() PtrIter
}
