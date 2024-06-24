package main

import "testing"

func BenchmarkSlugify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slugify("A 100x Investment (2019)")
	}
}

// func BenchmarkSlugifyRegExp(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		slugifyRegExp("A 100x Investment (2019)")
// 	}
// }

// func BenchmarkSlugifyMy(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		slugifyMy("A 100x Investment (2019)")
// 	}
// }
func BenchmarkSlugifyMyNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slugifyMyNew("A 100x Investment (2019)")
	}
}
// func BenchmarkH(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		test("A 100x Investment (2019)")	
// 	}
// }