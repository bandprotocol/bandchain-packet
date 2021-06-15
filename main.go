package main

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"

	"github.com/bandprotocol/export-packet/packet"
)

func main() {
	calldata, _ := hex.DecodeString("030000004254436400000000000000")
	req := packet.OracleRequestPacketData{
		ClientID:       "test",
		OracleScriptID: 1,
		Calldata:       calldata,
		AskCount:       1,
		MinCount:       1,
		FeeLimit:       sdk.NewCoins(sdk.NewCoin("uband", sdk.NewInt(10000))),
		RequestKey:     "TEST_KEY",
		PrepareGas:     100,
		ExecuteGas:     100,
	}
	// This can be data of packet
	fmt.Println(string(req.GetBytes()))
	packet := channeltypes.NewPacket(req.GetBytes(), 0, "src-port", "src-channel", "dst-port", "dst-channel", clienttypes.NewHeight(0, 0), 10)
	_ = packet
}
