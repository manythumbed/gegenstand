package protocol

type Qos byte

const (
	AtMostOnce  Qos = 0
	AtLeastOnce     = 1
	ExactlyOnce     = 2
)

type Subscription struct {
	topic   string
	quality Qos
}
