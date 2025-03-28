package main
import "fmt"

func profit(awal, akhir float64){
	if awal > akhir {
		awal = awal - akhir
		fmt.Print("penurunan sebesar: ", awal)
	}else if akhir > awal{
		akhir -= awal
		fmt.Print("peningkatan sebesar: ", akhir)
	}else{
		fmt.Print("tetap")
	}
}
func main(){	
	var n, num float64
	fmt.Scan(&n, &num)
	profit(n, num)
}