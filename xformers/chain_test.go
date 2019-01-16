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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/klmitch/mathom/object"
)

func TestXFormersImplementsInterfaces(t *testing.T) {
	assert.Implements(t, (*XFormer)(nil), XFormers{})
	assert.Implements(t, (*KeyXFormer)(nil), XFormers{})
	assert.Implements(t, (*DataXFormer)(nil), XFormers{})
}

type mockKey struct {
	mock.Mock
}

func (mk *mockKey) MaskKey(data []byte) []byte {
	args := mk.MethodCalled("MaskKey", data)

	masked := args.Get(0)
	if masked == nil {
		return nil
	}
	return masked.([]byte)
}

func (mk *mockKey) MaskName(name string) string {
	args := mk.MethodCalled("MaskName", name)

	return args.String(0)
}

type mockData struct {
	mock.Mock
}

func (mk *mockData) WrapData(data []byte, meta object.ObjMeta) []byte {
	args := mk.MethodCalled("WrapData", data, meta)

	wrapped := args.Get(0)
	if wrapped == nil {
		return nil
	}
	return wrapped.([]byte)
}

func (mk *mockData) UnwrapData(data []byte, meta object.ObjMeta) ([]byte, error) {
	args := mk.MethodCalled("UnwrapData", data, meta)

	unwrapped := args.Get(0)
	if unwrapped == nil {
		return nil, args.Error(1)
	}
	return unwrapped.([]byte), args.Error(1)
}

func TestXFormersPushKey(t *testing.T) {
	a := assert.New(t)
	obj := XFormers{}
	xf := &mockKey{}

	err := obj.Push(xf)

	a.NoError(err)
	a.Equal(obj, XFormers{xf})
}

func TestXFormersPushData(t *testing.T) {
	a := assert.New(t)
	obj := XFormers{}
	xf := &mockData{}

	err := obj.Push(xf)

	a.NoError(err)
	a.Equal(obj, XFormers{xf})
}

func TestXFormersPushBadType(t *testing.T) {
	a := assert.New(t)
	obj := XFormers{}
	xf := "invalid"

	err := obj.Push(xf)

	a.EqualError(err, "invalid xformer")
	a.Equal(obj, XFormers{})
}

func TestXFormersMaskKey(t *testing.T) {
	a := assert.New(t)
	k1 := &mockKey{}
	k2 := &mockKey{}
	d1 := &mockData{}
	d2 := &mockData{}
	obj := XFormers{k1, d1, k2, d2}
	k1.On("MaskKey", []byte("data")).Return([]byte("masked1"))
	k2.On("MaskKey", []byte("masked1")).Return([]byte("masked2"))

	result := obj.MaskKey([]byte("data"))

	a.Equal(result, []byte("masked2"))
	k1.AssertExpectations(t)
	k2.AssertExpectations(t)
	d1.AssertExpectations(t)
	d2.AssertExpectations(t)
}

func TestXFormersMaskName(t *testing.T) {
	a := assert.New(t)
	k1 := &mockKey{}
	k2 := &mockKey{}
	d1 := &mockData{}
	d2 := &mockData{}
	obj := XFormers{k1, d1, k2, d2}
	k1.On("MaskName", "name").Return("masked1")
	k2.On("MaskName", "masked1").Return("masked2")

	result := obj.MaskName("name")

	a.Equal(result, "masked2")
	k1.AssertExpectations(t)
	k2.AssertExpectations(t)
	d1.AssertExpectations(t)
	d2.AssertExpectations(t)
}

func TestXFormersWrapData(t *testing.T) {
	a := assert.New(t)
	k1 := &mockKey{}
	k2 := &mockKey{}
	d1 := &mockData{}
	d2 := &mockData{}
	obj := XFormers{k1, d1, k2, d2}
	d1.On("WrapData", []byte("data"), object.Meta(object.BLOCK, 0)).Return([]byte("wrapped1"))
	d2.On("WrapData", []byte("wrapped1"), object.Meta(object.BLOCK, 0)).Return([]byte("wrapped2"))

	result := obj.WrapData([]byte("data"), object.Meta(object.BLOCK, 0))

	a.Equal(result, []byte("wrapped2"))
	k1.AssertExpectations(t)
	k2.AssertExpectations(t)
	d1.AssertExpectations(t)
	d2.AssertExpectations(t)
}

func TestXFormersUnwrapDataHappyPath(t *testing.T) {
	a := assert.New(t)
	k1 := &mockKey{}
	k2 := &mockKey{}
	d1 := &mockData{}
	d2 := &mockData{}
	obj := XFormers{k1, d1, k2, d2}
	d1.On("UnwrapData", []byte("wrapped1"), object.Meta(object.BLOCK, 0)).Return([]byte("data"), nil)
	d2.On("UnwrapData", []byte("wrapped2"), object.Meta(object.BLOCK, 0)).Return([]byte("wrapped1"), nil)

	result, err := obj.UnwrapData([]byte("wrapped2"), object.Meta(object.BLOCK, 0))

	a.NoError(err)
	a.Equal(result, []byte("data"))
	k1.AssertExpectations(t)
	k2.AssertExpectations(t)
	d1.AssertExpectations(t)
	d2.AssertExpectations(t)
}

func TestXFormersUnwrapDataErrorPath(t *testing.T) {
	a := assert.New(t)
	k1 := &mockKey{}
	k2 := &mockKey{}
	d1 := &mockData{}
	d2 := &mockData{}
	obj := XFormers{k1, d1, k2, d2}
	d2.On("UnwrapData", []byte("wrapped2"), object.Meta(object.BLOCK, 0)).Return(nil, errors.New("an error"))

	result, err := obj.UnwrapData([]byte("wrapped2"), object.Meta(object.BLOCK, 0))

	a.EqualError(err, "an error")
	a.Nil(result)
	k1.AssertExpectations(t)
	k2.AssertExpectations(t)
	d1.AssertExpectations(t)
	d2.AssertExpectations(t)
}
