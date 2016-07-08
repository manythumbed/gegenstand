package packets

import (
	"bytes"
	"testing"
)

func TestWritePubAck(t *testing.T) {
	expected := []byte{4<<4 | 0, 2, 0x03, 0xFB}
	actual := WritePubAck(NewPacketId(1019))

	if !bytes.Equal(expected, actual) {
		t.Errorf("WritePubAck() = %v, expected = %v", actual, expected)
	}
}

func TestWritePubRec(t *testing.T) {
	expected := []byte{5<<4 | 0, 2, 0x03, 0xFC}
	actual := WritePubRec(NewPacketId(1020))

	if !bytes.Equal(expected, actual) {
		t.Errorf("WritePubRec() = %v, expected = %v", actual, expected)
	}
}

func TestWritePubRel(t *testing.T) {
	expected := []byte{6<<4 | 1<<1, 2, 0x03, 0xFD}
	actual := WritePubRel(NewPacketId(1021))

	if !bytes.Equal(expected, actual) {
		t.Errorf("WritePubRel() = %v, expected = %v", actual, expected)
	}
}

func TestWritePubComp(t *testing.T) {
	expected := []byte{7<<4 | 0, 2, 0x03, 0xFE}
	actual := WritePubComp(NewPacketId(1022))

	if !bytes.Equal(expected, actual) {
		t.Errorf("WritePubComp() = %v, expected = %v", actual, expected)
	}
}

func TestWriteSubAck(t *testing.T) {
	expected := []byte{9<<4 | 0, 6, 0x03, 0xE8, 0x01, 0x80, 0x02, 0x00}
	actual, err := WriteSubAck(NewPacketId(1000), SubSuccessOne, SubFailure, SubSuccessTwo, SubSuccessZero)

	if err != nil {
		t.Errorf("Unexpected error with WriteSubAck, err = %v", err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("WriteSubAck() = %v, expected = %v", actual, expected)
	}
}

func TestWriteUnsubscribe(t *testing.T) {
	expected := []byte{10<<4 | 1<<1, 0x0C, 0x03, 0xE9, 0x00, 0x03, 0x61, 0x2F, 0x62, 0x00, 0x03, 0x63, 0x2F, 0x64}
	actual, err := WriteUnsubscribe(NewPacketId(1001), "a/b", "c/d")

	if err != nil {
		t.Errorf("Unexpected error with WriteUnsubscribe, err = %v", err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("WriteUnsubscribe() = %v, expected = %v", actual, expected)
	}
}

func TestWriteUnsubAck(t *testing.T) {
	expected := []byte{11<<4 | 0, 2, 0x03, 0xFF}
	actual := WriteUnsubAck(NewPacketId(1023))

	if !bytes.Equal(expected, actual) {
		t.Errorf("WriteUnsubAck() = %v, expected = %v", actual, expected)
	}
}

func TestWritePingReq(t *testing.T) {
	expected := []byte{12<<4 | 0, 0}
	actual := WritePingReq()

	if !bytes.Equal(expected, actual) {
		t.Errorf("WritePingReq() = %v, expected = %v", actual, expected)
	}
}

func TestWritePingResp(t *testing.T) {
	expected := []byte{13<<4 | 0, 0}
	actual := WritePingResp()

	if !bytes.Equal(expected, actual) {
		t.Errorf("WritePingResp() = %v, expected = %v", actual, expected)
	}
}

func TestWriteDisconnect(t *testing.T) {
	expected := []byte{14<<4 | 0, 0}
	actual := WriteDisconnect()

	if !bytes.Equal(expected, actual) {
		t.Errorf("WriteDisconnect() = %v, expected = %v", actual, expected)
	}
}
