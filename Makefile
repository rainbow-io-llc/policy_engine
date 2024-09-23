G4_FILES := $(wildcard antlr/*.g4)

gen-build:
	mkdir -p gen/antlr
	@$(foreach file, $(G4_FILES), \
		mkdir -p gen/antlr/$(basename $(notdir $(file))); \
		antlr4 -Dlanguage=Go -o gen/antlr/$(basename $(notdir $(file))) $(file); \
	)

gen-clean:
	rm -fr gen/antlr

.PHONY: gen-build gen-clean
