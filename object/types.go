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

package object

// ObjType classifies the type of an object.
type ObjType uint8

// Defined object types.
const (
	BLOCK     ObjType = iota // Object is a Block object.
	FILE                     // Object is a File object.
	DIRECTORY                // Object is a Directory object.
	SNAPSHOT                 // Object is a Snapshot object.
)

// Types maps object type values to strings.
var types = []string{
	"BLOCK",
	"FILE",
	"DIRECTORY",
	"SNAPSHOT",
}

// String returns the string representation of the object type.
func (t ObjType) String() string {
	return types[t]
}

// Defined object flags
const (
	COMPRESSED uint8               = 1 << (7 - iota) // Content is compressed.
	ENCRYPTED                                        // Content is encrypted.
	TYPEMASK   = ^uint8(0) >> iota                   // Mask for object type.
)

// ObjMeta combines the object type with the content flags.
type ObjMeta uint8

// Meta constructs an ObjMeta from the object type and the content
// flags.
func Meta(typ ObjType, flags uint8) ObjMeta {
	return ObjMeta(flags | uint8(typ))
}

// Compressed returns true if the content flags indicate the content
// is compressed.
func (om ObjMeta) Compressed() bool {
	return uint8(om)&COMPRESSED == COMPRESSED
}

// SetCompressed sets the COMPRESSED flag and returns the meta for
// convenience.
func (om *ObjMeta) SetCompressed() ObjMeta {
	*om = ObjMeta(uint8(*om) | COMPRESSED)
	return *om
}

// Encrypted returns true if the content flags indicate the content is
// encrypted.
func (om ObjMeta) Encrypted() bool {
	return uint8(om)&ENCRYPTED == ENCRYPTED
}

// SetEncrypted sets the ENCRYPTED flag and returns the meta for
// convenience.
func (om *ObjMeta) SetEncrypted() ObjMeta {
	*om = ObjMeta(uint8(*om) | ENCRYPTED)
	return *om
}

// Type returns the object type.
func (om ObjMeta) Type() ObjType {
	return ObjType(uint8(om) & TYPEMASK)
}
