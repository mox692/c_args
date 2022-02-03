b:;
	go build -buildmode=c-archive -o lib.a example.go
	gcc -pthread -o example example.c lib.a

clean:;
	rm lib* example