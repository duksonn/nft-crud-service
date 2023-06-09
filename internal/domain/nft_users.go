package domain

type NftUsers struct {
	nft   Nft
	users []User
}

func NewNftUsers(nft Nft, users []User) *NftUsers {
	return &NftUsers{nft, users}
}

func (n *NftUsers) Nft() Nft {
	return n.nft
}

func (n *NftUsers) Users() []User {
	return n.users
}
