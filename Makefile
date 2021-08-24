DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh

proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace tendermintdev/docker-build-proto \
	find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) check breaking --against-input $(HTTPS_GIT)#branch=master

.PHONY: proto-all proto-gen proto-gen-any proto-swagger-gen proto-format proto-lint proto-check-breaking 
