build:
	go build -buildmode=c-archive ../restclient.go
	gcc -pthread -I ../ -o rest-c rest.c ./restclient.a
	g++ -pthread -I ../ -o rest-cpp rest.cpp ./restclient.a

run:
	./rest

clean:
	rm rest-c rest-cpp restclient.a restclient.h
