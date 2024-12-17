package audio

type Handler struct {
	usecase usecaseInterface
}

func New(usecase usecaseInterface) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
