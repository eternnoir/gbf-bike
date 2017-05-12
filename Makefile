# This is how we want to name the binary output
BINARY=gbfbike

BUILDFOLDER = build/bin

# These are the values we want to pass for Version and BuildTime
VERSION=0.0.1

# goxc flag
GOXCFLAG= -tasks-=validate -pv=${VERSION} -d ${BUILDFOLDER}

default:
	go build -o ${BUILDFOLDER}/${VERSION}/${BINARY} *.go
	@echo "Your binary is ready. Check "${BUILDFOLDER}/${VERSION}/${BINARY}

test:
	go test -v `go list ./... | grep -v vendor`

cross-all:
	goxc ${GOXCFLAG}

