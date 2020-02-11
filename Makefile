.PHONY: all clean idl uidl

OUTPUT=prototool

all: clean
	go build -o ./example/${OUTPUT} main.go

clean:
	rm -f example/${OUTPUT}