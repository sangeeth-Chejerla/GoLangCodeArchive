package main

import "fmt"
const (
	Asia =2.0 * iota
	Africa 
	Australia
	Antarctica
	Europe
	NorthAmerica
	SouthAmerica
)	

const (
	PacificOcean = iota
	_
	AtlanticOcean
	IndianOcean
	ArcticOcean
	AntarcticaOcean

)

func main(){
	fmt.Printf("Asia = %v [%T]\n",Asia,Asia)
	fmt.Printf("Africa = %v [%T]\n",Africa,Africa)
	fmt.Printf("Australia = %v [%T]\n",Australia,Australia)
	fmt.Printf("Antarctica = %v [%T]\n",Antarctica,Antarctica)
	fmt.Printf("Europe = %v [%T]\n",Europe,Europe)
	fmt.Printf("NorthAmerica = %v [%T]\n",NorthAmerica,NorthAmerica)
	fmt.Printf("SouthAmerica = %v [%T]\n",SouthAmerica,SouthAmerica)

	fmt.Printf("PacificOcean = %v [%T]\n",PacificOcean,PacificOcean)
	fmt.Printf("AtlanticOcean = %v [%T]\n",AtlanticOcean,AtlanticOcean)
	fmt.Printf("IndianOcean = %v [%T]\n",IndianOcean,IndianOcean)
	fmt.Printf("ArcticOcean = %v [%T]\n",ArcticOcean,ArcticOcean)
	fmt.Printf("AntarcticaOcean = %v [%T]\n",AntarcticaOcean,AntarcticaOcean)
	
}
