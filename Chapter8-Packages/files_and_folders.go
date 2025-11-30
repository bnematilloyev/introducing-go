package main

func main() {
	// Input / Output
	// Reader / Writer

	// func Copy(dst Writer, src Reader) (written int64, err  error)

	// Buffer nima?
	// Bu bir vaqtning o'zida: Reader va Writer

	/*
		Buffer afzalliklari:
			- avtomatik kengayadi
			- o‘z ichiga byte yozadi
			- undan byte o‘qish mumkin
			- hech qanday init qilish shart emas
	*/

	/*
		var buf bytes.Buffer
		buf.Write([]byte("test"))
		fmt.Println(buf.Bytes())
		// [116 101 115 116]
	*/

	/*
		src := strings.NewReader("Botir ishni boshladi!")
		var dst bytes.Buffer

		io.Copy(&dst, src)

		fmt.Println(dst.String())
		// Botir ishni boshladi!
	*/

	// Read the contents of a file and display them on the terminal
	/*
		file, err := os.Open("test.txt")
		if err != nil {
			fmt.Println("File doesn't exist !!!")
			return
		}
		defer file.Close()

		// get the file size
		stat, err := file.Stat()
		if err != nil {
			return
		}
		// read the file
		bs := make([]byte, stat.Size())
		_, err = file.Read(bs)
		if err != nil {
			return
		}

		str := string(bs)
		fmt.Println(str)
	*/

	// Another way
	/*
		bs, err := ioutil.ReadFile("test.txt")
		if err != nil {
			return
		}
		str := string(bs)
		fmt.Println(str)
	*/

	// Create a file
	/*
		file, err := os.Create("test.txt")
		if err != nil {
			fmt.Println("File not created")
			return
		}
		defer file.Close()

		file.WriteString("I am King")
	*/

	// To get the contents of directory
	/*
		dir, err := os.Open(".")
		if err != nil {
			return
		}

		defer dir.Close()

		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			return
		}
		for _, fi := range fileInfos {
			fmt.Println(fi.Name())
		}
		// Result
		//files_and_folders.go
		//strings.go
		//test.txt
	*/

	// To print all recursively even files inside folders
	// use filepath.Walk()
	/*
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
	*/
}
