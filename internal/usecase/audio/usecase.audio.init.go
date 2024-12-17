package audio

type Usecase struct {
	repoFile       repoFile
	repoUser       repoUser
	repoPhrase     repoPhrase
	repoUserPhrase repoUserPhrase
}

func New(
	repoFile repoFile,
	repoUser repoUser,
	repoPhrase repoPhrase,
	repoUserPhrase repoUserPhrase,
) *Usecase {
	return &Usecase{
		repoFile:       repoFile,
		repoUser:       repoUser,
		repoPhrase:     repoPhrase,
		repoUserPhrase: repoUserPhrase,
	}
}
