package main

/*
func main() {
		//crc32:
		//	- ma’lumotning ichidan bitlar bo‘yicha hisoblab
		//	- 32-bitlik raqam qaytaradi
		//Agar ma’lumot bitta bitga ham o‘zgarsa → CRC ham o‘zgaradi.
		//Bu faqat xatoni aniqlash uchun.


	// create a hasher -> crc32
		h := crc32.NewIEEE()
		// write our data to it
		h.Write([]byte("test"))
		// calculate the crc32 checksum
		v := h.Sum32()
		fmt.Println(v)
}
*/

/*
func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// create a hasher
	h := crc32.NewIEEE()
	_, trr := io.Copy(h, f)

	if trr != nil {
		return 0, trr
	}
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		fmt.Println("1st files not exist")
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		fmt.Println("2nd files not exist")
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
*/

// SHA-1

/*
func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
	// [169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
}
*/
