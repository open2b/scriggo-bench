// this file was generated by gomacro command: import _b "runtime/debug"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	debug "runtime/debug"
)

// reflection: allow interpreted code to import "runtime/debug"
func init() {
	Packages["runtime/debug"] = Package{
	Name: "debug",
	Binds: map[string]Value{
		"FreeOSMemory":	ValueOf(debug.FreeOSMemory),
		"PrintStack":	ValueOf(debug.PrintStack),
		"ReadBuildInfo":	ValueOf(debug.ReadBuildInfo),
		"ReadGCStats":	ValueOf(debug.ReadGCStats),
		"SetGCPercent":	ValueOf(debug.SetGCPercent),
		"SetMaxStack":	ValueOf(debug.SetMaxStack),
		"SetMaxThreads":	ValueOf(debug.SetMaxThreads),
		"SetPanicOnFault":	ValueOf(debug.SetPanicOnFault),
		"SetTraceback":	ValueOf(debug.SetTraceback),
		"Stack":	ValueOf(debug.Stack),
		"WriteHeapDump":	ValueOf(debug.WriteHeapDump),
	}, Types: map[string]Type{
		"BuildInfo":	TypeOf((*debug.BuildInfo)(nil)).Elem(),
		"GCStats":	TypeOf((*debug.GCStats)(nil)).Elem(),
		"Module":	TypeOf((*debug.Module)(nil)).Elem(),
	}, 
	}
}
