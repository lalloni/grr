bin/grr:
	go build -o bin/grr src/grr.go

clean:
	rm bin/grr

default:
	bin/grr
