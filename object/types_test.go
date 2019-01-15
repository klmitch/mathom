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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjTypeString(t *testing.T) {
	a := assert.New(t)

	result := BLOCK.String()

	a.Equal(result, "BLOCK")
}

func TestMetaBase(t *testing.T) {
	a := assert.New(t)

	result := Meta(BLOCK, 0)

	a.Equal(result, ObjMeta(0))
}

func TestMetaAlt(t *testing.T) {
	a := assert.New(t)

	result := Meta(DIRECTORY, COMPRESSED|ENCRYPTED)

	a.Equal(result, ObjMeta(0xc2))
}

func TestObjMetaCompressed(t *testing.T) {
	a := assert.New(t)

	a.False(Meta(BLOCK, 0).Compressed())
	a.True(Meta(BLOCK, COMPRESSED).Compressed())
	a.False(Meta(BLOCK, ENCRYPTED).Compressed())
	a.True(Meta(BLOCK, COMPRESSED|ENCRYPTED).Compressed())
}

func TestObjMetaSetCompressed(t *testing.T) {
	a := assert.New(t)
	obj := Meta(BLOCK, 0)

	result := obj.SetCompressed()

	a.Equal(obj, result)
	a.True(result.Compressed())
}

func TestObjMetaEncrypted(t *testing.T) {
	a := assert.New(t)

	a.False(Meta(BLOCK, 0).Encrypted())
	a.False(Meta(BLOCK, COMPRESSED).Encrypted())
	a.True(Meta(BLOCK, ENCRYPTED).Encrypted())
	a.True(Meta(BLOCK, COMPRESSED|ENCRYPTED).Encrypted())
}

func TestObjMetaSetEncrypted(t *testing.T) {
	a := assert.New(t)
	obj := Meta(BLOCK, 0)

	result := obj.SetEncrypted()

	a.Equal(obj, result)
	a.True(result.Encrypted())
}

func TestObjMetaType(t *testing.T) {
	a := assert.New(t)

	a.Equal(Meta(BLOCK, 0).Type(), BLOCK)
	a.Equal(Meta(FILE, COMPRESSED).Type(), FILE)
	a.Equal(Meta(DIRECTORY, ENCRYPTED).Type(), DIRECTORY)
	a.Equal(Meta(SNAPSHOT, COMPRESSED|ENCRYPTED).Type(), SNAPSHOT)
}
