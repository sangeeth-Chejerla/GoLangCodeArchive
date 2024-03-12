package main

import "fmt"

//constants
const avogadro float64 = 6.0221413e+23
const grams =100.0

//a type
type amu float64

//a method
func (mass amu) float() float64 {
	return float64(mass)
}

// a struct
type metalloid struct{
	name string
	number int32
	weight amu
}

//slice
var metalloids []metalloid

func addMetalloid(name string,number int32 , weight amu){
	//append(slice_to_modify, elements_to_append...)
	metalloids = append(metalloids,metalloid{name,number,weight})
}


//finds # of moles
//functions
func moles(mass amu)float64{
	return float64(mass)/grams
}

func atoms(moles float64)float64{
	return moles * avogadro
}

func headers() string {
    return fmt.Sprintf(
        "%-10s %-10s %-10s Atoms in %.2f Grams\n",
        "Element", "Number", "AMU", grams,
    )
}

func main(){
	addMetalloid("Boron",5,10.81)
	addMetalloid("Silicon", 14, 28.085)
    addMetalloid("Germanium", 32, 74.63)
    addMetalloid("Arsenic", 33, 74.921)
    addMetalloid("Antimony", 51, 121.760)
    addMetalloid("Tellerium", 52, 127.60)
    addMetalloid("Polonium", 84, 209.0)


fmt.Print(headers())

for _,m := range metalloids{
	fmt.Printf(
		"%-10s %-10d %-10.3f %e\n",
		m.name, m.number, m.weight.float(), atoms(moles(m.weight)),
	)
}
}