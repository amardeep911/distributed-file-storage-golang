package p2p

// peer is na interface which represent a connected remote node
type Peer interface {
}

// Transport is anything that handles the communication
// between the nodes. This can be of the form (Tcp, Udp, Websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
