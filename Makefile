.PHONY: all clean idl uidl

OUTPUT=bup

all: clean
	go build -o ./example/${OUTPUT} main.go

clean:
	rm -f example/${OUTPUT}