package adapter

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
)

type encoder interface {
	encode(any) (io.Reader, error)
}

type decoder interface {
	decode(io.Reader, any) error
}

type serializer interface {
	encoder
	decoder
}

func newSerializer(f Format) serializer {
	switch f {
	case FormatJSON:
		return &serializerJSON{}
	case FormatXML:
		return &serializerXML{}
	}

	panic("unknown format '%s'")
}

type serializerJSON struct{}

func (enc *serializerJSON) encode(v any) (io.Reader, error) {
	jsonData, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonData), nil
}

func (enc *serializerJSON) decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

type serializerXML struct{}

func (enc *serializerXML) encode(v any) (io.Reader, error) {
	jsonData, err := xml.Marshal(v)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonData), nil
}

func (enc *serializerXML) decode(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}
