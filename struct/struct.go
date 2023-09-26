package main

import "fmt"



type Mahasiswa struct{
	NIM			int
	FirstName 	string
	LastName	string
	Email		string
	IsActive	bool
}

// Ini Method
func (mhs Mahasiswa) tampil() string {
	result := fmt.Sprintf("Nim : %d\nFirstName : %s\nLastName : %s\nEmail : %s\nActive : %t\n", mhs.NIM, mhs.FirstName, mhs.LastName, mhs.Email, mhs.IsActive)
	return result
}

func myMhs(mhs Mahasiswa) string {
	result := fmt.Sprintf("Nim : %d\nFirstName : %s\nLastName : %s\nEmail : %s\nActive : %t\n", mhs.NIM, mhs.FirstName, mhs.LastName, mhs.Email, mhs.IsActive)
	return result
}

type Kelas struct{
	Name		string
	Ketua		Mahasiswa
	Mhs			[]Mahasiswa
	IsActive 	bool
}

// Ini Method
func (kelas Kelas) display()  {
	fmt.Printf("Kelas : %s\n", kelas.Name)
	fmt.Printf("Jumlah Mahasiswa : %d", len(kelas.Mhs))
}

func displayKelas(kelas Kelas){
	fmt.Printf("Kelas : %s\n", kelas.Name)
	fmt.Printf("Jumlah Mahasiswa : %d", len(kelas.Mhs))
}



func main(){

	// inisalisasi object
	mhs := Mahasiswa{1212, "Joko", "Susilo", "jokos@gmail.com", true}
	mhs1 := Mahasiswa{1213, "Ferdinand", "Susilo", "ferdinand@gmail.com", false}
	result := mhs.tampil()
	result1 := mhs1.tampil()

	fmt.Println(result, result1)

	myMhs := []Mahasiswa{mhs, mhs1}

	kelas := Kelas{"10.7A.25", mhs, myMhs, true}
	kelas.display()


}




