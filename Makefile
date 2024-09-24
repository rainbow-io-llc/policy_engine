G4_FILES := $(wildcard antlr/*.g4)

gen-build:
	@$(foreach file, $(G4_FILES), \
		cd antlr && antlr4 -Dlanguage=Go -o ../gen/antlr/$(basename $(notdir $(file))) -DgoTargetPackage=$(basename $(notdir $(file))) $(notdir $(file)) && cd ..; \
	)

gen-clean:
	rm -fr gen

.PHONY: gen-build gen-clean
