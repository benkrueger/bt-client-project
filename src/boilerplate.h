//This h file will include error checking and quick filesystem functions so we don't have to put it in the main
#ifndef BOILERPLATE_H
#define BOILERPLATE_H
#include <sys/stat.h>

bool check_paths( char * dirpath,  char * filepath)
{
  struct stat buffer;
  if(stat (dirpath, &buffer) == 0){
      if(!S_ISDIR(buffer.st_mode)){
          return false;
      }
  }
  if(stat (filepath, &buffer) == 0){
      if(!S_ISREG(buffer.st_mode)){
          return false;
      }
  }
  return true;
}
#endif