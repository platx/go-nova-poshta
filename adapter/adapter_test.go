package adapter

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/platx/go-nova-poshta/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestHTTPAdapter(t *testing.T) {
	t.Parallel()

	httpClientMock := &mockHTTPClient{}

	testCases := []struct {
		name   string
		format Format

		apiKey       string
		modelName    string
		calledMethod string
		reqProps     any
		res          any

		fakeResponseCode int
		fakeResponseErr  error
		fakeResponseBody response

		expectedErr error
	}{
		{
			name:         "WithRequestPropsAndResponse/Success",
			format:       FormatJSON,
			apiKey:       testdata.FakeApiKey,
			modelName:    "testModel",
			calledMethod: "testMethod",
			reqProps: map[string]interface{}{
				"prop1": "value1",
				"prop2": "value2",
			},

			fakeResponseCode: http.StatusOK,
			fakeResponseErr:  nil,
			fakeResponseBody: response{
				Success:      true,
				Data:         map[string]interface{}{"data": "fake-data"},
				Errors:       nil,
				Warnings:     nil,
				Info:         nil,
				MessageCodes: nil,
				ErrorCodes:   nil,
				WarningCodes: nil,
				InfoCodes:    nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s/%s", strings.ToUpper(string(tc.format)), tc.name), func(t *testing.T) {
			fakeReq := request{
				ApiKey:           tc.apiKey,
				ModelName:        tc.modelName,
				CalledMethod:     tc.calledMethod,
				MethodProperties: tc.reqProps,
			}

			c := NewAdapter(CreateConfig(
				tc.apiKey,
				WithHTTPClient(httpClientMock),
				WithFormat(tc.format),
			))

			assertUrl(t, tc.format, c.(*httpAdapter).baseUrl)
			assertSerializer(t, tc.format, c.(*httpAdapter).serializer)
			assert.Equal(t, fakeReq.ApiKey, c.(*httpAdapter).apiKey)
			assert.IsType(t, httpClientMock, c.(*httpAdapter).http)

			httpClientMock.mockDo = func(req *http.Request) (*http.Response, error) {
				decReq := &request{}

				if fakeReq.MethodProperties != nil {
					decReq.MethodProperties = reflect.New(reflect.TypeOf(fakeReq.MethodProperties))
				}

				unmarshalData(t, tc.format, streamToByte(t, req.Body), decReq)

				assert.Equal(t, http.MethodPost, req.Method)
				assertUrl(t, tc.format, req.URL.String())
				assert.Equal(t, fakeReq.ApiKey, decReq.ApiKey)
				assert.Equal(t, fakeReq.ModelName, decReq.ModelName)
				assert.Equal(t, fakeReq.CalledMethod, decReq.CalledMethod)
				assert.Equal(t, fakeReq.MethodProperties, decReq.MethodProperties)

				return createHttpResponse(t, tc.format, tc.fakeResponseCode, tc.fakeResponseBody), tc.fakeResponseErr
			}

			err := c.Req(fakeReq.ModelName, fakeReq.CalledMethod, fakeReq.MethodProperties, tc.res)

			if tc.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

	t.Run("Unknown", func(t *testing.T) {
		httpClient, expectedFormat, apiKey := &mockHTTPClient{}, Format("unknown"), testdata.FakeApiKey

		require.Panics(t, func() {
			NewAdapter(CreateConfig(
				apiKey,
				WithHTTPClient(httpClient),
				WithFormat(expectedFormat),
			))
		})
	})
}

type mockHTTPClient struct {
	mockDo func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.mockDo(req)
}

func unmarshalData(t *testing.T, format Format, data []byte, res any) {
	var err error

	switch format {
	case FormatJSON:
		err = json.Unmarshal(data, res)
	case FormatXML:
		err = xml.Unmarshal(data, res)
	}

	require.NoError(t, err)
}

func marshalData(t *testing.T, format Format, data any) []byte {
	var (
		encData []byte
		err     error
	)

	switch format {
	case FormatJSON:
		encData, err = json.Marshal(data)
	case FormatXML:
		encData, err = xml.Marshal(data)
	}

	require.NoError(t, err)

	return encData
}

func dataToStream(t *testing.T, format Format, data any) io.ReadCloser {
	return io.NopCloser(bytes.NewBuffer(marshalData(t, format, data)))
}

func streamToByte(t *testing.T, stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	require.NoError(t, err)

	return buf.Bytes()
}

func assertUrl(t *testing.T, format Format, url string) {
	switch format {
	case FormatJSON:
		assert.Equal(t, "https://api.novaposhta.ua/v2.0/json/", url)
	case FormatXML:
		assert.Equal(t, "https://api.novaposhta.ua/v2.0/xml/", url)
	}
}

func assertSerializer(t *testing.T, format Format, serializer serializer) {
	switch format {
	case FormatJSON:
		assert.IsType(t, &serializerJSON{}, serializer)
	case FormatXML:
		assert.IsType(t, &serializerXML{}, serializer)
	}
}

func createHttpResponse(t *testing.T, format Format, status int, data any) *http.Response {
	return &http.Response{
		Status:     http.StatusText(status),
		StatusCode: status,
		Body:       dataToStream(t, format, data),
	}
}
