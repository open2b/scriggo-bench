// this file was generated by gomacro command: import _b "net/http/cgi"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	cgi "net/http/cgi"
)

// reflection: allow interpreted code to import "net/http/cgi"
func init() {
	Packages["net/http/cgi"] = Package{
	Name: "cgi",
	Binds: map[string]Value{
		"Request":	ValueOf(cgi.Request),
		"RequestFromMap":	ValueOf(cgi.RequestFromMap),
		"Serve":	ValueOf(cgi.Serve),
	}, Types: map[string]Type{
		"Handler":	TypeOf((*cgi.Handler)(nil)).Elem(),
	}, 
	}
}
