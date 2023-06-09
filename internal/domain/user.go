package domain

type User struct {
	id      string
	balance float64
}

func NewUser(id string, balance float64) *User {
	return &User{
		id:      id,
		balance: balance,
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Balance() float64 {
	return u.balance
}
