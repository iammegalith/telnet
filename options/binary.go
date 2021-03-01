package options

import "github.com/PatrickRudolph/telnet"

// BINARY Telnet binary Option - https://tools.ietf.org/html/rfc857

// BinaryTransmissionOption enables Binary transmission negotiation on a Server.
func BinaryTransmissionOption(c *telnet.Connection) telnet.Negotiator {
	return &BinaryHandler{client: false}
}

// BinaryHandler negotiates BINARY for a specific connection.
type BinaryHandler struct {
	client bool
}

// OptionCode returns with the code used to negotiate BINARY modes.
func (e *BinaryHandler) OptionCode() byte {
	return telnet.TeloptBINARY
}

// Offer is called when a new connection is initiated. It offers the handler
// an opportunity to advertise or request an option.
func (e *BinaryHandler) Offer(c *telnet.Connection) {
	if !e.client {
		c.Conn.Write([]byte{telnet.IAC, telnet.DO, e.OptionCode()})
	}
}

// HandleDo is called when an IAC DO command is received for this option,
// indicating the client is requesting the option to be enabled.
func (e *BinaryHandler) HandleDo(c *telnet.Connection) {

}

// HandleWill is called when an IAC WILL command is received for this
// option, indicating the client is willing to enable this option.
func (e *BinaryHandler) HandleWill(c *telnet.Connection) {

}

// HandleSB is called when a subnegotiation command is received for this
// option. body contains the bytes between `IAC SB <OptionCode>` and `IAC
// SE`.
func (e *BinaryHandler) HandleSB(c *telnet.Connection, body []byte) {

}
