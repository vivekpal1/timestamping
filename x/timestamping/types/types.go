package types

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Timestamp struct {
	Creator   sdk.AccAddress `json:"creator"`
	Data      string         `json:"data"`
	Timestamp time.Time      `json:"timestamp"`
	Hash      string         `json:"hash"`
	BTCHeight int64          `json:"btc_height"`
}

func NewTimestamp(creator sdk.AccAddress, data string, btcHeight int64) Timestamp {
	timestamp := time.Now().UTC()
	return Timestamp{
		Creator:   creator,
		Data:      data,
		Timestamp: timestamp,
		Hash:      calculateHash(data, timestamp),
		BTCHeight: btcHeight,
	}
}

func (t Timestamp) Validate() error {
	if t.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator address cannot be empty")
	}
	if t.Data == "" {
		return sdkerrors.Wrap(ErrInvalidData, "data cannot be empty")
	}
	if t.Timestamp.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "timestamp cannot be zero")
	}
	if t.Hash == "" {
		return sdkerrors.Wrap(ErrInvalidHash, "hash cannot be empty")
	}
	if t.BTCHeight <= 0 {
		return sdkerrors.Wrap(ErrInvalidBTCHeight, "BTC height must be positive")
	}
	return nil
}

func calculateHash(data string, timestamp time.Time) string {
	input := data + timestamp.String()
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
