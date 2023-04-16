package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// Trivia :
// The gzip format and algorithm were initially developed by Jean-loup Gailly and Mark Adler in the 1990s
// as an alternative to the older Unix compress utility.
// The gzip format is based on the Deflate algorithm, which is a combination of LZ77 and Huffman coding.

// empty container used to group together related functionality

type GzipHandler struct {
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create a gziped response
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip") // tell the client this response needs to be decompressed
		}
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw) // returns a pointer to a new gzip.Writer
	return &WrappedResponseWriter{rw: rw, gw: gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
