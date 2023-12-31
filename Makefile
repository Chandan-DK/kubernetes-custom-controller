#########
# TOOLS #
#########

TOOLS_DIR                          := $(PWD)/.tools
DEEPCOPY_GEN                       := $(TOOLS_DIR)/deepcopy-gen
REGISTER_GEN                       := $(TOOLS_DIR)/register-gen
CLIENT_GEN                         := $(TOOLS_DIR)/client-gen
LISTER_GEN                         := $(TOOLS_DIR)/lister-gen
INFORMER_GEN                       := $(TOOLS_DIR)/informer-gen
CODE_GEN_VERSION                   := v0.28.0
CONTROLLER_GEN                     := $(TOOLS_DIR)/controller-gen
CONTROLLER_GEN_VERSION             := v0.12.0

$(DEEPCOPY_GEN):
	@echo Install deepcopy-gen... 
	@GOBIN=$(TOOLS_DIR) go install k8s.io/code-generator/cmd/deepcopy-gen@$(CODE_GEN_VERSION)

$(REGISTER_GEN):
	@echo Install register-gen... 
	@GOBIN=$(TOOLS_DIR) go install k8s.io/code-generator/cmd/register-gen@$(CODE_GEN_VERSION)

$(CLIENT_GEN):
	@echo Install client-gen... 
	@GOBIN=$(TOOLS_DIR) go install k8s.io/code-generator/cmd/client-gen@$(CODE_GEN_VERSION)

$(LISTER_GEN):
	@echo Install lister-gen...
	@GOBIN=$(TOOLS_DIR) go install k8s.io/code-generator/cmd/lister-gen@$(CODE_GEN_VERSION)

$(INFORMER_GEN):
	@echo Install informer-gen...
	@GOBIN=$(TOOLS_DIR) go install k8s.io/code-generator/cmd/informer-gen@$(CODE_GEN_VERSION)

$(CONTROLLER_GEN):
	@echo Install controller-gen...
	@GOBIN=$(TOOLS_DIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_GEN_VERSION)

install-tools: $(DEEPCOPY_GEN)
	@echo "All tools installed successfully."

clean-tools:
	@echo "Cleaning up tools..."
	@rm -rf $(TOOLS_DIR)
	@echo "Tools cleaned up successfully."

.PHONY: install-tools clean-tools


###########
# CODEGEN #
###########

PACKAGE                     := github.com/Chandan-DK/kubernetes-custom-controller
GOPATH_SHIM                 := ${PWD}/.gopath
PACKAGE_SHIM                := $(GOPATH_SHIM)/src/$(PACKAGE)
OUT_PACKAGE                 := $(PACKAGE)/pkg/generated/client
INPUT_DIRS                  := $(PACKAGE)/api/mygroup/v1alpha1
CLIENTSET_PACKAGE           := $(OUT_PACKAGE)/clientset
LISTERS_PACKAGE             := $(OUT_PACKAGE)/listers
INFORMERS_PACKAGE           := $(OUT_PACKAGE)/informers
CRDS_PATH                   := ${PWD}/config/crds/generated

$(GOPATH_SHIM):
	@echo Create gopath shim... 
	@mkdir -p $(GOPATH_SHIM)

.INTERMEDIATE: $(PACKAGE_SHIM)
$(PACKAGE_SHIM): $(GOPATH_SHIM)
	@echo Create package shim... 
	@mkdir -p $(GOPATH_SHIM)/src/github.com/Chandan-DK && ln -s -f ${PWD} $(PACKAGE_SHIM)

.PHONY: codegen-deepcopy
codegen-deepcopy: $(PACKAGE_SHIM) $(DEEPCOPY_GEN) ## Generate deep copy functions
	@echo Generate deep copy functions... 
	@GOPATH=$(GOPATH_SHIM) $(DEEPCOPY_GEN) \
		--go-header-file=./scripts/boilerplate.go.txt \
		--input-dirs=$(INPUT_DIRS) \
		--output-file-base=zz_generated.deepcopy

.PHONY: codegen-register
codegen-register: $(PACKAGE_SHIM) $(REGISTER_GEN) ## Generate types registrations
	@echo Generate registration... 
	@GOPATH=$(GOPATH_SHIM) $(REGISTER_GEN) \
		--go-header-file=./scripts/boilerplate.go.txt \
		--input-dirs=$(INPUT_DIRS)

.PHONY: codegen-client-clientset
codegen-client-clientset: $(PACKAGE_SHIM) $(CLIENT_GEN) ## Generate clientset
	@echo Generate clientset... 
	@GOPATH=$(GOPATH_SHIM) $(CLIENT_GEN) \
		--go-header-file ./scripts/boilerplate.go.txt \
		--clientset-name versioned \
		--output-package $(CLIENTSET_PACKAGE) \
		--input-base "" \
		--input $(INPUT_DIRS)

.PHONY: codegen-client-listers
codegen-client-listers: $(PACKAGE_SHIM) $(LISTER_GEN) ## Generate listers
	@echo Generate listers... 
	@GOPATH=$(GOPATH_SHIM) $(LISTER_GEN) \
		--go-header-file ./scripts/boilerplate.go.txt \
		--output-package $(LISTERS_PACKAGE) \
		--input-dirs $(INPUT_DIRS)

.PHONY: codegen-client-informers
codegen-client-informers: $(PACKAGE_SHIM) $(INFORMER_GEN) ## Generate informers
	@echo Generate informers... 
	@GOPATH=$(GOPATH_SHIM) $(INFORMER_GEN) \
		--go-header-file ./scripts/boilerplate.go.txt \
		--output-package $(INFORMERS_PACKAGE) \
		--input-dirs $(INPUT_DIRS) \
		--versioned-clientset-package $(CLIENTSET_PACKAGE)/versioned \
		--listers-package $(LISTERS_PACKAGE)

.PHONY: codegen-crds-mygroup
codegen-crds-mygroup: $(CONTROLLER_GEN) ## Generate CRDs
	@echo Generate mygroup crds... 
	@$(CONTROLLER_GEN) crd paths=./api/mygroup/... crd:crdVersions=v1 output:dir=$(CRDS_PATH)

codegen-all: codegen-deepcopy codegen-register codegen-client-clientset codegen-client-listers codegen-client-informers codegen-crds-mygroup