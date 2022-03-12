package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

const (
	bufferSize = 4096

	gUID            = "258EAFA5-E914-47DA- 95CA-C5AB0DC85B11"
	SecWebsocketKey = "Sec-WebSocket-Key"
)

type Conn interface {
	Close() error
}

type Websocket struct {
	conn           Conn
	buffReadWriter *bufio.ReadWriter
	header         http.Header
	status         uint16
}

func New(w http.ResponseWriter, req *http.Request) (*Websocket, error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		return nil, errors.New("webserver doesn't support http hijacking")
	}

	conn, buffReadWriter, err := hj.Hijack()
	if err != nil {
		return nil, err
	}

	return &Websocket{
		conn:           conn,
		buffReadWriter: buffReadWriter,
		header:         req.Header,
		status:         1000,
	}, nil
}

func getAcceptHash(key string) string {
	h := sha1.New()
	h.Write([]byte(key))
	h.Write([]byte(gUID))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (ws *Websocket) HandShake() error {
	hash := getAcceptHash(ws.header.Get(SecWebsocketKey))
	lines := []string{
		"HTTP/1.1 101 Web Socket Protocol Handshake",
		"Server: go/echoserver",
		"Upgrade: WebSocket",
		"Connection: Upgrade",
		"Sec-WebSocket-Accept: " + hash,
		"",
		"", // required for extra CRLF
	}

	return ws.write([]byte(strings.Join(lines, "\r\n")))
}

func (ws *Websocket) write(data []byte) error {
	if _, err := ws.buffReadWriter.Write(data); err != nil {
		return err
	}
	return ws.buffReadWriter.Flush()
}
