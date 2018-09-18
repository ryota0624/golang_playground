build:
	statik -src=static && vgo run main.go

initialize:
	go get -u golang.org/x/vgo
	
lint:
	golint ./... && vgo tool vet ./

format:
	gofmt -s -w . 