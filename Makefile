# makefile has a very stupid relation with tabs,
# all actions of every rule are identified by tabs ......
# and No 4 spaces don't make a tab,
# only a tab makes a tab...
#
# to check I use the command 'cat -e -t -v makefile_name'
# It shows the presence of tabs with ^I and line endings with $
# both are vital to ensure that dependencies end properly and tabs mark the action for the rules so that they are easily identifiable to the make utility.....

.PHONY: build compile

all: compile

build:
	go build -o applications/app/cmd/main applications/app/cmd/main.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o applications/app/cmd/main-linux  applications/app/cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o applications/app/cmd/main-darwin applications/app/cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o applications/app/cmd/main-win applications/app/cmd/main.go
