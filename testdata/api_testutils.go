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
	T       *testing.T
	Adapter adapter.RequestAdapter
	Name    string
	Format  adapter.Format

	ReqBody     []byte
	ReqCallback func(m M) (any, error)

	ResBody     []byte
	ResCallback func(res any)

	ExpectErr      error
	HttpStatusCode int
}

type ApiTestParam[M api.Model] func(tc *apiTestCase[M])

type ApiTestParams[M api.Model] []ApiTestParam[M]

func (params ApiTestParams[M]) Run(t *testing.T, name string, m func(adp adapter.RequestAdapter) M) {
	tc := &apiTestCase[M]{
		T:              t,
		Name:           name,
		Format:         adapter.FormatJSON,
		ExpectErr:      nil,
		HttpStatusCode: http.StatusOK,
	}

	tc.Adapter = createFakeAdapter(t, tc)

	for _, fn := range params {
		fn(tc)
	}

	tc.Run(m(tc.Adapter))
}

func WithTestDataPath[M api.Model](v string) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		WithReqPath[M](fmt.Sprintf("%s/request/%s.%s", v, tc.Name, tc.Format))(tc)
		WithResPath[M](fmt.Sprintf("%s/response/%s.%s", v, tc.Name, tc.Format))(tc)
	}
}

func WithFormat[M api.Model](f adapter.Format) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.Format = f
	}
}

func WithReqPath[M api.Model](v string) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		require.NotEmpty(tc.T, v)

		var reqBody []byte

		_, err := os.Stat(v)
		require.NoError(tc.T, err)

		reqBody, err = os.ReadFile(v)
		require.NoError(tc.T, err)

		WithReqBody[M](reqBody)(tc)
	}
}

func WithReqBody[M api.Model](v []byte) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.ReqBody = v
	}
}

func WithReqCallback[M api.Model](v func(m M) (any, error)) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.ReqCallback = v
	}
}

func WithResPath[M api.Model](v string) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		require.NotEmpty(tc.T, v)

		var body []byte

		_, err := os.Stat(v)
		require.NoError(tc.T, err)

		body, err = os.ReadFile(v)
		require.NoError(tc.T, err)

		WithResBody[M](body)(tc)
	}
}

func WithResBody[M api.Model](v []byte) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.ResBody = v
	}
}

func WithResCallback[M api.Model](v func(res any)) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.ResCallback = v
	}
}

func WithExpectErr[M api.Model](v error) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.ExpectErr = v
	}
}

func WithHttpStatusCode[M api.Model](v int) ApiTestParam[M] {
	return func(tc *apiTestCase[M]) {
		tc.HttpStatusCode = v
	}
}

func (tc *apiTestCase[M]) Run(m M) {
	tc.T.Run(tc.Name, func(t *testing.T) {
		require.NotNil(t, tc.ReqCallback)
		res, err := tc.ReqCallback(m)

		if tc.ExpectErr != nil {
			require.ErrorContains(t, err, tc.ExpectErr.Error(), tc.Name)
		} else {
			require.NoError(t, err, tc.Name)
		}

		if tc.ResCallback != nil {
			tc.ResCallback(res)
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
	if c.tc.ReqBody != nil {
		c.checkRequestBody(req)
	}

	return &http.Response{
		Status:     http.StatusText(c.tc.HttpStatusCode),
		StatusCode: c.tc.HttpStatusCode,
		Body:       io.NopCloser(bytes.NewReader(c.tc.ResBody)),
	}, nil
}

func (c *fakeApiHTTPClient[M]) checkRequestBody(req *http.Request) {
	actualReqBody, err := io.ReadAll(req.Body)
	require.NoError(c.t, err)

	assert.Equal(c.t, string(c.tc.ReqBody), string(actualReqBody), c.tc.Name)
}

func createFakeAdapter[M api.Model](t *testing.T, tc *apiTestCase[M]) api.Client {
	return adapter.NewAdapter(adapter.NewConfig(
		fakeApiKey,
		adapter.WithHTTPClient(createFakeHTTPClient(t, tc)),
		adapter.WithFormat(tc.Format),
		adapter.WithDebug(true),
	))
}
