package log

import (
	"bytes"
	"time"

	"github.com/zalgonoise/zlog/log/event"
)

var mockBuffer = &bytes.Buffer{}
var mockLogger = struct {
	logger Logger
	buf    *bytes.Buffer
}{
	logger: New(
		WithPrefix("test-message"),
		WithFormat(FormatJSON),
		WithOut(mockBuffer),
	),
	buf: mockBuffer,
}

var mockChBufs = [][]*bytes.Buffer{
	{{}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}},
}

var mockLogLevelsOK = []event.LogLevel{
	event.LogLevel(0),
	event.LogLevel(1),
	event.LogLevel(2),
	event.LogLevel(3),
	event.LogLevel(4),
	event.LogLevel(5),
	event.LogLevel(9),
}

var mockLogLevelsNOK = []event.LogLevel{
	event.LogLevel(6),
	event.LogLevel(7),
	event.LogLevel(8),
	event.LogLevel(10),
	event.LogLevel(-1),
	event.LogLevel(200),
	event.LogLevel(500),
}

var mockPrefixes = []string{
	"test-logger",
	"test-prefix",
	"test-log",
	"test-service",
	"test-module",
	"test-logic",
}

var mockEmptyPrefixes = []string{
	"",
	"",
	"",
	"",
	"",
	"",
}

var mockMessages = []string{
	"message test #1",
	"message test #2",
	"message test #3",
	"message test #4",
	"message test #5",
	"mock message",
	"{ logger text in brackets }",
}

var mockFmtMessages = []struct {
	format string
	v      []interface{}
}{
	{
		format: "mockLogLevelsOK length: %v",
		v: []interface{}{
			len(mockLogLevelsOK),
		},
	},
	{
		format: "'Hello world!' in a list: %s",
		v: []interface{}{
			[]string{"H", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d", "!"},
		},
	},
	{
		format: "seven times three = %v",
		v: []interface{}{
			21,
		},
	},
}

var testObjects = []map[string]interface{}{
	{
		"testID": 0,
		"desc":   "this is a test with custom metadata",
		"content": map[string]interface{}{
			"nestLevel": 1,
			"data":      "nested object #1",
			"content": map[string]interface{}{
				"nestLevel": 2,
				"data":      "nested object #2",
			},
		},
		"date": time.Now().Format(time.RFC3339),
	}, {
		"testID": 1,
		"desc":   "this is a test with custom metadata",
		"content": map[string]interface{}{
			"nestLevel": 1,
			"data":      "nested object #1",
			"content": map[string]interface{}{
				"nestLevel": 2,
				"data":      "nested object #2",
				"content": map[string]interface{}{
					"nestLevel": 3,
					"data":      "nested object #3",
				},
			},
		},
		"date": time.Now().Format(time.RFC3339),
	}, {
		"testID": 2,
		"desc":   "this is a test with custom metadata",
		"content": map[string]interface{}{
			"nestLevel": 1,
			"data":      "nested object #1",
			"content": map[string]interface{}{
				"nestLevel": 2,
				"data":      "nested object #2",
				"content": map[string]interface{}{
					"nestLevel": 3,
					"data":      "nested object #3",
					"content": map[string]interface{}{
						"nestLevel": 4,
						"data":      "nested object #4",
					},
				},
			},
		},
		"date": time.Now().Format(time.RFC3339),
	}, {
		"testID": 3,
		"desc":   "this is a test with custom metadata",
		"content": map[string]interface{}{
			"nestLevel": 1,
			"data":      "nested object #1",
			"content": map[string]interface{}{
				"nestLevel": 2,
				"data":      "nested object #2",
				"content": map[string]interface{}{
					"nestLevel": 3,
					"data":      "nested object #3",
					"content": map[string]interface{}{
						"nestLevel": 4,
						"data":      "nested object #4",
						"content": map[string]interface{}{
							"nestLevel": 5,
							"data":      "nested object #5",
						},
					},
				},
			},
		},
		"date": time.Now().Format(time.RFC3339),
	},
}

var testEmptyObjects = []map[string]interface{}{
	nil,
	nil,
	{},
	{},
}