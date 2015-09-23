package hub

// MessageHub interface defines what methods the message hub should have
type MessageHub interface {
	// Start the hub and binds it to the specified address
	Start(laddr string) error
	// ClientIDs Returns client ids for the currently connected clients
	ClientIDs() []uint64
	// Stop stops the server
	Stop() error
}

// MessageEncoder defines the message encoder
type MessageEncoder func(receivers []uint64, payload []byte) []byte
