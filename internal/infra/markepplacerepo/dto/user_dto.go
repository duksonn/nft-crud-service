package dto

import "ssr/internal/domain"

type UserMsg struct {
	Id      string
	Balance float64
}

func (u *UserMsg) ToUserDomain() *domain.User {
	return domain.NewUser(u.Id, u.Balance)
}

func FromUserDomain(user domain.User) UserMsg {
	return UserMsg{
		Id:      user.Id(),
		Balance: user.Balance(),
	}
}
