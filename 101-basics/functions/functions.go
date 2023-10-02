package main

import (
	"fmt"
	"time"
)

func main() {
	returnTypeTest(1)
	timer(time.After(3 * time.Second))
}

//func timer(c <-chan time.Time, messages ...string) { // messages alanına slice vererek dilediğin kadar string parametresi gönderebilirsin
//	for {
//		select {
//		case <-c:
//			return
//		default:
//			fmt.Println(time.Now())
//			time.Sleep(1 * time.Second)
//		}
//	}
//}

//func timer(c <-chan time.Time) {
//	defer fmt.Println("sonraki asamaya geciyor") // fonksiyon bittiğinde çalışır
//	defer fmt.Println("timer bitti")             // en son yazan defer ilk calisir (FILO)
//
//	for {
//		select {
//		case <-c:
//			return
//		default:
//			fmt.Println(time.Now())
//			time.Sleep(1 * time.Second)
//		}
//	}
//}

func timer(c <-chan time.Time) {
	defer fmt.Println("sonraki asamaya geciyor")
	defer fmt.Println("timer bitti")

	defer func() { // bu fonksiyonu başta yazmalıyız. panic atmaya müsait bir yapı varsa recover kullanılır
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)

			if time.Now().Day() == 2 {
				panic("beklenmeyen bir sorun olustu")
			}
		}
	}

}

func returnTypeTest(val int) (res string) {
	if val == 10 {
		return "++++++++"
	}

	return "---------"
}
