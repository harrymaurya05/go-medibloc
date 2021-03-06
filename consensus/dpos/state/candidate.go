package dposstate

import (
	"github.com/gogo/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	dpospb "github.com/medibloc/go-medibloc/consensus/dpos/pb"
	"github.com/medibloc/go-medibloc/util"
)

// Candidate is struct for save candidate state
type Candidate struct {
	ID         []byte // candidate id = txHash
	Addr       common.Address
	Collateral *util.Uint128 // candidate collateral
	VotePower  *util.Uint128 // sum of voters' staking
	URL        string
	Timestamp  int64
}

func (c *Candidate) fromProto(pbCandidate *dpospb.Candidate) error {
	var err error
	c.ID = pbCandidate.Id
	if err := c.Addr.FromBytes(pbCandidate.Addr); err != nil {
		return err
	}
	c.Collateral, err = util.NewUint128FromFixedSizeByteSlice(pbCandidate.Collateral)
	if err != nil {
		return err
	}
	c.VotePower, err = util.NewUint128FromFixedSizeByteSlice(pbCandidate.VotePower)
	if err != nil {
		return err
	}
	c.URL = pbCandidate.Url
	c.Timestamp = pbCandidate.Timestamp
	return nil
}

func (c *Candidate) toProto() (*dpospb.Candidate, error) {
	collateral, err := c.Collateral.ToFixedSizeByteSlice()
	if err != nil {
		return nil, err
	}
	votePower, err := c.VotePower.ToFixedSizeByteSlice()
	if err != nil {
		return nil, err
	}
	return &dpospb.Candidate{
		Id:         c.ID,
		Addr:       c.Addr.Bytes(),
		Collateral: collateral,
		VotePower:  votePower,
		Url:        c.URL,
		Timestamp:  c.Timestamp,
	}, nil
}

// FromBytes set Candidate struct from bytes
func (c *Candidate) FromBytes(bytes []byte) error {
	var err error
	pbCandidate := new(dpospb.Candidate)
	if err := proto.Unmarshal(bytes, pbCandidate); err != nil {
		return err
	}
	c.ID = pbCandidate.Id
	if err := c.Addr.FromBytes(pbCandidate.Addr); err != nil {
		return err
	}
	c.Collateral, err = util.NewUint128FromFixedSizeByteSlice(pbCandidate.Collateral)
	if err != nil {
		return err
	}
	c.VotePower, err = util.NewUint128FromFixedSizeByteSlice(pbCandidate.VotePower)
	if err != nil {
		return err
	}
	c.URL = pbCandidate.Url
	c.Timestamp = pbCandidate.Timestamp
	return nil
}

// ToBytes marshal Candidate struct to bytes
func (c *Candidate) ToBytes() ([]byte, error) {
	collateral, err := c.Collateral.ToFixedSizeByteSlice()
	if err != nil {
		return nil, err
	}
	votePower, err := c.VotePower.ToFixedSizeByteSlice()
	if err != nil {
		return nil, err
	}

	pbCandidate := &dpospb.Candidate{
		Id:         c.ID,
		TxHash:     c.ID,
		Addr:       c.Addr.Bytes(),
		Collateral: collateral,
		VotePower:  votePower,
		Url:        c.URL,
		Timestamp:  c.Timestamp,
	}
	return proto.Marshal(pbCandidate)
}
