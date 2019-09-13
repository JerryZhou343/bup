.PHONY: all clean idl uidl

OUTPUT=prototool

all: clean
	go build -o ./example/${OUTPUT} main.go config.go desc_source.go helper.go

clean:
	rm -f example/${OUTPUT}