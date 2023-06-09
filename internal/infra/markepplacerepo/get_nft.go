package markepplacerepo

import (
	"context"
	"errors"
	"fmt"
	"ssr/internal/domain"
	"ssr/internal/infra/markepplacerepo/dto"
)

func (m *MysqlRepository) GetNftById(ctx context.Context, nftId string) (*domain.Nft, error) {
	var nftMsg dto.NftMsg
	err := m.findOne(ctx, nftId, NftTable).Scan(
		&nftMsg.Id,
		&nftMsg.Image,
		&nftMsg.Description,
		&nftMsg.Owner,
		&nftMsg.CoCreators,
		&nftMsg.CreatedAt,
		&nftMsg.CreatedBy,
	)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New(fmt.Sprintf("nft '%v' doesnt exist in database", nftId))
		}
		return nil, err
	}

	return nftMsg.ToNftDomain(), nil
}
