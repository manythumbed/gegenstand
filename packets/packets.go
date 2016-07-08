package packets

import (
	"bytes"
	"errors"
	"fmt"
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

func WriteUnsubscribe(id PacketIdentifier, topics ...string) ([]byte, error) {
	if len(topics) == 0 {
		return nil, errors.New("No topics have been provided for unsubscribe packet")
	}

	encodedTopics := bytes.Buffer{}
	for _, t := range topics {
		s, err := encode(t)
		if err != nil {
			return nil, errors.New("Topic is too long")
		}

		encodedTopics.Write(s)
	}

	l, err := remainingLength(encodedTopics.Len() + 2)
	if err != nil {
		return nil, errors.New("encoded topics are too long for unsubscribe packet")
	}

	fmt.Println(l)

	packet := bytes.Buffer{}
	packet.Write(fixedHeader(unsubscribe, 1<<1, l))
	packet.WriteByte(id[0])
	packet.WriteByte(id[1])
	packet.Write(encodedTopics.Bytes())

	return packet.Bytes(), nil
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

	bytes := bytes.Buffer{}

	for {
		d := byte(length % 128)
		length /= 128
		if length > 0 {
			d |= 0x80
		}
		bytes.WriteByte(d)

		if length == 0 {
			break
		}
	}

	return bytes.Bytes(), nil
}
