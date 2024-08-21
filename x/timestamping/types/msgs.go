package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTimestamp{}

func NewMsgCreateTimestamp(creator sdk.AccAddress, data string) *MsgCreateTimestamp {
	return &MsgCreateTimestamp{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgCreateTimestamp) Route() string {
	return RouterKey
}

func (msg *MsgCreateTimestamp) Type() string {
	return "CreateTimestamp"
}

func (msg *MsgCreateTimestamp) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

func (msg *MsgCreateTimestamp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTimestamp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Data == "" {
		return sdkerrors.Wrap(ErrInvalidData, "data cannot be empty")
	}
	return nil
}
