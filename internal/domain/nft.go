package domain

type Nft struct {
	id          string
	image       string
	description string
	owner       string
	coCreators  []string
	createdAt   string
	createdBy   string
}

func NewNft(id, image, description, owner string, coCreators []string, createdAt, createdBy string) *Nft {
	return &Nft{
		id:          id,
		image:       image,
		description: description,
		owner:       owner,
		coCreators:  coCreators,
		createdAt:   createdAt,
		createdBy:   createdBy,
	}
}

func (n *Nft) Id() string {
	return n.id
}

func (n *Nft) Image() string {
	return n.image
}

func (n *Nft) Description() string {
	return n.description
}

func (n *Nft) Owner() string {
	return n.owner
}

func (n *Nft) CoCreators() []string {
	return n.coCreators
}

func (n *Nft) CreatedAt() string {
	return n.createdAt
}

func (n *Nft) CreatedBy() string {
	return n.createdBy
}
