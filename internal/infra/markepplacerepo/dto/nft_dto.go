package dto

import (
	"ssr/internal/domain"
	"strings"
)

type NftMsg struct {
	Id          string
	Image       string
	Description string
	Owner       string
	CoCreators  string
	CreatedAt   string
	CreatedBy   string
}

func (n *NftMsg) ToNftDomain() *domain.Nft {
	coCreators := transformCoCreatorsToStringList(&n.CoCreators)
	return domain.NewNft(
		n.Id,
		n.Image,
		n.Description,
		n.Owner,
		coCreators,
		n.CreatedAt,
		n.CreatedBy,
	)
}

func FromNftDomain(nft domain.Nft) NftMsg {
	coCreators := transformCoCreatorsToString(nft.CoCreators())
	return NftMsg{
		Id:          nft.Id(),
		Image:       nft.Image(),
		Description: nft.Description(),
		Owner:       nft.Owner(),
		CoCreators:  coCreators,
		CreatedAt:   nft.CreatedAt(),
		CreatedBy:   nft.CreatedBy(),
	}
}

func transformCoCreatorsToString(coCreators []string) string {
	if len(coCreators) > 0 {
		return strings.Join(coCreators, ",")
	}
	return ""
}

func transformCoCreatorsToStringList(coCreators *string) []string {
	if *coCreators != "" {
		return strings.Split(*coCreators, ",")
	}
	return nil
}

type NftListMsg struct {
	Data []NftMsg
	Next *int
	Took *int
}
