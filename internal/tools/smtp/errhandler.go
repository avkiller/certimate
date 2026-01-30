package smtp

import (
	"bytes"
	"errors"
	"io"
	"net/textproto"
)

// REF: https://github.com/wneessen/go-mail/wiki/Error-Registry
type wQQMailQuitErrorHandler struct{}

func (q *wQQMailQuitErrorHandler) HandleError(_, _ string, conn *textproto.Conn, err error) error {
	var tpErr textproto.ProtocolError
	if errors.As(err, &tpErr) {
		if len(tpErr.Error()) < 16 {
			return err
		}
		if !bytes.Equal([]byte(tpErr.Error()[16:]), []byte("\x00\x00\x00\x1a\x00\x00\x00")) {
			return err
		}
		_, _ = io.ReadFull(conn.R, make([]byte, 8))
		return nil
	}
	return err
}
