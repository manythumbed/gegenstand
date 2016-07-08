package packets

import (
	"bytes"
	"errors"
)

type PacketIdentifier [2]byte

func NewPacketId(i uint16) PacketIdentifier {
	return PacketIdentifier{uint8(i >> 8), uint8(i & 0xFF)}
}

const (
	connect     byte = 1
	connack          = 2
	publish          = 3
	puback           = 4
	pubrec           = 5
	pubrel           = 6
	pubcomp          = 7
	subscribe        = 8
	suback           = 9
	unsubscribe      = 10
	unsuback         = 11
	pingreq          = 12
	pingresp         = 13
	disconnect       = 14
)

func WritePubAck(id PacketIdentifier) []byte {
	return packetWithIdentifier(puback, id)
}

func WritePubRec(id PacketIdentifier) []byte {
	return packetWithIdentifier(pubrec, id)
}

func WritePubRel(id PacketIdentifier) []byte {
	return packetWithIdentifierAndFlags(pubrel, 1<<1, id)
}

func WritePubComp(id PacketIdentifier) []byte {
	return packetWithIdentifier(pubcomp, id)
}

func WriteUnsubAck(id PacketIdentifier) []byte {
	return packetWithIdentifier(unsuback, id)
}

func WritePingReq() []byte {
	return emptyPacket(pingreq)
}

func WritePingResp() []byte {
	return emptyPacket(pingresp)
}

func WriteDisconnect() []byte {
	return emptyPacket(disconnect)
}

func packetWithIdentifier(packetType byte, id PacketIdentifier) []byte {
	return packetWithIdentifierAndFlags(packetType, 0, id)
}

func packetWithIdentifierAndFlags(packetType byte, flags byte, id PacketIdentifier) []byte {
	packet := bytes.Buffer{}
	packet.Write(fixedHeader(packetType, flags, []byte{2}))
	packet.WriteByte(id[0])
	packet.WriteByte(id[1])

	return packet.Bytes()
}

func emptyPacket(packetType byte) []byte {
	return fixedHeader(packetType, 0, []byte{0})
}

func fixedHeader(packetType byte, flags byte, remainingLength []byte) []byte {
	buffer := bytes.Buffer{}
	buffer.WriteByte(packetType<<4 | flags)
	buffer.Write(remainingLength)

	return buffer.Bytes()
}

func remainingLength(length int) ([]byte, error) {
	if length < 0 || length > 268435455 {
		return nil, errors.New("Invalid remaining length")
	}

	return nil, nil
}
