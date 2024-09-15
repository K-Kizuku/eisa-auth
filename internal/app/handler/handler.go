package handler

type Root struct {
	UserHandler IUserHandler
}

func New(UserHandler IUserHandler) *Root {
	return &Root{
		UserHandler: UserHandler,
	}
}
