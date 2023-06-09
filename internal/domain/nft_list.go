package domain

type NftList struct {
	data []Nft
	next *int
	took *int
}

func NewNftList(data []Nft, next, took *int) *NftList {
	return &NftList{
		data: data,
		next: next,
		took: took,
	}
}

func (n *NftList) Data() []Nft {
	return n.data
}

func (n *NftList) Next() *int {
	return n.next
}

func (n *NftList) Took() *int {
	return n.took
}
