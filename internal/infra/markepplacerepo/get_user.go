package markepplacerepo

import (
	"context"
	"errors"
	"fmt"
	"nft-crud-service/internal/domain"
	"nft-crud-service/internal/infra/markepplacerepo/dto"
)

func (m *MysqlRepository) GetUserById(ctx context.Context, userId string) (*domain.User, error) {
	var userMsg dto.UserMsg
	err := m.findOne(ctx, userId, UserTable).Scan(&userMsg.Id, &userMsg.Balance)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New(fmt.Sprintf("user '%v' doesnt exist in database", userId))
		}
		return nil, err
	}
	return userMsg.ToUserDomain(), nil
}
