package main

import "fmt"

func main() {
	fmt.Println(operation(15, 4))
	//TODO
	// projede değişken tanımlaması yok 
	// fonksiyon kullanımı çok yetersiz belirli bir senaryo üzerinde çeşitlendirerek
	// fonksiyonların farklı tanımlama biçimlerinide burada kullanalım
}

func operation(x int, y int) int {
	//TODO
	//fonksiyon ismi ile yapılan işlem arasında tutarsızlık var
	// fonksiyon ismi sadece toplama yapacağını söylüyor 
	fmt.Println("İşlem Başarılı!")
	
	if x < 6 && x > 0 || y > 16 {
		return x + y
	} else if x > 10 || y < 5 {
		return x / y
	} else if x == 10 || y == 4 {
		return x * y
	} else if x < 0 && y > 0 {
		return (x * y * -1)
	} else {
		return 0
	}
}
