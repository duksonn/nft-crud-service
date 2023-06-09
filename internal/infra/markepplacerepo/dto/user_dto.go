package dto

import "nft-crud-service/internal/domain"

type UserDTO struct {
	Id      string
	Balance float64
}

func (u *UserDTO) ToUserDomain() *domain.User {
	return domain.NewUser(u.Id, u.Balance)
}

func FromUserDomain(user domain.User) UserDTO {
	return UserDTO{
		Id:      user.Id(),
		Balance: user.Balance(),
	}
}
