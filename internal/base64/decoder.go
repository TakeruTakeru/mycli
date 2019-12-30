package base64

func Decode(s string) (string, error) {
	encoder := encodeType.GetEecoder()
	ds, err := encoder.DecodeString(s)
	return string(ds), err
}
