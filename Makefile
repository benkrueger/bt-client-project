CC=g++
<<<<<<< HEAD
DEPS=src/main.cpp
CFLAGS=-o
EXE=client

all: ${DEPS}
	${CC} ${DEPS} ${CFLAGS} ${EXE}

clean: ${EXE}
	rm -rf ${EXE}
=======
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
>>>>>>> 0e04ebb83617bde87267cbf485fe651eb5935a93
