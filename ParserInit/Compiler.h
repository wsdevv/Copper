#include <string>
#include <vector>
#include <functional>
using namespace std;
class DataTransfer;
class Compiler {
    public:
    string read_file;
    string temp_storage;
    vector<string> data;

    //current index in file
    int    current_index;
    int    last_index;

    Compiler();
    Compiler*        set(string);
    Compiler*        set(int);
    DataTransfer Start_Parse();
    

};

//just for chaining things, its kinda satisfying
class DataTransfer {
    public:
    Compiler c;
    bool     transfer;
    DataTransfer();
    DataTransfer init(Compiler comp, bool t) {
        this->c = comp;
        this->transfer = t;
        return *this;
    }
    DataTransfer init(bool t) {
        this->transfer = t;
        return *this;
    }

    //gets the next token, removes all spaces to find it
    DataTransfer next(string);

    /*gets space until token
    ex:
         hello = <- "=" is token, so gets " hello "
    */
    DataTransfer get_until(string);
    

    template<typename F>
    //runs function and deletes self
    Compiler* empty(F lambada){
        if (this->transfer == true) {
            this->c.temp_storage = "";
            lambada(this->c.data);
            this->c.data.clear();

            
     }
        this->c.temp_storage = "";
        
        return (new Compiler())->set(this->c.read_file)->set(this->c.current_index);
    }
};



string removeChr(string str, string fi, string ot) {
     string l = str;
     while (l.find(fi) != std::string::npos) {
         l.replace(l.find(fi), fi.length(), ot);
     }
     return l;
}