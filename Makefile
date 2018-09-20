build:
	statik -src=static && sh app.sh

initialize:
	go get -u golang.org/x/vgo
	
lint:
	golint ./... && vgo tool vet ./

format:
	gofmt -s -w . 