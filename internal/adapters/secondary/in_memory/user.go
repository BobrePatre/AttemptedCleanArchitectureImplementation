package in_memory

type UserCounterRepository struct {
	counter int
}

func NewUserCounterRepository() *UserCounterRepository {
	return &UserCounterRepository{
		counter: 0,
	}
}

func (usr *UserCounterRepository) Increment() int {
	usr.counter++
	return usr.counter
}
