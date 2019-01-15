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

// Encrypted returns true if the content flags indicate the content is
// encrypted.
func (om ObjMeta) Encrypted() bool {
	return uint8(om)&ENCRYPTED == ENCRYPTED
}

// Type returns the object type.
func (om ObjMeta) Type() ObjType {
	return ObjType(uint8(om) & TYPEMASK)
}
