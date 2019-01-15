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

func TestObjMetaEncrypted(t *testing.T) {
	a := assert.New(t)

	a.False(Meta(BLOCK, 0).Encrypted())
	a.False(Meta(BLOCK, COMPRESSED).Encrypted())
	a.True(Meta(BLOCK, ENCRYPTED).Encrypted())
	a.True(Meta(BLOCK, COMPRESSED|ENCRYPTED).Encrypted())
}

func TestObjMetaType(t *testing.T) {
	a := assert.New(t)

	a.Equal(Meta(BLOCK, 0).Type(), BLOCK)
	a.Equal(Meta(FILE, COMPRESSED).Type(), FILE)
	a.Equal(Meta(DIRECTORY, ENCRYPTED).Type(), DIRECTORY)
	a.Equal(Meta(SNAPSHOT, COMPRESSED|ENCRYPTED).Type(), SNAPSHOT)
}
