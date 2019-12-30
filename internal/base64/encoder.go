package base64

func Encode(s string) string {
	encoder := encodeType.GetEecoder()
	ds := encoder.EncodeToString([]byte(s))
	return ds
}
