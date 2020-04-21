package interfaces

type userRepository struct {
	SQLHandler
}

func NewUserRepository(sqlhandler SQLHandler) *userRepository {
	return &userRepository{sqlhandler}
}
