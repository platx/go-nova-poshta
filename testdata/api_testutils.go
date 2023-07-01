package testdata

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatXML  Format = "xml"
)

const (
	FakeUUID   = "00000000-0000-0000-0000-000000000000"
	FakeApiKey = "fake-token"
)

type ApiTestCase struct {
	Name string

	ReqBody     []byte
	ReqCallback func(api any) (any, error)

	ResBody     []byte
	ResCallback func(res any)

	ExpectErr      error
	HttpStatusCode int
}

func RunApiTestCase(t *testing.T, tc ApiTestCase, api any) {
	t.Run(tc.Name, func(t *testing.T) {
		require.NotNil(t, tc.ReqCallback)
		res, err := tc.ReqCallback(api)

		if tc.ExpectErr != nil {
			require.ErrorContains(t, err, tc.ExpectErr.Error())
		} else {
			require.NoError(t, err)
		}

		if tc.ResCallback != nil {
			tc.ResCallback(res)
		}
	})
}

type FakeApiHTTPClient struct {
	t  *testing.T
	tc ApiTestCase
}

func CreateFakeHTTPClient(t *testing.T, tc ApiTestCase) *FakeApiHTTPClient {
	return &FakeApiHTTPClient{t, tc}
}

func (c *FakeApiHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.tc.ReqBody != nil {
		actualReqBody, err := io.ReadAll(req.Body)
		require.NoError(c.t, err)

		assert.Equal(c.t, string(c.tc.ReqBody), string(actualReqBody))
	} else {
		c.t.Logf("ReqBody not provided")
	}

	httpStatusCode := c.tc.HttpStatusCode

	return &http.Response{
		Status:     http.StatusText(httpStatusCode),
		StatusCode: httpStatusCode,
		Body:       io.NopCloser(bytes.NewReader(c.tc.ResBody)),
	}, nil
}
