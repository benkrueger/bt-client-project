CC=g++
DEPS=src/main.cpp
CFLAGS=-o
EXE=client

all: ${DEPS}
	${CC} ${DEPS} ${CFLAGS} ${EXE}

clean: ${EXE}
	rm -rf ${EXE}