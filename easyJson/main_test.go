package main

import "testing"

func BenchmarkMyJson(b *testing.B) {


	for n := 0; n < b.N; n++ {

		easyJsonHandler()
	}


}


func BenchmarkNativeJson(b *testing.B) {


	for n := 0; n < b.N; n++ {
		jsonNativeHandler()

	}


}