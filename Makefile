.POSIX:

build:
	go generate
	go build

install:
	cp ./md2pdf /usr/local/bin/md2pdf

clean:
	rm ./md2pdf ./themes.go
