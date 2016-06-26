package main

import (
	"github.com/hyan1979/hello/gotoc"
	"github.com/hyan1979/hello/grammar"
	"github.com/hyan1979/hello/io"
	"github.com/hyan1979/hello/math"
	"github.com/hyan1979/hello/net"
	"github.com/hyan1979/hello/oop"
)

func go_grammar() {
	// Basic
	grammar.Hello()
	grammar.Typesize()
	grammar.Enum()
	grammar.Scope()
	grammar.Scanf()
	grammar.Pointer()
	grammar.Average()
	grammar.Multi_return()
	grammar.First_class()
	grammar.Anonymous_func()

	// Go routine and channel
	grammar.Call_pingpong()

	// Interface
	grammar.GetArea()
}

func go_math() {
	math.Sqrt()
	math.Call_fibonacci()
}

func go_gotoc() {
	// Call c routine in golang
	gotoc.Rand()
}

func go_io() {
	io.Call_read()
}

func go_net() {
	net.Call_http()
}

func go_oop() {
	oop.Call_bicycle()
	oop.Vertex_method()
	oop.Entity_method()
}

func main() {
	go_grammar()
	go_math()
	go_gotoc()
	go_io()
	go_net()
	go_oop()
}
