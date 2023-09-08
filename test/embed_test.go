package test

import (
	"embed"
	_ "embed" //jika kita tidak ingin menggunakan tambahkan _
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// embed tidak bisa di dalam function
// bisa berkali kali walaupun file sama atau file berbeda

/*
Embed File ke String
1. Embed File bisa kita lakukan ke variable dengan tipe data String
2. Secara otomatiis isi file akan di baca sebagai text dan masukkan ke variable tersebut

*/
//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

/*
Embed File ke []byte
1. Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
2. Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain lain
file file binary : gambar, vidio, atau musik
*/

//go:embed ../belajar-golang-embed/logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_baru.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
Embed Multiple Files
1. Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
2. Hal ini juga bisa dilakukan menggunakan embed package
3. Kita bisa menambahkan komentar //go:embed lebih dari satu baris
4. Selain itu variablenya bisa kita gunakan tipe data embed.FS (File System)
*/

//go:embed test/files/a.txt
//go:embed test/files/b.txt
//go:embed test/files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	//a itu adalah byte slice ([]byte)
	a, _ := files.ReadFile("files/a.txt")
	//jika ingin di print, konversikan ke string terlebih dahulu
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	//jika ingin di print, konversikan ke string terlebih dahulu
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	//jika ingin di print, konversikan ke string terlebih dahulu
	fmt.Println(string(c))

}

/*
Path Matcher
1. Selain Manual satu per satu, kita bisa menggunakan patch matcher untuk membaca multiple file yang kita inginkan\
2. Inii sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
3. Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match
https://golang.org/pkg/path/#Match
*/
//kalau * saja gaperduli filenya apa akan di load
// kalau *.txt akan di ambil semua file yang bentuknya txt

//go:embed test/files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() { // ! adalah bukan
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

/*
Hasil Embed di Compile
1. Perlu diketahui, Bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file golangnya
2. Artinya bukan dilakukan secara realtime membaca file yang ada diluar
3. Hal ini menjadikan jika binary file golang sudah fi compile, kita tidak butuh lagi file externalnya, dan bahkan jika diunah file externalnya, isi variable nya tidak akan berubah lagi

Contoh ada di Folder Test, dan file main
*/
