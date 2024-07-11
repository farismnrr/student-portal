package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	// Method Login untuk proses login mahasiswa
	Login(id string, name string) error
	// Method Register untuk proses registrasi mahasiswa
	Register(id string, name string, studyProgram string) error
	// Method GetStudyProgram untuk mendapatkan nama program studi berdasarkan kode
	GetStudyProgram(code string) (string, error)
	// Method ModifyStudent untuk memodifikasi data mahasiswa
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	// slice students untuk menyimpan data mahasiswa
	students []model.Student
	// map studentStudyPrograms untuk menyimpan hubungan antara kode program studi dan nama program studi
	studentStudyPrograms map[string]string
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	// Kembalikan pointer ke instance baru dari InMemoryStudentManager
	return &InMemoryStudentManager{
		// Inisialisasi slice students dengan data mahasiswa
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		// Inisialisasi map studentStudyPrograms dengan hubungan kode program studi dan nama program studi
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
	}
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	// Kembalikan slice students dari InMemoryStudentManager
	return sm.students
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	// Iterasi melalui slice students
	for _, student := range sm.students {
		// Periksa apakah ID dan nama mahasiswa cocok dengan parameter yang diberikan
		if student.ID == id && student.Name == name {
			// Jika cocok, kembalikan pesan login berhasil dan tidak ada error
			return fmt.Sprintf("Login berhasil: %s", name), nil
		}
	}
	// Jika tidak cocok, kembalikan pesan kosong dan error login gagal
	return "", fmt.Errorf("Login gagal: ID atau nama tidak sesuai")
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// Periksa apakah id, name, atau studyProgram kosong
	if id == "" || name == "" || studyProgram == "" {
		// Jika salah satu kosong, kembalikan pesan error
		errorMsg := "ID, Name or StudyProgram is undefined!"
		return "", fmt.Errorf("%s", errorMsg)
	}

	// Periksa apakah id sudah digunakan
	for _, student := range sm.students {
		if student.ID == id {
			errorMsg := "Registrasi gagal: id sudah digunakan"
			return "", fmt.Errorf("%s", errorMsg)
		}
	}

	// Periksa apakah studyProgram tersedia
	_, ok := sm.studentStudyPrograms[studyProgram]
	if !ok {
		errorMsg := "Study program " + studyProgram + " is not found"
		return "", fmt.Errorf(errorMsg)
	}

	// Buat instance baru dari Student dengan data yang diberikan
	newStudent := model.Student{
		ID:           id,
		Name:         name,
		StudyProgram: studyProgram,
	}
	// Tambahkan mahasiswa baru ke slice students
	sm.students = append(sm.students, newStudent)
	// Kembalikan pesan registrasi berhasil dan tidak ada error
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	// Cari nama program studi berdasarkan kode
	studyProgram, ok := sm.studentStudyPrograms[code]
	// Periksa apakah kode program studi ditemukan
	if !ok {
		// Jika tidak ditemukan, kembalikan pesan error
		errorMsg := "Program studi dengan kode " + code + " tidak ditemukan"
		return "", errors.New(errorMsg)
	}
	// Jika ditemukan, kembalikan nama program studi dan tidak ada error
	return studyProgram, nil
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	// Inisialisasi variabel untuk menandai apakah mahasiswa ditemukan
	var found bool
	// Iterasi melalui slice students
	for idx, student := range sm.students {
		// Periksa apakah nama mahasiswa cocok dengan parameter yang diberikan
		if student.Name == name {
			// Jika nama mahasiswa cocok, tandai bahwa mahasiswa ditemukan
			found = true
			// Panggil fungsi fn untuk memodifikasi data mahasiswa
			if err := fn(&sm.students[idx]); err != nil {
				return "", err
			}
			// Hentikan iterasi setelah modifikasi dilakukan
			break
		}
	}
	// Periksa apakah mahasiswa ditemukan
	if !found {
		// Jika tidak ditemukan, kembalikan pesan error
		errorMsg := "Mahasiswa dengan nama " + name + " tidak ditemukan"
		return "", fmt.Errorf(errorMsg)
	}
	// Jika ditemukan, kembalikan pesan berhasil dan tidak ada error
	return "Program studi mahasiswa berhasil diubah.", nil
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	// Kembalikan sebuah fungsi yang menerima pointer ke Student dan mengembalikan error
	return func(s *model.Student) error {
		// Periksa apakah program studi yang diminta tersedia
		_, ok := sm.studentStudyPrograms[programStudi]
		if !ok {
			// Jika tidak tersedia, kembalikan pesan error
			errorMsg := "Program studi " + programStudi + " tidak ditemukan"
			return fmt.Errorf(errorMsg)
		}
		// Jika tersedia, ubah program studi mahasiswa sesuai dengan yang diminta
		s.StudyProgram = programStudi
		// Kembalikan nil karena tidak ada error
		return nil
	}
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
