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

func newSerializer(f Format, verbose bool) serializer {
	switch f {
	case FormatJSON:
		return &serializerJSON{verbose: verbose}
	case FormatXML:
		return &serializerXML{verbose: verbose}
	}

	panic("unknown format '%s'")
}

type serializerJSON struct{ verbose bool }

func (enc *serializerJSON) encode(v any) (io.Reader, error) {
	var (
		encData []byte
		err     error
	)

	if enc.verbose {
		encData, err = json.MarshalIndent(v, "", "  ")
	} else {
		encData, err = json.Marshal(v)
	}

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(encData), nil
}

func (enc *serializerJSON) decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

type serializerXML struct{ verbose bool }

func (enc *serializerXML) encode(v any) (io.Reader, error) {
	var (
		encData []byte
		err     error
	)

	if enc.verbose {
		encData, err = xml.MarshalIndent(v, "", "  ")
	} else {
		encData, err = xml.Marshal(v)
	}

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(encData), nil
}

func (enc *serializerXML) decode(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}
