# Set the default goal
.DEFAULT_GOAL=all

# A phony target is one that is not really the name of a file; rather it is just a name for a recipe to be executed when you make an explicit request. There are two reasons to use a phony target: to avoid a conflict with a file of the same name, and to improve performance.
.PHONY: all protocol msgid fmt clean mostlyclean distclean realclean clobber install print tar shar dist TAGS check test

EXEC=bin/3rd

PROTO_FILES=$(wildcard pb/*/*.proto)
PB_GO_FILES=$(patsubst %.proto,%.pb.go,$(PROTO_FILES))
GO_FILES=$(wildcard **/*.go)

# Make all the top-level targets the makefile knows about.
all: $(EXEC)

bin/%: protocol $(GO_FILES)
	@echo building $@
	@go build -o $@ github.com/enjoypi/$(notdir $@)

# Generate all protocol
protocol: $(PB_GO_FILES)

pb/cspb/msg.go: $(PROTO_FILES)
	@echo generating $@
#	@python3 scripts/gen_msg.go.py --in pb/cspb/req.proto --in pb/cspb/ack.proto --in pb/cspb/login.proto --in pb/cspb/ntf.proto  --out pb/cspb/msg.go
	@goimports -w $@

# *.proto to *.pb.go
%.pb.go: %.proto
	@echo generating $@
	@protoc --proto_path=./ --gofast_out=plugins=grpc:./  $(patsubst %.pb.go,%.proto,$@)

# Message ID
msgid: pb/cspb/msgid.def

pb/cspb/msgid.def: bin/genmsgid
	@echo generating $@
	@bin/genmsgid

bin/genmsgid: protocol
	@echo building $@
	@go build -o bin/genmsgid  git.tap4fun.com/fw/gate/utility/genmsgid

# Format all sources
fmt: GO_FILES=$(wildcard *.go) $(wildcard */*.go) $(wildcard */*/*.go)
fmt:
	goimports -w $(GO_FILES)

# Delete all files that are normally created by running make.
clean:
	git clean -dxf

# Like ‘clean’, but may refrain from deleting a few files that people normally don’t want to recompile. For example, the ‘mostlyclean’ target for GCC does not delete libgcc.a, because recompiling it is rarely necessary and takes a lot of time.
mostlyclean:

# Any of these targets might be defined to delete more files than ‘clean’ does. For example, this would delete configuration files or links that you would normally create as preparation for compilation, even if the makefile itself cannot create these files.
distclean:
realclean:
clobber:

# Copy the executable file into a directory that users typically search for commands; copy any auxiliary files that the executable uses into the directories where it will look for them.
install:

# Print listings of the source files that have changed.
print:
	@echo $(PROTO_FILES)
	@echo $(PB_GO_FILES)
	@echo $(GO_FILES)

# Create a tar file of the source files.
tar:

# Create a shell archive (shar file) of the source files.
shar:

# Create a distribution file of the source files. This might be a tar file, or a shar file, or a compressed version of one of the above, or even more than one of the above.
dist:

# Update a tags table for this program.
TAGS:

# Perform self tests on the program this makefile builds.
check:
test: