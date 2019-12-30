package base64

import (
	"encoding/base64"
)

var encodeType EncoderType

const (
	STANDARD EncoderType = iota
	URL
	STANDARD_NO_PADDING
	URL_NO_PADDING
)

var dtypeMap = map[EncoderType]*base64.Encoding{
	STANDARD:            base64.StdEncoding,
	URL:                 base64.URLEncoding,
	STANDARD_NO_PADDING: base64.RawStdEncoding,
	URL_NO_PADDING:      base64.RawURLEncoding,
}

type EncoderType int

func (dt EncoderType) GetEecoder() *base64.Encoding {
	return dtypeMap[dt]
}

func SetEncodeType(etype EncoderType) {
	encodeType = etype
}

func init() {
	SetEncodeType(STANDARD)
}
