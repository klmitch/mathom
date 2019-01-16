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
	"errors"

	"github.com/klmitch/mathom/object"
)

// XFormers describes a list of transformers to apply to database
// requests.  A transformer can mask keys or pointer names, and can
// wrap data in compression algorithms, encryption algorithms, and
// integrity protection algorithms.
type XFormers []XFormer

// Push adds a transformer to the list of transformers.  A transformer
// must meet either the KeyXFormer interface or the DataXFormer
// interface (or both).
func (xfs *XFormers) Push(xf XFormer) error {
	_, keyOk := xf.(KeyXFormer)
	_, dataOk := xf.(DataXFormer)
	if !keyOk && !dataOk {
		return errors.New("invalid xformer")
	}

	*xfs = append(*xfs, xf)

	return nil
}

// MaskKey masks a raw key.
func (xfs XFormers) MaskKey(key []byte) []byte {
	for _, xf := range xfs {
		if _, ok := xf.(KeyXFormer); ok {
			key = xf.(KeyXFormer).MaskKey(key)
		}
	}

	return key
}

// MaskName masks a pointer name.
func (xfs XFormers) MaskName(name string) string {
	for _, xf := range xfs {
		if _, ok := xf.(KeyXFormer); ok {
			name = xf.(KeyXFormer).MaskName(name)
		}
	}

	return name
}

// WrapData wraps the data.
func (xfs XFormers) WrapData(data []byte, meta object.ObjMeta) []byte {
	for _, xf := range xfs {
		if _, ok := xf.(DataXFormer); ok {
			data = xf.(DataXFormer).WrapData(data, meta)
		}
	}

	return data
}

// UnwrapData unwraps the data.
func (xfs XFormers) UnwrapData(data []byte, meta object.ObjMeta) ([]byte, error) {
	var err error
	for i := len(xfs) - 1; i >= 0; i-- {
		xf := xfs[i]
		if _, ok := xf.(DataXFormer); ok {
			data, err = xf.(DataXFormer).UnwrapData(data, meta)
			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}
