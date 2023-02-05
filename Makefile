hello_:
	go build -o ./hello_server ./hello/hello.go

concdl_:
	go build -o ./concurrent_dl ./concdl/dl.go

books_:
	cd books && go build -o ../books_server books.go && cd ..

csv_:
	go build -o ./csv_parser ./csv/csv.go

all:
	make hello_
	make concdl_
	make books_
	make csv_
