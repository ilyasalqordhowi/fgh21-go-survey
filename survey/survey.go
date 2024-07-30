package survey

import "fmt"

type Bio struct {
	name         string
	age          int
	gender       string
	isSmoker     bool
	cigarVariant []string
	
}


func Survey(name string,age int,gender string,isSmoker bool) {
	survey := []Bio{
		Bio{
			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
			
		
		},
		Bio{
			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: false,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
		},
		Bio{
			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
		},
		Bio{
			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
		},
		Bio{
			name:     "ilyas ",
			age:      18,
			gender:   "laki-laki",
			isSmoker: false,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
				
			},
			
		},
	}
	for _, j := range survey{
		fmt.Printf("%s,%d,%s,%t ", j.name,j.age,j.gender,j.isSmoker)
		for _, y := range j.cigarVariant {
			fmt.Printf(" %s ", y)
		}
		fmt.Println("")
	}
}
func SurveyCigar(){
	var name string
	var age int
	var gender string
	var isSmoker bool
	fmt.Print("siapa nama anda :")
	fmt.Scanln(&name)
	fmt.Print("berapa umur anda :")
	fmt.Scanln(&age)
	fmt.Print("apa jenis kelamin anda :")
	fmt.Scanln(&gender)
	fmt.Print("apakah anda perokok:")
	fmt.Scanln(&isSmoker)
	Survey(name,age,gender,isSmoker)


}