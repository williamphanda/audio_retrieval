package userphrase

// MimeType is an enum for user phrase entry
type MimeType int8

const (
	MimeTypeMPEG MimeType = iota + 1
	MimeTypeWAV
	MimeTypeOGG
)

// TODO: use library / better way to check mime type
var MapMimeTypeToEnum = map[string]MimeType{
	"audio/mpeg": MimeTypeMPEG,
	"audio/wav":  MimeTypeWAV,
	"audio/wave": MimeTypeWAV,
	"audio/ogg":  MimeTypeOGG,
	"mpeg":       MimeTypeMPEG,
	"wav":        MimeTypeWAV,
	"wave":       MimeTypeWAV,
	"ogg":        MimeTypeOGG,
}
