package packets

import (
	"reflect"
	"testing"
)

func TestWritePubAck(t *testing.T) {
	expected := []byte{4<<4 | 0, 2, 0x03, 0xFB}
	actual := WritePubAck(NewPacketId(1019))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WritePubAck() = %v, expected = %v", actual, expected)
	}
}

func TestWritePubRec(t *testing.T) {
	expected := []byte{5<<4 | 0, 2, 0x03, 0xFC}
	actual := WritePubRec(NewPacketId(1020))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WritePubRec() = %v, expected = %v", actual, expected)
	}
}

func TestWritePubRel(t *testing.T) {
	expected := []byte{6<<4 | 1<<1, 2, 0x03, 0xFD}
	actual := WritePubRel(NewPacketId(1021))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WritePubRel() = %v, expected = %v", actual, expected)
	}
}

func TestWritePubComp(t *testing.T) {
	expected := []byte{7<<4 | 0, 2, 0x03, 0xFE}
	actual := WritePubComp(NewPacketId(1022))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WritePubComp() = %v, expected = %v", actual, expected)
	}
}

func TestWriteUnsubAck(t *testing.T) {
	expected := []byte{11<<4 | 0, 2, 0x03, 0xFF}
	actual := WriteUnsubAck(NewPacketId(1023))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WriteUnsubAck() = %v, expected = %v", actual, expected)
	}
}

func TestWritePingReq(t *testing.T) {
	expected := []byte{12<<4 | 0, 0}
	actual := WritePingReq()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WritePingReq() = %v, expected = %v", actual, expected)
	}
}

func TestWritePingResp(t *testing.T) {
	expected := []byte{13<<4 | 0, 0}
	actual := WritePingResp()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WritePingResp() = %v, expected = %v", actual, expected)
	}
}

func TestWriteDisconnect(t *testing.T) {
	expected := []byte{14<<4 | 0, 0}
	actual := WriteDisconnect()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("WriteDisconnect() = %v, expected = %v", actual, expected)
	}
}
