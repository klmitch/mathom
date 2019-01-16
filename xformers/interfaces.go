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

package xformers

import (
	"github.com/klmitch/mathom/object"
)

// XFormer is a transformer for keys and data in the store.  This can
// come in one of two flavors: a KeyXFormer or a DataXFormer.
type XFormer interface{}

// KeyXFormer transforms keys in the database by masking them in some
// fashion.  This enhances security by making it difficult for an
// attacker to discover the data by brute-forcing or otherwise
// breaking the hashing used to generate a key.
type KeyXFormer interface {
	// MaskKey masks a raw key.
	MaskKey([]byte) []byte

	// MaskName masks a pointer name.
	MaskName(string) string
}

// DataXFormer transforms data in the database in some fashion.  This
// could be in the form of applying data compression, or it could be
// encrypting the data.
type DataXFormer interface {
	// WrapData wraps the data in the transform.
	WrapData([]byte, object.ObjMeta) []byte

	// UnwrapData unwraps the data in the transform.  If the
	// unwrapping cannot be performed, an error is returned.
	UnwrapData([]byte, object.ObjMeta) ([]byte, error)
}
