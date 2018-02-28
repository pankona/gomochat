

OUTDIR = $(CURDIR)/android/app/libs
AAR    = $(OUTDIR)/gomochat.aar

all: cmd bind
	cd $(CURDIR)/android && ./gradlew build

.PHONY: cmd
cmd:
	cd $(CURDIR)/cmd/gomochat && go build

.PHONY: bind
bind: $(AAR)

$(AAR): $(CURDIR)/gomochat.go
	mkdir -p $(OUTDIR)
	gomobile bind -o $(OUTDIR)/gomochat.aar

.PHONY: lint
lint:
	gometalinter $(CURDIR)/... --exclude='cmd/server'

clean:
	rm -rf $(OUTDIR)/*
	cd $(CURDIR)/android && ./gradlew clean
