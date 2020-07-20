RM := rm -f
GO := go

SOURCES := bytes.go
TEST_SOURCES := bytes_test.go

ALL_SOURCES := $(SOURCES) $(TEST_SOURCES)

.PHONY:	all clean

all:	test

test:	$(ALL_SOURCES)
	$(GO) test -coverprofile=coverage.txt
	$(GO) tool cover -html=coverage.txt -o coverage.html
	$(GO) tool cover -func=coverage.txt
	@echo "Run your borwser to see coverage.html"

clean:
	$(RM) coverage.txt coverage.html
