package testdata

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/api"
)

const (
	fakeApiKey = "fake-token"
)

type apiTestCase[M api.Model] struct {
	t       *testing.T
	adapter adapter.RequestAdapter
	name    string
	format  adapter.Format

	reqBody     []byte
	reqCallback func(m M) (any, error)

	resBody     []byte
	resCallback func(res any)

	expectErr      error
	httpStatusCode int
}

type ApiTestParam[M api.Model] func(tc *apiTestCase[M])

type ApiTestParams[M api.Model] []ApiTestParam[M]

func (params ApiTestParams[M]) Run(t *testing.T, name string, m func(adp adapter.RequestAdapter) M) {
	tc := &apiTestCase[M]{
		t:              t,
		name:           name,
		format:         adapter.FormatJSON,
		expectErr:      nil,
		httpStatusCode: http.StatusOK,
	}

	tc.adapter = createFakeAdapter(t, tc)

	for _, fn := range params {
		fn(tc)
	}

	tc.run(m(tc.adapter))
}

func WithTestDataPath[M api.Model](v string) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		WithReqPath[M](fmt.Sprintf("%s/request/%s.%s", v, tc.name, tc.format))(tc)
		WithResPath[M](fmt.Sprintf("%s/response/%s.%s", v, tc.name, tc.format))(tc)
	}
}

func WithFormat[M api.Model](f adapter.Format) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.format = f
	}
}

func WithReqPath[M api.Model](v string) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		require.NotEmpty(tc.t, v)

		var reqBody []byte

		_, err := os.Stat(v)
		require.NoError(tc.t, err)

		reqBody, err = os.ReadFile(v)
		require.NoError(tc.t, err)

		WithReqBody[M](reqBody)(tc)
	}
}

func WithReqBody[M api.Model](v []byte) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.reqBody = v
	}
}

func WithReqCallback[M api.Model](v func(m M) (any, error)) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.reqCallback = v
	}
}

func WithResPath[M api.Model](v string) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		require.NotEmpty(tc.t, v)

		var body []byte

		_, err := os.Stat(v)
		require.NoError(tc.t, err)

		body, err = os.ReadFile(v)
		require.NoError(tc.t, err)

		WithResBody[M](body)(tc)
	}
}

func WithResBody[M api.Model](v []byte) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.resBody = v
	}
}

func WithResCallback[M api.Model](v func(res any)) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.resCallback = v
	}
}

func WithExpectErr[M api.Model](v error) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.expectErr = v
	}
}

func WithHttpStatusCode[M api.Model](v int) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.httpStatusCode = v
	}
}

func (tc *apiTestCase[M]) run(m M) {
	tc.t.Run(tc.name, func(t *testing.T) {
		require.NotNil(t, tc.reqCallback)
		res, err := tc.reqCallback(m)

		if tc.expectErr != nil {
			require.ErrorContains(t, err, tc.expectErr.Error(), tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}

		if tc.resCallback != nil {
			tc.resCallback(res)
		}
	})
}

type fakeApiHTTPClient[M api.Model] struct {
	t  *testing.T
	tc *apiTestCase[M]
}

func createFakeHTTPClient[M api.Model](t *testing.T, tc *apiTestCase[M]) *fakeApiHTTPClient[M] {
	return &fakeApiHTTPClient[M]{t, tc}
}

func (c *fakeApiHTTPClient[M]) Do(req *http.Request) (*http.Response, error) {
	if c.tc.reqBody != nil {
		c.checkRequestBody(req)
	}

	return &http.Response{
		Status:     http.StatusText(c.tc.httpStatusCode),
		StatusCode: c.tc.httpStatusCode,
		Body:       io.NopCloser(bytes.NewReader(c.tc.resBody)),
	}, nil
}

func (c *fakeApiHTTPClient[M]) checkRequestBody(req *http.Request) {
	actualReqBody, err := io.ReadAll(req.Body)
	require.NoError(c.t, err)

	assert.Equal(c.t, string(c.tc.reqBody), string(actualReqBody), c.tc.name)
}

func createFakeAdapter[M api.Model](t *testing.T, tc *apiTestCase[M]) api.Client {
	return adapter.NewAdapter(adapter.NewConfig(
		fakeApiKey,
		adapter.WithHTTPClient(createFakeHTTPClient(t, tc)),
		adapter.WithFormat(tc.format),
		adapter.WithDebug(true),
	))
}
