package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

// Write -
func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.writer.Write(data)
}

// CompressResponse -
func CompressResponse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		acceptEncoding := ctx.GetHeader("Accept-Encoding")
		if !strings.Contains(acceptEncoding, "gzip") {
			ctx.Next()
			return
		}

		gz := gzip.NewWriter(ctx.Writer)
		defer gz.Close()

		ctx.Header("Content-Encoding", "gzip")
		ctx.Writer = &gzipWriter{
			ResponseWriter: ctx.Writer,
			writer:         gz,
		}
		ctx.Next()
	}
}

type gzipReader struct {
	io.ReadCloser
	reader *gzip.Reader
}

// Read -
func (g *gzipReader) Read(data []byte) (int, error) {
	return g.reader.Read(data)
}

// Close -
func (g *gzipReader) Close() error {
	return g.reader.Close()
}

// DecompressRequest -
func DecompressRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !strings.Contains(ctx.GetHeader("Content-Encoding"), "gzip") {
			ctx.Next()
			return
		}
		gz, err := gzip.NewReader(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		defer gz.Close()

		ctx.Request.Body = &gzipReader{
			ReadCloser: ctx.Request.Body,
			reader:     gz,
		}
		ctx.Next()
	}
}
