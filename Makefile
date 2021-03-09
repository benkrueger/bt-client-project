CC=g++
DEPS=src/main.cpp lib/include/bencode.hpp
CFLAGS=-o
EXE=client
LIB=lib
REMOTES=https://github.com/jimporter/bencode.hpp.git
all: ${DEPS} 
	${CC} ${DEPS} ${CFLAGS} ${EXE}

fetch: ${LIB}
	git clone ${REMOTES} ${LIB} 

clean: ${EXE}
	rm -rf ${EXE} ${LIB}