build:
	go build -buildmode=c-archive ../restclient.go
	gcc -pthread -I ../ -o rest rest.c ./restclient.a

run:
	./rest

clean:
	rm rest restclient.a restclient.h
