CC=g++
DEPS=src/main.cpp lib/include/bencode.hpp
CFLAGS=-o
EXE=client
LIB=lib
REMOTES=https://github.com/jimporter/bencode.hpp.git
TEMPFILES=*.log *.aux
REPORTS=reports/
TEXC=pdflatex
TEXFILES=*.tex

all: ${DEPS} 
	${CC} ${DEPS} ${CFLAGS} ${EXE}

fetch: ${LIB}
	git clone ${REMOTES} ${LIB} 

clean: ${EXE} ${REPORTS} ${TEMPFILES}
	rm -rf ${EXE} ${LIB}
	rm -rf ${TEMPFILES}
	rm -rf ${REPORTS}
reports: ${TEXFILES}
	${TEXC} ${TEXFILES}
	mv *.pdf ${REPORTS}
