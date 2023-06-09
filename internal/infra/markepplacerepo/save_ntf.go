package markepplacerepo

import (
	"context"
	"fmt"
	"nft-crud-service/internal/domain"
	"nft-crud-service/internal/infra/markepplacerepo/dto"
)

func (m *MysqlRepository) SaveNft(ctx context.Context, nft domain.Nft) (*domain.Nft, error) {
	nftMsg := dto.FromNftDomain(nft)
	query := fmt.Sprintf(
		"insert into nft(id, image, description, owner, co_creators, created_at, created_by) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v') "+
			"on duplicate key update id=values(id),image=values(image),description=values(description),owner=values(owner),co_creators=values(co_creators),created_at=values(created_at),created_by=values(created_by)",
		nftMsg.Id,
		nftMsg.Image,
		nftMsg.Description,
		nftMsg.Owner,
		nftMsg.CoCreators,
		nftMsg.CreatedAt,
		nftMsg.CreatedBy,
	)
	_, err := m.client.Exec(query)
	if err != nil {
		return nil, err
	}

	var nftResponse dto.NftDTO
	err = m.findOne(ctx, nft.Id(), NftTable).Scan(
		&nftResponse.Id,
		&nftResponse.Image,
		&nftResponse.Description,
		&nftResponse.Owner,
		&nftResponse.CoCreators,
		&nftResponse.CreatedAt,
		&nftResponse.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return nftResponse.ToNftDomain(), nil
}
