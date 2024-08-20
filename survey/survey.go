package survey

import "fmt"

type Bio struct {
	name         string
	age          string
	gender       string
	isSmoker     string
	cigarVariant []string
	
}


func Survey(name string,age string,gender string,isSmoker bool) {	
}
func SurveyCigar(){
	data := []Bio{}
	var input string
	var name string
	var age string
	var gender string
	var isSmoker string
	var  variant string
	
	fmt.Print("1.Input survey")
	fmt.Scanln(&input)
	fmt.Print("siapa nama anda? :")
	fmt.Scanln(&name)
	fmt.Print("berapa umur anda? :")
	fmt.Scanln(&age)
	fmt.Print("apa jenis kelamin anda? :")
	fmt.Scanln(&gender)
	fmt.Print("apakah anda perokok? :")
	fmt.Scanln(&isSmoker)
	cigarVariant := []string{}
	cont := true
	for cont{
		fmt.Scanln(&variant)
		if variant == "0" {
			cont = false
			}else{
				cigarVariant =append(cigarVariant,variant)
				fmt.Print("rokok apa yang pernah anda coba? :")	
			}
		}
	
		data = append(data, Bio{name: name,age: age,gender: gender,isSmoker: isSmoker,cigarVariant:  cigarVariant})
		fmt.Print(data)
} 
