

OUTDIR = $(CURDIR)/android/app/libs

	
all: cmd bind

.PHONY: cmd
cmd:
	cd $(CURDIR)/cmd/gomochat && go build

.PHONY: bind
bind:
	gomobile bind -o $(OUTDIR)/gomochat.aar

lint:
	gometalinter $(CURDIR)/...

clean:
	rm -rf $(OUTDIR)/*
