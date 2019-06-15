BUILD := `git rev-parse HEAD`

TAGS   := "netgo osusergo"
LFLAGS := -ldflags "-X=main.Build=$(BUILD)"
SRCS   := $(shell find . -name "main.go" | awk -F"/" '{ print $$3 }')

.PHONY: all clean install

all: $(SRCS) ui

$(SRCS):
	go build $(LFLAGS) -tags $(TAGS) -o $@.out ./cmd/$@ 

clean:
	-rm -f *.out

ui: template
	qtc -dir $<
	yarn run css