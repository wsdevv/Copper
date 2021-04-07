#include <string>
#include <vector>
#include "../ParserInit/Compiler.h"
using namespace std;
// This file contains all of the functions provided by the compiler (in c++)
string log(string msg, string len) {
   return 
   "   mov ecx,"+msg+";\n"
   "   mov edx,"+len+";\n"
   "   call log       ;\n"
   "   xor ecx, ecx   ;\n"
   "   xor edx, edx   ;\n";
}


//creates a variable
string createVar(int stackpos, string value) {
    return
    "   mov DWORD [esp+"+to_string(stackpos)+"], "+value+";\n";
}

//deletes a variable
string deleteVar(int stackpos) {
    return
    "   mov DWORD [esp+"+to_string(stackpos)+"],"+"0;\n";
}

//formats a memory location
string formatVar(int stackpos) {
    return "[esp+"+to_string(stackpos)+"]";
}

//gets the location of a variable
int getVarLoc(vector<string> names, string name) {
   for (int x = 0; x< names.size(); x++) {
       if (names[x] == removeChr(name, " ", "")) {
           return x;
       }
   }
   return -1;
}

//gets the location the next var has to be in
int movVar(vector<int> locations) {
  int loc = 4;
  for (int x = 0; x< locations.size(); x++) {
       if (locations[x] == loc) {
           loc += 4;
       }
   }
  
   return loc;
}

//creates constant variables for strings and gets any stack vars
vector<string> parse(vector<string>* names,vector<int>* locations, vector<string>* Values, int* Const_index, string to_parse) {
    int task = 0;
    string temp = "";
    string var  = "";
    string finals = "";
    vector<string> real_final;

    //constant names to include
    for (int i=0; i<to_parse.length();i++) {
        if (to_parse[i] == '`') {
            if (task==0) {
                if (removeChr(var, " ", "")!="") {
                  finals += formatVar(getVarLoc(*names, var)) + ",";
                }
                var = "";
                task = 1;
            }
            else {
                //constant name
                string name  = ".Const."+to_string(*Const_index);
                int    exists = -1;

                //checks if a constant exists allready, so does not waste memory
                for (int fc = 0; fc<Values->size();fc++) {
                   
                    if ((*Values)[fc].find(removeChr(("DB \"" + temp + "\""), "%n", "\", 0xA, \""))<removeChr(("DB \"" + temp + "\""), "%n", "\", 0xA, \"").length()) {
                       exists = fc;
                       break;
                    }
                }

                if (exists == -1) {
                  //the value of the new constant
                  Values->push_back(removeChr((name+":\n    DB \"" + temp + "\""), "%n", "\", 0xA, \""));
                
                  //the length of the constant
                  Values->push_back(name+".length equ $-"+name);
                  real_final.push_back(name);
                  finals += name;
                  *Const_index += 1;
                }else {
                   //if value exists, use it instead
                   real_final.push_back(".Const."+to_string(exists));
                   finals += ".Const."+to_string(exists);
                }

                
                temp = "";
                task = 0;
            }

        }
        else if (task==1) {
          temp += to_parse[i];
        }
        else {
           var += to_parse[i];
        }
    }


    //adds/concatinates any variables
    if (removeChr(var, " ", "")!="") {
        finals += formatVar((*locations)[getVarLoc(*names, var)]);
        real_final.push_back(formatVar((*locations)[getVarLoc(*names, removeChr(var, ".length", "")+".length")]));
    }
    
    real_final.push_back(finals);
    return real_final;
}
