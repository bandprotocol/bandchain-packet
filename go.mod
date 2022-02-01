module github.com/bandprotocol/bandchain-packet

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0
	github.com/cosmos/ibc-go v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/stretchr/testify v1.7.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
