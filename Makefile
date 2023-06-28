DOCKER=docker
protoVer=0.12.1
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto: proto-format proto-gen

proto-gen:
	@$(protoImage) sh ./scripts/protocgen.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json ./proto

proto-check-breaking:
	@$(protoImage) buf breaking --against-input '.git#branch=main'

.PHONY: proto proto-gen proto-format proto-lint proto-check-breaking
