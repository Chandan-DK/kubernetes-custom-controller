#########
# TOOLS #
#########

TOOLS_DIR                          := $(PWD)/.tools
DEEPCOPY_GEN                       := $(TOOLS_DIR)/deepcopy-gen
CODE_GEN_VERSION                   := v0.28.0

$(DEEPCOPY_GEN):
	@echo Install deepcopy-gen... 
	@GOBIN=$(TOOLS_DIR) go install k8s.io/code-generator/cmd/deepcopy-gen@$(CODE_GEN_VERSION)

install-tools: $(DEEPCOPY_GEN)
	@echo "All tools installed successfully."

clean-tools:
	@echo "Cleaning up tools..."
	@rm -rf $(TOOLS_DIR)
	@echo "Tools cleaned up successfully."

.PHONY: install-tools clean-tools
