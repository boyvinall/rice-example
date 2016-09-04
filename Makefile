BIN:=$(shell basename $(realpath .))

all: $(BIN)

$(BIN):
	go build
	rice append --exec $(BIN)

clean:
	rm -f $(BIN)

list:
	unzip -l $(BIN)

run:
	./$(BIN)

test:
	./$(BIN) &
	sleep 1
	curl http://localhost:8080
	pkill $(BIN)
