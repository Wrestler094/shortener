package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// Compressor - middleware для сжатия HTTP-ответов и распаковки HTTP-запросов.
// Поддерживает сжатие gzip для ответов и распаковку gzip для запросов.
// Сжимает только ответы с типами контента application/json и text/html.
func Compressor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции
		ow := w

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		acceptEncoding := r.Header.Get("Accept-Encoding")
		isGzipSupported := strings.Contains(acceptEncoding, "gzip")
		contentType := r.Header.Get("Content-Type")
		isCompressibleContentType := strings.Contains(contentType, "application/json") || strings.Contains(contentType, "text/html")

		if isGzipSupported && isCompressibleContentType {
			// оборачиваем оригинальный http.ResponseWriter новым с поддержкой сжатия
			cw := newCompressorWriter(w)
			// меняем оригинальный http.ResponseWriter на новый
			ow = cw
			// не забываем отправить клиенту все сжатые данные после завершения middleware
			defer cw.Close()
		}

		// проверяем, что клиент отправил серверу сжатые данные в формате gzip
		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			// оборачиваем тело запроса в io.Reader с поддержкой декомпрессии
			cr, err := newCompressorReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// меняем тело запроса на новое
			r.Body = cr
			defer cr.Close()
		}

		// передаём управление хендлеру
		next.ServeHTTP(ow, r)
	})
}

// compressorWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
// сжимать передаваемые данные и выставлять правильные HTTP-заголовки
type compressorWriter struct {
	w  http.ResponseWriter // Оригинальный ResponseWriter
	zw *gzip.Writer        // Writer для сжатия данных
}

// newCompressorWriter создает новый экземпляр compressorWriter
func newCompressorWriter(w http.ResponseWriter) *compressorWriter {
	return &compressorWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

// Header возвращает HTTP-заголовки ответа
func (c *compressorWriter) Header() http.Header {
	return c.w.Header()
}

// Write записывает сжатые данные в ответ
func (c *compressorWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

// WriteHeader устанавливает код ответа и заголовок Content-Encoding
func (c *compressorWriter) WriteHeader(statusCode int) {
	c.w.Header().Set("Content-Encoding", "gzip")
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера
func (c *compressorWriter) Close() error {
	return c.zw.Close()
}

// compressorReader реализует интерфейс io.ReadCloser и позволяет прозрачно для сервера
// декомпрессировать получаемые от клиента данные
type compressorReader struct {
	r  io.ReadCloser // Оригинальный Reader
	zr *gzip.Reader  // Reader для распаковки данных
}

// newCompressorReader создает новый экземпляр compressorReader
func newCompressorReader(r io.ReadCloser) (*compressorReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressorReader{
		r:  r,
		zr: zr,
	}, nil
}

// Read читает и распаковывает данные из запроса
func (c compressorReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

// Close закрывает оба reader'а
func (c *compressorReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}
