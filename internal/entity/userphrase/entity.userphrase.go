package userphrase

// UserPhrase define data model structure of a user phrase
type UserPhrase struct {
	UserID   int64
	PhraseID int64
	FilePath string
	MimeType MimeType
}
