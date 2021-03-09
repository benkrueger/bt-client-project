#include <iostream>
#include "boilerplate.h"
using namespace std;

int main(int argc,char *argv[]){
    string dirpath;
    string filepath;
    if (argc != 3 || !check_paths(argv[2],argv[1])){
        std::cerr << "Incorrect arguments, usage :" << argv[0] << " <file.torrent> <download directory>" <<std::endl;
        return 1;
    }else{
        dirpath = std::string(argv[2]);
        filepath = std::string(argv[1]);
    }
    //First thing we need to do is parse the Torrent file.
    
    return 0;
}