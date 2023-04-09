package json

import jsoniter "github.com/json-iterator/go"

var (
	json                = jsoniter.ConfigCompatibleWithStandardLibrary
	MarshalToString     = json.MarshalToString
	Marshal             = json.Marshal
	MarshalIndent       = json.MarshalIndent
	UnmarshalFromString = json.UnmarshalFromString
	Unmarshal           = json.Unmarshal
	Get                 = json.Get
	NewEncoder          = json.NewEncoder
	NewDecoder          = json.NewDecoder
	Valid               = json.Valid
	RegisterExtension   = json.RegisterExtension
	DecoderOf           = json.DecoderOf
	EncoderOf           = json.EncoderOf
)
