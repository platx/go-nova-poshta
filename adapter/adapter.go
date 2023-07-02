package adapter

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/platx/go-nova-poshta/custom/types"
)

const baseUrlTemplate = "https://api.novaposhta.ua/v2.0/%s/"

type Format string

const (
	FormatJSON Format = "json"
	FormatXML  Format = "xml"
)

type request struct {
	ApiKey           string `json:"apiKey" xml:"apiKey"`
	ModelName        string `json:"modelName" xml:"modelName"`
	CalledMethod     string `json:"calledMethod" xml:"calledMethod"`
	MethodProperties any    `json:"methodProperties" xml:"methodProperties"`
}

type response struct {
	Success      bool                   `json:"success" xml:"success"`
	Data         any                    `json:"data" xml:"data"`
	Errors       types.Messages[string] `json:"errors" xml:"errors"`
	Warnings     types.Messages[string] `json:"warnings" xml:"warnings"`
	Info         types.Messages[string] `json:"info" xml:"info"`
	MessageCodes types.Messages[string] `json:"messageCodes" xml:"messageCodes"`
	ErrorCodes   types.Messages[string] `json:"errorCodes" xml:"errorCodes"`
	WarningCodes types.Messages[string] `json:"warningCodes" xml:"warningCodes"`
	InfoCodes    types.Messages[string] `json:"infoCodes" xml:"infoCodes"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestAdapter interface {
	Call(modelName string, calledMethod string, props any, resData any) error
}

type httpAdapter struct {
	http       httpClient
	serializer serializer
	logger     logger
	baseUrl    string
	apiKey     string
}

func NewAdapter(cfg Config) RequestAdapter {
	var baseUrl string
	switch cfg.Format {
	case FormatJSON, FormatXML:
		baseUrl = fmt.Sprintf(baseUrlTemplate, cfg.Format)
	default:
		panic(fmt.Sprintf("unknown format '%s'", cfg.Format))
	}

	return &httpAdapter{
		http:       cfg.HTTPClient,
		serializer: newSerializer(cfg.Format, cfg.Debug),
		logger:     cfg.Logger,
		baseUrl:    baseUrl,
		apiKey:     cfg.ApiKey,
	}
}

func (c *httpAdapter) Call(model string, method string, props any, resData any) error {
	c.logger.Debugf("call %s.%s with props %+v", model, method, props)

	req := request{
		ModelName:        model,
		CalledMethod:     method,
		MethodProperties: props,
	}

	res := &response{Data: resData}

	if err := c.do(req, res); err != nil {
		return err
	}

	if !res.Success {
		var errs []error

		for i, errCode := range res.ErrorCodes {
			trans, ok := ErrorTranslations[errCode]["en"]
			if !ok {
				trans = res.Errors[i]
			}

			errs = append(errs, errors.New(trans))
		}

		return fmt.Errorf("request failed => %w", errors.Join(errs...))
	}

	c.logger.Debugf("got success response %+v", resData)

	return nil
}

func (c *httpAdapter) do(req request, res *response) error {
	if req.MethodProperties == nil {
		req.MethodProperties = struct{}{}
	}

	if req.ApiKey == "" {
		req.ApiKey = c.apiKey
	}

	encReq, err := c.serializer.encode(req)
	if err != nil {
		return fmt.Errorf("encode request failed => %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.baseUrl, encReq)
	if err != nil {
		return fmt.Errorf("create http request failed => %w", err)
	}

	httpRes, err := c.http.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send http request failed => %w", err)
	}

	if err = c.serializer.decode(httpRes.Body, res); err != nil {
		return fmt.Errorf("decode response failed => %w", err)
	}

	return nil
}
