package handlers

// todo

//
//func TestSavePlainURL(t *testing.T) {
//	type want struct {
//		code        int
//		contentType string
//		bodyContent string
//	}
//
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "positive test #1",
//			want: want{
//				code:        201,
//				contentType: "text/plain",
//				bodyContent: "http://yandex.ru",
//			},
//		},
//		{
//			name: "negative test #1",
//			want: want{
//				code:        400,
//				contentType: "text/plain",
//				bodyContent: "yandex.ru",
//			},
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			body := strings.NewReader(test.want.bodyContent)
//			request := httptest.NewRequest(http.MethodPost, "/", body)
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			SavePlainURL(w, request)
//
//			res := w.Result()
//			// проверяем код ответа
//			assert.Equal(t, test.want.code, res.StatusCode)
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			resBody, err := io.ReadAll(res.Body)
//
//			require.NoError(t, err)
//			assert.NotEqual(t, "", string(resBody))
//			assert.Contains(t, res.Header.Get("Content-Type"), test.want.contentType)
//		})
//	}
//}
//
//func TestSaveJSONURL(t *testing.T) {
//	type want struct {
//		code        int
//		contentType string
//		bodyContent string
//	}
//
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "positive test #1",
//			want: want{
//				code:        201,
//				contentType: "application/json",
//				bodyContent: `{"url": "http://yandex.ru"}`,
//			},
//		},
//		{
//			name: "negative test #1",
//			want: want{
//				code:        400,
//				contentType: "text/plain",
//				bodyContent: `{"url": "yandex.ru"}`,
//			},
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			body := strings.NewReader(test.want.bodyContent)
//			request := httptest.NewRequest(http.MethodPost, "/api/shorten", body)
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			SaveJSONURL(w, request)
//
//			res := w.Result()
//			// проверяем код ответа
//			assert.Equal(t, test.want.code, res.StatusCode)
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			resBody, err := io.ReadAll(res.Body)
//
//			require.NoError(t, err)
//			assert.NotEqual(t, "", string(resBody))
//			assert.Contains(t, res.Header.Get("Content-Type"), test.want.contentType)
//		})
//	}
//}
//
//func TestGetURL(t *testing.T) {
//	type want struct {
//		code        int
//		contentType string
//	}
//
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "negative test #1",
//			want: want{
//				code:        400,
//				contentType: "text/plain",
//			},
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodPost, "/test", nil)
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			GetURL(w, request)
//
//			res := w.Result()
//			// проверяем код ответа
//			assert.Equal(t, test.want.code, res.StatusCode)
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			assert.Contains(t, res.Header.Get("Content-Type"), test.want.contentType)
//		})
//	}
//}
