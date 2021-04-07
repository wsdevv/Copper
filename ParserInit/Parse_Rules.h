#include <string>
using namespace std;
namespace Parse_Rules {
    static string declare_variable    = "var:";
    static string set_variable        = "$=";
    static string add_var             = "$+";
    static string subtract_var        = "$-";
    static string multiply_var        = "$*";
    static string devide_var          = "$/";

    static string declare_function    = "chunk:";
    static string set_function_params = "~";
    static string run_function        = "run:";
    static string start_chunk         = "{";
    static string end_chunk           = "}";

    static string end_line            = ";";
    static string smart_pointer       = "@";

    static string change_properties   = "#";
}