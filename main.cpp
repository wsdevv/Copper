#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <functional>
#include "./Init/Compiler.cpp"
/*
 * NOTICE:
 * The Copper programming language is in development,
 * The compiler (this code) does not yet "compile" anything,
 * This project will start accepting contributions when
 * a. It gets further into development
 * b. The creator adds more commets
 * c. The creator makes a contribution and use wiki
 * d. The creator has motivation
 * 
 * The compiler is not that efficient, but it works, haha!
 * Thank you for visiting! I suggest you visit again when this project progresses! 
*/

using namespace std;
int main(int argc, char* argv[]) {
  ifstream getf(argv[1]);
  ifstream std("./libs/std.asm");
  ofstream out("./out.asm");
  string ff = "";
  string fstd = "";

  //puts std file and user input file into string
  if (getf && std) {
     ostringstream full_file;
     full_file << getf.rdbuf();
     ff = full_file.str();
     ostringstream full_std;
     full_std << std.rdbuf();
     fstd = full_std.str();
  }
  else {
     return -1;
  }

  //just in case line does not include ";" add ";" there
   if (removeChr(ff, " ", "")[removeChr(ff, " ", "").length()-1] != ';' && removeChr(ff, " ", "")[removeChr(ff, " ", "").length()-1] != '}') {
      ff+=";";
   }
 

  //replaces "\n" with empty, if had "\n" would mess up compilation, thats why you use %n instead of \n for new lines
  while (ff.find("\n") != string::npos) {
     if (removeChr(ff, " ", "")[removeChr(ff, " ", "").find("\n")-1] == ';' || removeChr(ff, " ", "")[removeChr(ff, " ", "").find("\n")-1] == '{' || removeChr(ff, " ", "")[removeChr(ff, " ", "").find("\n")-1] == '}') {
       ff.replace(ff.find("\n"), 1, "");
     }
     else {
       ff.replace(ff.find("\n"), 1, ";");
     }
  }
  ff = removeChr(ff, "\t","");
  vector<string> names;
  vector<int>    pointers;
  vector<double> parent;
  

  // main is what is in _global start function and functions are just in the _section.text
  vector<string> main;
  vector<string> functions;

  Compile("!@auto_start@!", &names, &pointers, &parent, &main, ff);

 
  //for the global start funcion
  string start = "";

  //constants for the current function
  string constst = "";
  for (string x: main) {
     start += x;
  }
  out <<
  "section .text\n"
  +fstd+
  start+
  constst;
  
  out.close();
  std.close();
  getf.close();
  system("nasm -f elf32 -g out.asm");
  system("ld -m elf_i386 -s -o out out.o");
  
  return 0;
}