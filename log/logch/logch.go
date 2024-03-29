package logch

import (
	"github.com/zalgonoise/zlog/log"
	"github.com/zalgonoise/zlog/log/event"
)

// ChanneledLogger interface defines the behavior of a channeled logger in a goroutine.
//
// This interface is a bit more specific in the sense that it has also methods to interact
// with the goroutine, and not just to spawn it and retrieve the needed channels.
type ChanneledLogger interface {
	Log(msg ...*event.Event)
	Close()
	Channels() (logCh chan *event.Event, done chan struct{})
}

// LogChannel struct defines what a minimal logging channel must contain:
//   - a channel to receive pointers to event.Event
//   - a channel to receive a done signal (to close the goroutine)
type LogChannel struct {
	logCh chan *event.Event
	done  chan struct{}
}

// New function is a helper to spawn a channeled logger function and
// channel, for an existing Logger interface.
//
// Instead of implementing the logic below everytime, this function can be
// used to spawn a go routine and use its channel to send messages:
//
//
//     logger := log.New(log.WithPrefix("logger"), log.CfgTextFormat)
//     logCh := logch.New(logger)
//
//	   // then, either the "classic" channeled message approach:
//	   ch, done := logCh.Channels()
//
//     ch <- event.New().Level(log.Level_trace).Message("test message").Build()
//
//     // or using the embeded method
//     logCh.Log(event.New().Message("this works too").Build())
//
//     // and finally stop the goroutine (if needed)
//     logCh.Close()
//     // or
//     done <- struct{}{}
func New(logger log.Logger) (logCh ChanneledLogger) {

	msgCh := make(chan *event.Event)
	done := make(chan struct{})

	logCh = &LogChannel{
		logCh: msgCh,
		done:  done,
	}

	go func(done chan struct{}) {
		for {
			select {
			case msg := <-msgCh:
				logger.Log(msg)
			case <-done:
				return
			}

		}
	}(done)

	return
}

// Log method will take in any number of pointers to event.Event, and iterating through each of them,
// pushing them to the LogMessage channel.
//
// As these messages are queued, they will be then printed within the spawned goroutine, using a Logger.Log()
// method call
//
// This method is a wrapper for not having to call the Channels() method, and then working with these separately
//
//     logger := log.New(log.WithPrefix("logger"), log.CfgTextFormat)
//     logCh := NewLogCh(logger)
//
//     logCh.Log(
//       event.New().Message("this works too").Build(),
//       event.New().Message("with many messages").Build(),
//     )
//
//     logCh.Close()
//
func (c LogChannel) Log(msg ...*event.Event) {
	if len(msg) == 0 {
		return
	}

	for _, m := range msg {
		if m != nil {
			c.logCh <- m
		}
	}
}

// Close method will send a signal (an empty `struct{}`) to the done channel, triggering the spawned goroutine to
// return
//
//     logger := log.New(log.WithPrefix("logger"), log.CfgTextFormat)
//     logCh := NewLogCh(logger)
//
//     logCh.Log(
//       event.New().Message("this works too").Build(),
//       event.New().Message("with many messages").Build(),
//     )
//
//     logCh.Close()
//
func (c LogChannel) Close() {
	c.done <- struct{}{}
}

// Channels method will return the LogMessage channel and the done channel, so that they can be used
// directly with the same channel messaging patterns
//
//     logger := log.New(log.WithPrefix("logger"), log.CfgTextFormat)
//     logCh := log.NewLogCh(logger)
//
//	   ch, done := logCh.Channels()
//
//     ch <- event.New().Level(log.Level_trace).Message("test message").Build()
//
//     done <- struct{}{}
//
func (c LogChannel) Channels() (logCh chan *event.Event, done chan struct{}) {
	return c.logCh, c.done
}
