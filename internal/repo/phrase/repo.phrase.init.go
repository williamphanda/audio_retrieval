package userphrase

type Repo struct {
	dbInterface dbInterface
}

func New(dbInterface dbInterface) *Repo {
	return &Repo{
		dbInterface: dbInterface,
	}
}
