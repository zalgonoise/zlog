package db

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	ErrShortWrite = errors.New("short write")
)

// DBWriter interface is a small wrapper for a Database Writer, to be used in a Logger
type DBWriter interface {
	io.WriteCloser
}

type multiWriteCloser struct {
	writers []io.WriteCloser
}

// Write method is a wraper for io.Writer, which calls this method across all
// WriteClosers.
//
// Writes are sequential. If an error is retrieved from a write, the remainder of the
// operation is cancelled, returning the error encountered.
func (m *multiWriteCloser) Write(p []byte) (n int, err error) {
	for _, w := range m.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = ErrShortWrite
			return
		}
	}
	return len(p), nil
}

// Close method is a wrapper for io.Closer, which calls this method across all
// WriteClosers.
//
// It only returns an error as per the io.Closer signature. However, unlike Write(),
// the operation will not halt if errors are retrieved when closing a Writer.
//
// Instead, the errors are collected and returned as one, if any. If there is only
// one error, the original error is returned. If there are multiple errors, a single
// error is returned, encapsulating all errors:
//
//     "multiple errors when closing writers: {errors...}"
func (m *multiWriteCloser) Close() error {
	var errs []error

	for _, w := range m.writers {
		err := w.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		if len(errs) == 1 {
			return errs[0]
		}

		var sb = strings.Builder{}

		for _, e := range errs {
			sb.WriteString(e.Error())
			sb.WriteString("; ")
		}

		return fmt.Errorf("multiple errors when closing writers: %s", sb.String())
	}

	return nil
}

// MultiWriteCloser creates a WriteCloser that duplicates its writes to all the
// provided writers, similar to the Unix tee(1) command.
//
// Similarly, it extends the support for io.Closer, by implementing a multiWriteCloser
// which can commit to the same Close call across all writers
//
// Each write is written to each listed writer, one at a time.
// If a listed writer returns an error, that overall write operation
// stops and returns the error; it does not continue down the list. This does not happen
// with the Close() call, which is intended to be sent to all writers regardless of errors
// retrieved. It will return a single error encapsulating all errors if existing
func MultiWriteCloser(dbs ...DBWriter) io.WriteCloser {
	allWriters := make([]io.WriteCloser, 0, len(dbs))
	for _, w := range dbs {
		if mw, ok := w.(*multiWriteCloser); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}
	return &multiWriteCloser{allWriters}
}