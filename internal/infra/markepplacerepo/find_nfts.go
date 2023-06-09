package markepplacerepo

import (
	"context"
	"database/sql"
	"ssr/internal/domain"
	"ssr/internal/infra/markepplacerepo/dto"
)

func (m *MysqlRepository) FindNfts(ctx context.Context, next, took *int) (*domain.NftList, error) {
	var nextResp int
	var res *sql.Rows
	var err error
	if next != nil && *next != 0 {
		res, err = m.findAllWithPagination(ctx, next, took, NftTable)
		if err != nil {
			return nil, err
		}

		countRes, err := m.getCount(ctx, NftTable)
		if err != nil {
			return nil, err
		}
		var countInt int
		for countRes.Next() {
			err = countRes.Scan(&countInt)
			if err != nil {
				return nil, err
			}
		}
		nextResp = countInt - *next - *took

	} else {
		res, err = m.findAll(ctx, NftTable)
		if err != nil {
			return nil, err
		}
	}

	var nfts []domain.Nft
	for res.Next() {
		var nftResponse dto.NftMsg
		err := res.Scan(
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
		nfts = append(nfts, *nftResponse.ToNftDomain())
	}

	nftList := domain.NewNftList(nfts, &nextResp, next)
	return nftList, nil
}
