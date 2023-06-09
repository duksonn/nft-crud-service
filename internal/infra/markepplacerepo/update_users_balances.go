package markepplacerepo

import (
	"context"
	"fmt"
	"nft-crud-service/internal/domain"
	"nft-crud-service/internal/infra/markepplacerepo/dto"
)

func (m *MysqlRepository) UpdateUsersBalances(ctx context.Context, balances map[string]float64) ([]domain.User, error) {
	var values string
	ids := ""
	totalBalances := len(balances)
	for user, balance := range balances {
		values = fmt.Sprintf("%v('%v', '%v')", values, user, balance)
		ids = fmt.Sprintf("%v'%v'", ids, user)
		if totalBalances-1 > 0 {
			values = fmt.Sprintf("%v,", values)
			ids = fmt.Sprintf("%v,", ids)
		}
		totalBalances = totalBalances - 1
	}
	query := fmt.Sprintf("insert into %v(id, balance) VALUES %v on duplicate key update balance=values(balance)", UserTable, values)
	_, err := m.client.Exec(query)
	if err != nil {
		return nil, err
	}

	res, err := m.findMany(ctx, ids, UserTable)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for res.Next() {
		var userResponse dto.UserDTO
		err := res.Scan(&userResponse.Id, &userResponse.Balance)
		if err != nil {
			return nil, err
		}
		users = append(users, *userResponse.ToUserDomain())
	}

	return users, nil
}
