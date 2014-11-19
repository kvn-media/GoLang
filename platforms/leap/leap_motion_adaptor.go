package leap

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/hybridgroup/gobot"
	"io"
)

var _ gobot.AdaptorInterface = (*LeapMotionAdaptor)(nil)

type LeapMotionAdaptor struct {
	gobot.Adaptor
	ws      io.ReadWriteCloser
	connect func(*LeapMotionAdaptor) (err error)
}

// NewLeapMotionAdaptor creates a new leap motion adaptor using specified name and port
func NewLeapMotionAdaptor(name string, port string) *LeapMotionAdaptor {
	return &LeapMotionAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"LeapMotionAdaptor",
			port,
		),
		connect: func(l *LeapMotionAdaptor) (err error) {
			ws, err := websocket.Dial(
				fmt.Sprintf("ws://%v/v3.json", l.Port()),
				"",
				fmt.Sprintf("http://%v", l.Port()),
			)
			if err != nil {
				return err
			}
			l.ws = ws
			return
		},
	}
}

// Connect returns true if connection to leap motion is established succesfully
func (l *LeapMotionAdaptor) Connect() error {
	return l.connect(l)
}

// Finalize ends connection to leap motion
func (l *LeapMotionAdaptor) Finalize() error { return nil }
