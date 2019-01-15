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

// Store is an interface for all possible mathom storage systems.  It
// has distinct operations for objects in the store--typically keyed
// by their checksum--and pointers, which allow human-readable names
// to be associated with a given object key.  Mathom storage systems
// can also be wrapped by objects which modify it in some fashion,
// such as encrypting the data in the store.
type Store interface {
	// Tests if an object with the specified key exists.
	ObjExists(key []byte) bool

	// Creates an object with the specified key and data.
	// Additional metadata describes the type of the object.
	ObjCreate(key, data []byte, meta object.ObjMeta) error

	// Get an object with the specified key.  The data and
	// metadata are returned.
	ObjGet(key []byte) ([]byte, object.ObjMeta, error)

	// Delete an object with the specified key.
	ObjDelete(key []byte) error

	// Create or update a named pointer, mapping it to the
	// specified key.
	PtrSet(name, key []byte) error

	// Get the named pointer and return the key it points to.
	PtrGet(name []byte) ([]byte, error)

	// Delete a pointer.
	PtrDelete(name []byte) error
}
