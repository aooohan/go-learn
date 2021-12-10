package cache

// PeerPicker get PeerGetter by a specific key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter get
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
