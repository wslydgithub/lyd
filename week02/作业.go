// package main//只打印到9
//
// import "fmt"
//
//	func main() {
//		ch1 := make(chan byte, 2)
//		ch2 := make(chan byte, 2)
//		go func() {
//			var a byte
//			for a = 'A'; a < a+10; a++ {
//				ch1 <- a
//			}
//		}()
//		go func() {
//			var b byte
//			for b = '0'; b < b+10; b++ {
//				ch2 <- b
//			}
//		}()
//		var i int
//		for i = 0; i < 5; i++ {
//			fmt.Printf("%c", <-ch1)
//			fmt.Printf("%c", <-ch1)
//			fmt.Printf("%c", <-ch2)
//			fmt.Printf("%c", <-ch2)
//		}
//
// }
//package main//打印到字母结束
//
//import "fmt"
//
//func main() {
//	ch1 := make(chan byte, 2)
//	ch2 := make(chan int, 2)
//	go func() {
//		var a byte
//		for a = 'A'; a < a+26; a++ {
//			ch1 <- a
//		}
//	}()
//	go func() {
//		var b int
//		for b = 0; b < b+26; b++ {
//			ch2 <- b
//		}
//	}()
//	var i int
//	for i = 0; i < 13; i++ {
//		fmt.Printf("%c", <-ch1)
//		fmt.Printf("%c", <-ch1)
//		fmt.Printf("%d", <-ch2)
//		fmt.Printf("%d", <-ch2)
//	}
//
//}
