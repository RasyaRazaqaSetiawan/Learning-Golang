package main


//  import "fmt"
// // import (
// //	"math/rand"
// // )

//  func main() {
//  	var pemain1, pemain2 string
//  	var pilihanPemain1, pilihanPemain2 int

//  	fmt.Println("Pilihan mau pilihan apa?")
//  	fmt.Println("Ketik 1 untuk Gunting")
//  	fmt.Println("Ketik 2 untuk Kertas")
//  	fmt.Println("Ketik 3 untuk Batu")
//  	fmt.Println("Giliran Pemain 1")

//  	fmt.Scanln(&pilihanPemain1)

//  	switch pilihanPemain1 {
//  	case 1:
//  		pemain1 = "Gunting"
//  	case 2:
//  		pemain1 = "Kertas"
//  	case 3:
//  		pemain1 = "Ba\tu"
//  	default:
//  		fmt.Println("Tidak Ada Pilihan")
//  	}
//  	fmt.Printf("Pilihan pemain satu %s \n", pemain1)

//  	fmt.Println("Giliran Pemain 2")
//  	fmt.Scanln(&pilihanPemain2)

//  	switch pilihanPemain2 {
//  	case 1:
//  		pemain2 = "Gunting"
//  	case 2:
//  		pemain2 = "Kertas"
//  	case 3:
//  		pemain2 = "Batu"
//  	default:
//  		fmt.Println("Tidak Ada Pilihan")
//  	}
	
//  	fmt.Printf("Pilihan pemain dua %s \n", pemain2)

//  	if pemain1 == "Gunting" && pemain2 == "Kertas" {
//  		fmt.Println("Pemain 1 menang")
//  	} else if pemain1 == "Kertas" && pemain2 == "Gunting" {
//  		fmt.Println("Pemain 2 menang")
//  	} else if pemain1 == "Kertas" && pemain2 == "Batu" {
//  		fmt.Println("Pemain 1 menang")
//  	} else if pemain1 == "Batu" && pemain2 == "Kertas" {
//  		fmt.Println("Pemain 2 menang")
//  	} else if pemain1 == "Batu" && pemain2 == "Gunting" {
//  		fmt.Println("Pemain 1 menang")
//  	} else if pemain1 == "Gunting" && pemain2 == "Batu" {
//  		fmt.Println("Pemain 2 menang")
//  	} else if pemain1 == pemain2 {
//  		fmt.Println("Hasilnya Draw")
//  	} else {
//  		fmt.Println("Ada Pemain yang memilih diluar angka")
//  	}
//  }