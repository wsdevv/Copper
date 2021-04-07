#include "../Init/Functions.cpp"
#include "../ParserInit/Parse_Rules.h"
#include <iostream>
#include <vector>
#include <random>
#include <algorithm>

//hippity hoppity this code is now my property
double fRand(double fMin, double fMax)
{
    double f = (double)rand() / RAND_MAX;
    return fMin + f * (fMax - fMin);
}

Compiler::Compiler() {
   this->current_index = 0;
   this->last_index    = 0;
   this->temp_storage  = "";
};
Compiler* Compiler::set(string f) {
    this->read_file = f;
    return this;
}
Compiler* Compiler::set(int i) {
   this->current_index = i;
   return this;
}

DataTransfer::DataTransfer() {};

DataTransfer Compiler::Start_Parse() {
    return (new DataTransfer())->init(*this, true);
};

//checks if the next string is equal to another string
DataTransfer DataTransfer::next(string st) {
    
    if (this->transfer == true) {
     this->c.last_index = this->c.current_index;
     //keeps on checking for that string until the end or break
     while (removeChr(this->c.temp_storage, " ", "") != st) {
          
          if (this->c.current_index >= this->c.read_file.length()) {
             this->c.temp_storage = "";
             this->c.current_index = this->c.last_index;
             return (new DataTransfer())->init(this->c,false);
          }
          else {
            this->c.temp_storage  += this->c.read_file[this->c.current_index];
            this->c.current_index += 1;
          }
      }
      this->c.temp_storage = "";
      return (new DataTransfer())->init(((*this).c), true);
    }
    return (new DataTransfer())->init(((*this).c), false);
};

//gets everything until a certain string
DataTransfer DataTransfer::get_until(string st) {

if(this->transfer == true) {
 
    this->c.last_index = this->c.current_index;
    int xn = this->c.current_index;
    //the final storage value
    string fin = "";

    while(this->c.temp_storage != st) {
        
        if (this->c.current_index > this->c.read_file.length() || this->c.current_index + st.length() > this->c.read_file.length()) {
             this->c.temp_storage  = "";
             this->c.current_index = this->c.last_index;
             return (new DataTransfer())->init(false);
        }
        else {
            //if the temporary holder is less than the token length, keep adding to it
            if (this->c.temp_storage.length()<st.length()) {
              
              this->c.temp_storage  += this->c.read_file[this->c.current_index +xn];
              xn += 1;
            }
            //if not equal to token yet, reset temp storage and add the current index by one
            else  {
              this->c.temp_storage  = "";
              fin += this->c.read_file[this->c.current_index];
              this->c.current_index += 1;
              xn = 0;
            }
            if(this->c.temp_storage == st) {
                   break;
            }
            
        }
    }
    this->c.current_index += st.length();
    this->c.data.push_back(fin);
    this->c.temp_storage  = "";
    return (new DataTransfer())->init(((*this).c), true);

}
return (new DataTransfer())->init(((*this).c), false);
};


//the actual compiler, found in Compiler.cpp, EDIT HERE IF YOU WANT TO CHANGE ANYTHING
void Compile(
  string         name,
  vector<string> *names, 
  vector<int>    *pointers, 
  vector<double> *parent, 
  vector<string> *start,
  string ff) {
    
    vector<string> Const_names;
    vector<string> Const_values;
    int Current_Const_index = 0;
    double id = fRand(-9000.0, 9000.0);

    Compiler* comp = new Compiler();
    comp->set(ff);
    
    vector<string> main;
    

     //reads the file and converts it to assembly, the real part of the compiler 
   while (true) {

      //erases everything once a file gets to the end
    if (comp->current_index > comp->read_file.length()-2) {
        int f = 0;
        //deletes all of the variables
        for (double p: *parent) {
           if (p == id) {
              main.push_back(deleteVar((*pointers)[f]));
              
              names->erase( names->begin()+f);
              pointers->erase( std::remove( pointers->begin(), pointers->end(), (*pointers)[f] ), pointers->end() );
              parent->erase( std::remove( parent->begin(), parent->end(), (*parent)[f] ), parent->end() );
           }
           f++;
        }
        break;
     }
     

     //reads a function
     comp = comp->Start_Parse()
     .next(Parse_Rules::declare_function)
     .get_until(Parse_Rules::set_function_params)
     .get_until(Parse_Rules::start_chunk)
     .get_until(Parse_Rules::end_chunk)
     .empty([comp,name,&names, &pointers, &parent, &start](vector<string> x) {
        //compiles a different function
       Compile(removeChr(x[0], " ",""), names, pointers, parent, start, x[2]);
     });
   
   //runs a function
     comp = comp->Start_Parse()
     .next(Parse_Rules::run_function)
     .get_until(Parse_Rules::set_function_params)
     .get_until(Parse_Rules::end_line)
     .empty([&main, &pointers, &names, &Const_values, &Current_Const_index, name](vector<string> x) {
        //tests if it is a preset function
        //TODO: PROBABLY NEED TO MAKE PRESET FUNCTIONS LOCATED SOMEWHERE ELSE
        if (removeChr(removeChr(x[0], " ", ""), "\t", "") == "log") {
           vector<string> val = parse(names, pointers,  &Const_values, &Current_Const_index,  x[1]);
         
            //adds up the constants to make the length
            string length = "";
            for (int constant=0;constant<val.size()-1;constant++) {
            
              if (val[constant].find(".Const.")<val[constant].length()-1) {
               if (length.length() == 0) {
                  length += val[constant]+".length";
               }
                else {
                  length += "+"+val[constant]+".length";
                }
              }
              else {
                   if (length.length() == 0) {
                    
                  length += val[constant]+"";
               }
                else {
                  length += "+"+val[constant]+"";
                }
              }
            }

           main.push_back(log(val[val.size()-1], length));
        }
        else {
            main.push_back("   call "+removeChr(x[0], " ", "")+"\n");
        }
     });

     
      
     //gets a varaible
     comp = comp->Start_Parse()
     .next(Parse_Rules::declare_variable)
     .get_until(Parse_Rules::set_variable)
     .get_until(Parse_Rules::end_line)
     .empty([&names, &pointers, &main, &Current_Const_index, &Const_values, &parent, id, comp](vector<string> x) {
        int loc = movVar(*pointers);
        vector<string> val = parse(names,  pointers,  &Const_values, &Current_Const_index,  x[2]);

        //adds up the constants to make the length
        string length = "";
  
        for (int constant=0;constant<val.size()-1;constant++) {
           if (length.length() == 0) {
              length += val[constant]+".length";
           }
           else {
              length += "+"+val[constant]+".length";
           }
        }
        //pushes the variable
        main.push_back(createVar(loc,val[val.size()-1]));
        names->push_back(removeChr(x[1], " ", ""));
        pointers->push_back(loc);
        parent->push_back(id);

        //length of the variable (if string)
        loc = movVar(*pointers);
        main.push_back(createVar(loc,length));
        names->push_back(removeChr(x[1]+".length", " ", ""));
        pointers->push_back(loc);
        parent->push_back(id);

     });
     
     //lets the user be able to change the syntax
     comp = comp->Start_Parse()
     .next(Parse_Rules::change_properties)
     .get_until(Parse_Rules::end_line)
     .empty([](vector<string> x){
        cout << "found: " << x[0] << "\n";
     });
 
  }
  //for the global start func
  string st = "";

  //constants, and compiled code for the current function
  string constst = "";
  for (string x: main) {
     st += x;
  }
  for (string x: Const_values) {
     constst += "\n"+x;
  }

  //initializing the program
  //TODO: Fix no entry mode, some stuff does not work
  if (name=="!@auto_start@!" && comp->read_file.find("#no-entry") == std::string::npos || name=="entry") {

  //Checks for entry errors, so no entry functions can be included
  if (comp->read_file.find("#no-entry") != std::string::npos && name=="entry") {
    cout << "Cppr Err:\n   cannot make entry function when code is already an entry point.\n"<<"\n";
    exit(-1);
  }
  else {
  start->push_back("\nglobal _start");
  start->push_back(
      "\n_start:\n"
      +st+
      "   xor ebx,ebx\n"
      "   mov eax, 1\n"
      "   int 0x80\n" 
      +constst
  );
  }
  }
  else {

     //custom functions that have been compiled
     if (name!= "!@auto_start@!") {
      start->push_back(
        "\n"+name+":\n"
         +st +
         "   ret ;"
         +constst
      );
     }
     else {
        start->push_back(
         st
         +constst
      );
     }
  }


};