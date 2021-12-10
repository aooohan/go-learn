package cache

// ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (b ByteView) Len() int {
	return len(b.b)
}

func (b ByteView) String() string {
	return string(b.b)
}

// ByteSlice returns a copy of the data as a byte slice
func (b ByteView) ByteSlice() []byte {
	return cloneBytes(b.b)
}

// cloneBytes deeply copy
func cloneBytes(b []byte) []byte {
	bytes := make([]byte, len(b))
	copy(bytes, b)
	return bytes
}
