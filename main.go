// main.go
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	// Metode untuk manajemen mahasiswa
	Login(id string, name string) (string, error)                         // Login mahasiswa
	Register(id string, name string, studyProgram string) (string, error) // Mendaftarkan mahasiswa baru
	GetStudyProgram(code string) (string, error)                          // Mendapatkan nama program studi
	ModifyStudent(name string, fn model.StudentModifier) (string, error)  // Memodifikasi data mahasiswa
	GetStudents() []model.Student                                         // Mendapatkan daftar mahasiswa
	ImportStudents(filenames []string) error                              // Mengimpor data mahasiswa
	SubmitAssignments(numAssignments int)                                 // Mengumpulkan tugas
}

type InMemoryStudentManager struct {
	sync.Mutex                             // Sinkronisasi akses ke data mahasiswa.
	students             []model.Student   // Daftar mahasiswa.
	studentStudyPrograms map[string]string // Pemetaan ID mahasiswa ke program studi.
	failedLoginAttempts  map[string]int    // Pemetaan ID mahasiswa ke jumlah upaya login yang gagal.
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		// Inisialisasi daftar mahasiswa dengan data yang sudah ditentukan.
		students: []model.Student{
			{ID: "A12345", Name: "Aditira", StudyProgram: "TI"},
			{ID: "B21313", Name: "Dito", StudyProgram: "TK"},
			{ID: "A34555", Name: "Afis", StudyProgram: "MI"},
		},
		// Inisialisasi pemetaan program studi ke nama program studi.
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		// Inisialisasi jumlah upaya login yang gagal.
		failedLoginAttempts: make(map[string]int),
	}
}

// Membaca data mahasiswa dari file CSV yang diberikan.
func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	// Membuka file CSV dengan nama file yang diberikan.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Menggunakan defer untuk memastikan file ditutup setelah selesai dibaca.

	// Membuat pembaca CSV dari file yang dibuka.
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name dan StudyProgram

	var students []model.Student
	// Membaca setiap baris dalam file CSV.
	for {
		// Membaca satu baris dari file CSV.
		record, err := reader.Read()
		// Menghentikan perulangan jika sudah mencapai akhir file.
		if err == io.EOF {
			break
		}
		// Mengembalikan error jika terjadi kesalahan saat membaca file.
		if err != nil {
			return nil, err
		}

		// Membuat objek model.Student dari data yang dibaca dan menambahkannya ke slice students.
		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	// Mengembalikan slice dari model.Student yang berisi data mahasiswa yang dibaca dari file CSV.
	return students, nil
}

// GetStudents mengembalikan daftar mahasiswa dari InMemoryStudentManager.
func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	// Mengembalikan slice dari mahasiswa yang disimpan di InMemoryStudentManager.
	return sm.students
}

// Login melakukan proses otentikasi untuk mahasiswa dengan ID dan nama yang diberikan.
func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	// Mengunci akses untuk operasi concurent.
	sm.Lock()
	defer sm.Unlock()

	// Memeriksa apakah jumlah upaya login yang gagal sudah mencapai batas maksimum.
	if sm.failedLoginAttempts[id] >= 3 {
		return "", fmt.Errorf("Login gagal: Batas maksimum login terlampaui")
	}

	// Memeriksa setiap mahasiswa dalam daftar untuk mencocokkan ID dan nama.
	for _, student := range sm.students {
		if student.ID == id && student.Name == name {
			// Reset jumlah upaya login yang gagal jika login berhasil.
			sm.failedLoginAttempts[id] = 0
			// Mendapatkan nama program studi mahasiswa yang berhasil login.
			studyProgram, err := sm.GetStudyProgram(student.StudyProgram)
			if err != nil {
				return "", err
			}
			// Mengembalikan pesan sukses login beserta nama dan program studi mahasiswa.
			return fmt.Sprintf("Login berhasil: Selamat datang %s! Kamu terdaftar di program studi: %s", name, studyProgram), nil
		}
	}

	// Menambahkan jumlah upaya login yang gagal jika data mahasiswa tidak ditemukan.
	sm.failedLoginAttempts[id]++
	return "", fmt.Errorf("Login gagal: data mahasiswa tidak ditemukan")
}

// Register mendaftarkan mahasiswa baru dengan ID, nama, dan program studi yang diberikan.
func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// Memeriksa jika salah satu dari field kosong
	if id == "" || name == "" || studyProgram == "" {
		return "", fmt.Errorf("%s", "ID, Name or StudyProgram is undefined!")
	}

	// Memeriksa apakah program studi yang diberikan valid
	if _, err := sm.GetStudyProgram(studyProgram); err != nil {
		return "", fmt.Errorf("Study program " + studyProgram + " is not found")
	}

	// Mengunci akses untuk operasi concurent.
	sm.Lock()
	defer sm.Unlock()

	// Memeriksa jika ID mahasiswa sudah digunakan sebelumnya
	for _, student := range sm.students {
		if student.ID == id {
			return "", fmt.Errorf("%s", "Registrasi gagal: id sudah digunakan")
		}
	}

	// Menambahkan mahasiswa baru ke dalam daftar mahasiswa
	sm.students = append(sm.students, model.Student{ID: id, Name: name, StudyProgram: studyProgram})
	// Mengembalikan pesan sukses registrasi
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil
}

// GetStudyProgram mendapatkan nama program studi berdasarkan kode program studi yang diberikan.
func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	// Memeriksa apakah kode program studi ditemukan dalam pemetaan program studi.
	program, ok := sm.studentStudyPrograms[code]
	if !ok {
		// Mengembalikan error jika kode program studi tidak ditemukan.
		return "", fmt.Errorf("Program studi dengan kode" + code + "tidak ditemukan")
	}
	// Mengembalikan nama program studi yang sesuai dengan kode program studi yang diberikan.
	return program, nil
}

// ModifyStudent memodifikasi data mahasiswa dengan nama yang diberikan menggunakan fungsi yang ditentukan.
func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	// Mengunci akses untuk operasi concurent.
	sm.Lock()
	defer sm.Unlock()

	// Memeriksa setiap mahasiswa dalam daftar untuk mencari mahasiswa dengan nama yang sesuai.
	for i, student := range sm.students {
		if student.Name == name {
			// Memanggil fungsi yang ditentukan untuk memodifikasi data mahasiswa.
			if err := fn(&sm.students[i]); err != nil {
				return "", err
			}
			// Mengembalikan pesan sukses setelah berhasil memodifikasi data mahasiswa.
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}
	// Mengembalikan error jika data mahasiswa dengan nama yang diberikan tidak ditemukan.
	return "", fmt.Errorf("Data mahasiswa" + name + "tidak ditemukan")
}

// ChangeStudyProgram mengembalikan sebuah fungsi model.StudentModifier yang mengubah program studi mahasiswa menjadi program studi yang diberikan.
func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	// Mengembalikan fungsi model.StudentModifier.
	return func(s *model.Student) error {
		// Mengubah program studi mahasiswa menjadi program studi yang diberikan.
		s.StudyProgram = programStudi
		return nil
	}
}

// ImportStudents mengimpor data mahasiswa dari file-file yang diberikan.
func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {
	// Membuat sebuah channel buffered untuk mengirimkan rekaman mahasiswa.
	studentsCh := make(chan model.Student, 100000)

	var wg sync.WaitGroup

	// Memulai goroutine worker
	for i := 0; i < len(filenames); i++ {
		wg.Add(1)
		go func(file string) {
			defer wg.Done() // Mengurangi WaitGroup setelah proses file selesai

			time.Sleep(40 * time.Millisecond) // Memberi jeda untuk simulasi
			students, err := ReadStudentsFromCSV(file)
			if err != nil {
				fmt.Printf("Error importing from %s: %s\n", file, err.Error())
				return
			}
			for _, student := range students {
				studentsCh <- student
			}
		}(filenames[i])
	}

	// Menutup channel setelah semua mahasiswa dikirim
	go func() {
		wg.Wait()
		close(studentsCh)
	}()

	// Memulai goroutine terpisah untuk memproses pendaftaran mahasiswa secara konkuren
	var regWg sync.WaitGroup
	for i := 0; i < 10; i++ { // Jumlah goroutine pekerja
		regWg.Add(1)
		go func() {
			defer regWg.Done()
			for student := range studentsCh {
				_, err := sm.Register(student.ID, student.Name, student.StudyProgram)
				if err != nil {
					fmt.Printf("Error registering student %s: %s\n", student.Name, err.Error())
				}
			}
		}()
	}

	// Menunggu semua goroutine pendaftaran selesai
	regWg.Wait()

	return nil
}

// SubmitAssignmentLongProcess adalah fungsi yang mensimulasikan proses pengumpulan tugas yang memakan waktu.
func SubmitAssignmentLongProcess() {
	// Memberi jeda eksekusi selama 30 milidetik untuk mensimulasikan proses yang memakan waktu.
	time.Sleep(30 * time.Millisecond)
}

// SubmitAssignments menyerahkan tugas sejumlah numAssignments untuk diproses.
func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {
	// Menggunakan channel buffered untuk membatasi jumlah tugas yang diproses secara konkuren.
	jobs := make(chan int, numAssignments)

	// Memulai goroutine worker untuk memproses tugas
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for job := range jobs {
				// Memanggil fungsi SubmitAssignmentLongProcess untuk memproses tugas.
				SubmitAssignmentLongProcess()
				_ = job
			}
		}(w)
	}

	// Mengirimkan tugas ke dalam channel untuk diproses
	for i := 1; i <= numAssignments; i++ {
		jobs <- i
	}

	// Menutup channel setelah semua tugas terkirim
	close(jobs)

	// Menunggu semua goroutine worker selesai
	wg.Wait()
}

// main adalah fungsi utama yang menjalankan aplikasi manajemen mahasiswa.
func main() {
	// Membuat instance dari InMemoryStudentManager.
	manager := NewInMemoryStudentManager()

	for {
		// Membersihkan layar konsol.
		helper.ClearScreen()

		// Mendapatkan daftar mahasiswa dari manajer mahasiswa.
		students := manager.GetStudents()
		for _, student := range students {
			// Menampilkan informasi mahasiswa.
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println("-----------")
		}

		// Menampilkan menu pilihan.
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Modify Student")
		fmt.Println("4. Import Students")
		fmt.Println("5. Submit Assignments")
		fmt.Println("6. Exit")

		var choice int
		// Meminta pengguna untuk memilih menu.
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Menu untuk login.
			var id, name string
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Name: ")
			fmt.Scanln(&name)
			// Memanggil fungsi Login pada manajer mahasiswa.
			message, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println(message)
			}
			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case 2:
			// Menu untuk registrasi mahasiswa baru.
			var id, name, studyProgram string
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Study Program: ")
			fmt.Scanln(&studyProgram)
			// Memanggil fungsi Register pada manajer mahasiswa.
			message, err := manager.Register(id, name, studyProgram)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println(message)
			}
			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case 3:
			// Menu untuk memodifikasi data mahasiswa.
			var name, newProgram string
			fmt.Print("Enter Name of Student: ")
			fmt.Scanln(&name)
			fmt.Print("Enter New Study Program: ")
			fmt.Scanln(&newProgram)
			// Memanggil fungsi ModifyStudent pada manajer mahasiswa.
			message, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(newProgram))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println(message)
			}
			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case 4:
			// Menu untuk mengimpor data mahasiswa dari file.
			var filenames string
			fmt.Print("Enter filenames separated by comma: ")
			fmt.Scanln(&filenames)
			files := strings.Split(filenames, ",")
			// Memanggil fungsi ImportStudents pada manajer mahasiswa.
			err := manager.ImportStudents(files)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Import successful")
			}
			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case 5:
			// Menu untuk mengumpulkan tugas.
			var numAssignmentsStr string
			fmt.Print("Enter number of assignments: ")
			fmt.Scanln(&numAssignmentsStr)
			numAssignments, err := strconv.Atoi(numAssignmentsStr)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				break
			}
			// Memanggil fungsi SubmitAssignments pada manajer mahasiswa.
			manager.SubmitAssignments(numAssignments)
			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		case 6:
			// Keluar dari aplikasi.
			os.Exit(0)
		default:
			// Pesan jika pilihan tidak valid.
			fmt.Println("Invalid choice")
		}
	}
}
