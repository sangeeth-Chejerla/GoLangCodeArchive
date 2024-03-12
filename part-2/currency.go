package main

import (
	"math/rand"
)

type Curr struct{
	Currency string
	Name     string
	country  string
	Number   int
}

var Currencies = []Curr{
	Curr{"NOK", "Norwegian Krone", "Norwary", 578},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"KHR", "Riel", "Cambodia", 116},
	Curr{"SZL", "Lilangeni", "Swaziland", 748},
	Curr{"GBP", "Pound Sterling", "Isle of Man", 826},
	Curr{"BWP", "Pula", "Botswana", 72},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"TRY", "Turkish Lira", "Turkey", 949},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"JMD", "Jamaican Dollar", "Jamaica", 388},
	Curr{"ALL", "Lek", "Albania", 8},
	Curr{"GEL", "Lari", "Georgia", 981},
	Curr{"KFM", "Coromo Franc", "Comoros", 174},
	Curr{"NZD", "New Zeland Dollar", "Tokelau", 554},
}

func main(){
	sort()
	print()
	print()
}

func shuffle(){
	n := len(Currencies)
	for i := range(Currencies){
		next := rand.Intn(n)
		temp := Currencies[i]
		Currencies[i] = Currencies[next]
		Currencies[next] = temp
	}
}

func sort(){
	N:= len(Currencies)
	for i:=0; i < N-1; i++{
		CurrMin := i
		for k:=i+1; k<N; k++{
			if Currencies[k].Number < Currencies[CurrMin].Number{
				CurrMin = k
			}
		}
		if CurrMin != i{
			temp := Currencies[i]
			Currencies[i] = Currencies[CurrMin]
			Currencies[CurrMin] = temp
		}
	}
}
