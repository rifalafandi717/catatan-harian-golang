package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)


func main(){
	for{
		fmt.Println("\n=== Aplikasi Catatan Harian ===")
		fmt.Println("1. Tambah catatan")
		fmt.Println("2. Lihat catatan hari ini")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih Menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan{
		case 1:
			TambahCatatan()
			fmt.Println("Tekan ENTER untuk kembali...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			ClearSceen()
		case 2:
			LihatCatatan()
			fmt.Println("Tekan ENTER untuk kembali...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			ClearSceen()
		case 3:
			fmt.Println("keluar....")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func AmbilFile() string  {
	hari := time.Now().Format("2006-01-03")
	return hari + ".txt"
}

func TambahCatatan(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Tulis catatan: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	filename := AmbilFile()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Gagal menulis pesan", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(text + "\n")
	if err != nil {
		fmt.Println("gagal menimpan catatan hari ini.")
	}else {
		fmt.Println("Catatan Disimpan")
	}
	
}

func LihatCatatan(){
	filename := AmbilFile()
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Catatan belum ada hari ini.")
		return
	}
	fmt.Println("Catatan hari ini: ")
	fmt.Println(string(data))
}

func ClearSceen(){
	if runtime.GOOS == "windows"{
		cmd := exec.Command("cmd", "/c", "cls", "clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}else {
		command := exec.Command("clear")
		command.Stdout = os.Stdout
		command.Run()
	}
}