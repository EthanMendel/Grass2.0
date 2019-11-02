.PHONY: all run clean prepare  binary

#####
# EXE - Name of executable
#####
EXE?=app



#####
# GOPATH - path to go code. All code needs to be under the GOPATH/src directory in the appropriate location
# If not set then it will default to "../../../../"  i.e. (assuming in gopath/src/github.ibm.com/Bluemix/content-mgmt as current directory)
# will take it back "gopath" 
#
# GOPATH?=../../../../
GOPATH?=/Users/gmendel/GO

#####
# Define DEBUG to have go compile with debug flags
#####
ifdef DEBUG
	DBG_FLAGS:=-gcflags "-N -l"
endif

#####
# Define GO_UPDATE if running in an existing project area AND you know there are updates to the current vendor projects. This will force them to be updated.
# By default only missing vendor packages are downloaded.
#####
ifdef GO_UPDATE
	updateGoPackages:=-u
endif

all: run

run: binary
	./$(EXE)

clean:
	echo "got to clean"
	-@rm -f $(EXE)
	-@rm -f $(GOPATH)/bin/$(EXE)
	-@rm -f debug
	-@find . -name '.DS_Store' -exec rm {} \;

prepare:
	go get $(updateGoPackages) -d ./...
	

binary: clean prepare
	go build $(DBG_FLAGS) -o $(EXE)

