// Code generated by 'yaegi extract compress/lzw'. DO NOT EDIT.

// +build go1.16,!go1.17

package stdlib

import (
	"compress/lzw"
	"reflect"
)

func init() {
	Symbols["compress/lzw/lzw"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"LSB":       reflect.ValueOf(lzw.LSB),
		"MSB":       reflect.ValueOf(lzw.MSB),
		"NewReader": reflect.ValueOf(lzw.NewReader),
		"NewWriter": reflect.ValueOf(lzw.NewWriter),

		// type definitions
		"Order": reflect.ValueOf((*lzw.Order)(nil)),
	}
}
