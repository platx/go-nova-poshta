package adapter

import (
	"errors"
	"fmt"
	"net/http"
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
	Success      bool     `json:"success" xml:"success"`
	Data         any      `json:"data" xml:"data"`
	Errors       []string `json:"errors" xml:"errors"`
	Warnings     any      `json:"warnings" xml:"warnings"` // FIXME: empty array or map[string]any or []map[string]string
	Info         any      `json:"info" xml:"info"`         // FIXME: empty array or map[string]any
	MessageCodes []string `json:"messageCodes" xml:"messageCodes"`
	ErrorCodes   []string `json:"errorCodes" xml:"errorCodes"`
	WarningCodes []string `json:"warningCodes" xml:"warningCodes"`
	InfoCodes    []string `json:"infoCodes" xml:"infoCodes"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestAdapter interface {
	Req(modelName string, calledMethod string, props any, res any) error
}

type httpAdapter struct {
	http       HTTPClient
	serializer serializer
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
		baseUrl:    baseUrl,
		serializer: newSerializer(cfg.Format),
		apiKey:     cfg.ApiKey,
	}
}

func (c *httpAdapter) Req(modelName string, calledMethod string, reqProps any, resData any) error {
	if reqProps == nil {
		reqProps = struct{}{}
	}

	req := request{
		ApiKey:           c.apiKey,
		ModelName:        modelName,
		CalledMethod:     calledMethod,
		MethodProperties: reqProps,
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

	res := &response{Data: resData}

	if err = c.serializer.decode(httpRes.Body, res); err != nil {
		return fmt.Errorf("decode response failed => %w", err)
	}

	if !res.Success {
		var errs []error

		for _, resErr := range res.Errors {
			errs = append(errs, errors.New(resErr))
		}

		return fmt.Errorf("request failed => %w", errors.Join(errs...))
	}

	return nil
}
