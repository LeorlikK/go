package generics

type IUser interface {
	GetId(id int) int
}

type User struct {
	ID int
}

func (u *User) GetId(id int) int {
	u.ID = 35
	return u.ID
}

func SumGeneric[T int | string](a, b T) T {
	return a + b
}

func GetIdGeneric[T IUser](user T) T {
	return user
}

func GetIdGenericTwo[T IUser](user *T) *T {
	return user
}
