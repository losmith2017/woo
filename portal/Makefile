GOPATH := $(shell pwd)
PORT    = 1380
SRCDIR  = . 

start:
	go run portal.go $(PORT) $(SRCDIR) &
	sleep 3 && open http://localhost:$(PORT)

stop:
	kill -HUP $(shell lsof -t -i:$(PORT))

