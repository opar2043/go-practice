package main

import ("os" 
"fmt")

func main() {
	  // create file
	  file, _ := os.Create("test.txt");

	  // write file
	  file.WriteString("Hello Opar ")
	  file.WriteString("i nam writinmg the file")

	  // read the text of the file
	  read , _ := os.ReadFile("test.txt")
	  fmt.Println(string(read))
 
	  // delete file
	 // os.Remove("test.txt")

	 file2 , _ := os.Create("new.txt")
	 file2.Close()
	 file2.WriteString("chi chi")
	 os.Remove("new.txt")
	 
	 defer file.Close();
}