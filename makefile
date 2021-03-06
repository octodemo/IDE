# Go parameters
GOCMD=go
GOS=gos
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=IDE
GOSOUTPUT=server_out.go

all: deps build test
build: 
		$(GOS) --export
		$(GOBUILD) -o $(BINARY_NAME) -v       
test: 
		$(GOTEST) -bench=. -benchmem
		$(GOTEST) ./...
clean: 
		$(GOCLEAN)
		rm -f bindata.*
		rm -f $(GOSOUTPUT)
		rm -f $(BINARY_NAME)
run:
		$(GOS) --export
		$(GOBUILD) -o $(BINARY_NAME) -v
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/cheikhshift/gos
		$(GOS) deps
uiview:
		npm install

