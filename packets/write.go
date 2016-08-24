package packets

import (
	"bytes"
	"errors"
	"github.com/manythumbed/gegenstand/protocol"
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

const (
	SubSuccessZero byte = 0x00
	SubSuccessOne       = 0x01
	SubSuccessTwo       = 0x02
	SubFailure          = 0x80
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

// TODO look at creating a common shared package for Subscription, Qos.
func WriteSubscribe(id PacketIdentifier, topics ...protocol.Subscription) ([]byte, error) {

	return nil, nil
}

func WriteSubAck(id PacketIdentifier, codes ...byte) ([]byte, error) {
	if len(codes) == 0 {
		return nil, errors.New("No return codes have been provided for suback packet")
	}

	l, err := remainingLength(len(codes) + 2)
	if err != nil {
		return nil, errors.New("Too many return codes")
	}

	packet := bytes.Buffer{}
	packet.Write(fixedHeader(suback, 0, l))
	packet.WriteByte(id[0])
	packet.WriteByte(id[1])
	packet.Write(codes)

	return packet.Bytes(), nil
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
