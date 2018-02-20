

OUTDIR = $(CURDIR)/android/app/libs

	
all: cmd bind lint

.PHONY: cmd
cmd:
	cd $(CURDIR)/cmd/gomochat && go build

.PHONY: bind
bind:
	gomobile bind -o $(OUTDIR)/gomochat.aar

.PHONY: lint
lint:
	gometalinter $(CURDIR)/...

clean:
	rm -rf $(OUTDIR)/*
