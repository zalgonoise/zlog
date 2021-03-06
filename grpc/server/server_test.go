package server

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/zalgonoise/zlog/log"
	"github.com/zalgonoise/zlog/log/event"
)

const maxTestWait time.Duration = time.Millisecond * 50

type failingWriter struct{}

func (failingWriter) Write(p []byte) (n int, err error) {
	if len(p) > 100 {
		return 100, errors.New("generic overflow error")
	}
	return 0, nil
}

func TestNew(t *testing.T) {
	module := "GRPCLogServer"
	funcname := "New()"

	_ = module
	_ = funcname

	type test struct {
		name    string
		cfg     []LogServerConfig
		wants   *GRPCLogServer
		optsLen int
	}

	var writers = []log.Logger{
		log.New(log.NilConfig),
		log.New(),
		log.New(log.SkipExit),
	}

	var tests = []test{
		{
			name:    "default config, no input",
			cfg:     []LogServerConfig{},
			wants:   New(defaultConfig),
			optsLen: 0,
		},
		{
			name: "with custom config (one entry)",
			cfg: []LogServerConfig{
				WithLogger(writers[0]),
			},
			wants:   New(WithLogger(writers[0])),
			optsLen: 0,
		},
		{
			name: "with custom config (two entries)",
			cfg: []LogServerConfig{
				WithLogger(writers[1]),
				WithServiceLoggerV(writers[1]),
			},
			wants: New(
				WithLogger(writers[1]),
				WithServiceLoggerV(writers[1]),
			),
			optsLen: 2,
		},
		{
			name:    "with nil input",
			cfg:     nil,
			wants:   New(defaultConfig),
			optsLen: 0,
		},
	}

	var verify = func(idx int, test test) {
		// wants := test.wants
		server := New(test.cfg...)

		// skip error channel deepequal
		if server.errCh == nil {
			t.Errorf(
				"#%v -- FAILED -- [%s] [%s] unexpected nil error channel -- action: %s",
				idx,
				module,
				funcname,
				test.name,
			)
			return
		}
		server.errCh = nil
		test.wants.errCh = nil

		// skip LogSv deepequal (tested in pb package)
		if server.logSv == nil {
			t.Errorf(
				"#%v -- FAILED -- [%s] [%s] unexpected nil log server -- action: %s",
				idx,
				module,
				funcname,
				test.name,
			)
			return
		}
		server.logSv = nil
		test.wants.logSv = nil

		// check options length first
		if len(server.opts) != test.optsLen {
			t.Errorf(
				"#%v -- FAILED -- [%s] [%s] grpc.ServerOption length mismatch error: wanted %v ; got %v -- action: %s",
				idx,
				module,
				funcname,
				test.optsLen,
				len(server.opts),
				test.name,
			)
			return
		}

		server.opts = nil
		test.wants.opts = nil

		if !reflect.DeepEqual(server, test.wants) {
			t.Errorf(
				"#%v -- FAILED -- [%s] [%s] output mismatch error: wanted %v ; got %v -- action: %s",
				idx,
				module,
				funcname,
				test.wants,
				server,
				test.name,
			)
			return
		}
	}

	for idx, test := range tests {
		verify(idx, test)
	}
}

func TestServe(t *testing.T) {
	module := "GRPCLogServer"
	funcname := "Serve()"

	_ = module
	_ = funcname

	type test struct {
		name string
		s    *GRPCLogServer
		ok   bool
	}

	var mockAddr = []string{
		"127.0.0.1:45051",
		"127.0.0.1:45052",
	}

	var bufs = []*bytes.Buffer{{}, {}}

	var writers = []log.Logger{
		log.New(log.WithOut(bufs[0]), log.SkipExit, log.CfgFormatJSON),
		log.New(log.WithOut(bufs[1]), log.SkipExit, log.CfgFormatJSON),
	}

	var tests = []test{
		{
			name: "working config",
			s: New(
				WithLogger(writers[0]),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[0]),
				WithGRPCOpts(),
			),
			ok: true,
		},
		{
			name: "invalid port error",
			s: New(
				WithLogger(writers[0]),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[1]),
				WithGRPCOpts(),
			),
		},
	}

	var reset = func() {
		for _, b := range bufs {
			b.Reset()
		}
	}

	var verify = func(idx int, test test) {
		defer test.s.Stop()
		defer reset()
		go test.s.Serve()
		go func() {
			for {
				select {
				case <-time.After(maxTestWait):
					return
				case err := <-test.s.errCh:
					if test.ok {
						t.Errorf(
							"#%v -- FAILED -- [%s] [%s] unexpected error: %v -- action: %s",
							idx,
							module,
							funcname,
							err,
							test.name,
						)
						return
					}
				}
			}
		}()
	}

	for idx, test := range tests {
		verify(idx, test)
	}
}

func TestHandleResponses(t *testing.T) {
	module := "GRPCLogServer"
	funcname := "handleResponses()"

	_ = module
	_ = funcname

	type test struct {
		name string
		s    *GRPCLogServer
		e    *event.Event
		ok   bool
	}

	var mockAddr = []string{
		"127.0.0.1:45053",
		"127.0.0.1:45054",
		"127.0.0.1:45055",
	}

	var bufs = []*bytes.Buffer{{}, {}}
	var failingL = log.New(log.WithOut(&failingWriter{}))

	var writers = []log.Logger{
		log.New(log.WithOut(bufs[0]), log.SkipExit, log.CfgFormatJSON),
		log.New(log.WithOut(bufs[1]), log.SkipExit, log.CfgFormatJSON),
	}

	var tests = []test{
		{
			name: "working config",
			s: New(
				WithLogger(writers[0]),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[0]),
				WithGRPCOpts(),
			),
			e:  event.New().Message("null").Build(),
			ok: true,
		},
		{
			name: "zero bytes writen error",
			s: New(
				WithLogger(failingL),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[1]),
				WithGRPCOpts(),
			),
			e: event.New().Message("null").Build(),
		},
		{
			name: "zero bytes writen error",
			s: New(
				WithLogger(failingL),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[2]),
				WithGRPCOpts(),
			),
			e: event.New().Message("very long message that will overflow the 100 byte threshold for the test failing writer").Build(),
		},
	}

	var handleErrors = func(idx int, test test) {
		for {
			select {
			case <-time.After(maxTestWait):
				return
			case err := <-test.s.errCh:
				if test.ok {
					t.Errorf(
						"#%v -- FAILED -- [%s] [%s] unexpected error: %v -- action: %s",
						idx,
						module,
						funcname,
						err,
						test.name,
					)
					return
				}
			case <-test.s.logSv.Done():
				return
			}
		}
	}

	var verify = func(idx int, test test) {
		defer test.s.Stop()
		go test.s.handleResponses(test.e)
		handleErrors(idx, test)
	}

	for idx, test := range tests {
		verify(idx, test)
	}
}

func TestHandleMesages(t *testing.T) {
	module := "GRPCLogServer"
	funcname := "handleMessages()"

	_ = module
	_ = funcname

	type test struct {
		name string
		s    *GRPCLogServer
		e    *event.Event
		ok   bool
	}

	var mockAddr = []string{
		"127.0.0.1:45056",
		"127.0.0.1:45057",
		"127.0.0.1:45058",
	}

	var bufs = []*bytes.Buffer{{}, {}}
	var failingL = log.New(log.WithOut(&failingWriter{}))

	var writers = []log.Logger{
		log.New(log.WithOut(bufs[0]), log.SkipExit, log.CfgFormatJSON),
		log.New(log.WithOut(bufs[1]), log.SkipExit, log.CfgFormatJSON),
	}

	var tests = []test{
		{
			name: "working config",
			s: New(
				WithLogger(writers[0]),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[0]),
				WithGRPCOpts(),
			),
			e:  event.New().Message("null").Build(),
			ok: true,
		},
		{
			name: "zero bytes writen error",
			s: New(
				WithLogger(failingL),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[1]),
				WithGRPCOpts(),
			),
			e: event.New().Message("null").Build(),
		},
		{
			name: "zero bytes writen error",
			s: New(
				WithLogger(failingL),
				WithServiceLoggerV(writers[1]),
				WithAddr(mockAddr[2]),
				WithGRPCOpts(),
			),
			e: event.New().Message("very long message that will overflow the 100 byte threshold for the test failing writer").Build(),
		},
	}

	var handleErrors = func(idx int, test test) {
		for {
			select {
			case <-time.After(maxTestWait):
				return
			case err := <-test.s.errCh:
				if test.ok {
					t.Errorf(
						"#%v -- FAILED -- [%s] [%s] unexpected error: %v -- action: %s",
						idx,
						module,
						funcname,
						err,
						test.name,
					)
					return
				}
			case <-test.s.logSv.Done():
				return
			}
		}
	}

	var verify = func(idx int, test test) {
		defer test.s.Stop()
		go test.s.handleMessages()
		test.s.logSv.MsgCh <- test.e
		handleErrors(idx, test)
	}

	for idx, test := range tests {
		verify(idx, test)
	}
}
