package main

import (
	// "fmt"
	// "bufio"
	// "fmt"
	"fmt"
	"os"
) 

func main(){
	// f, err := os.Open("example.txt")
	
	// if err != nil{
	// 	// log the error
	// 	panic(err)
	// }

	// defer f.Close()

	// finfo, err := f.Stat()
	// if err != nil{
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("file Name :",finfo.Name())
	// fmt.Println("is directory :", finfo.IsDir())
	// fmt.Println("file size :", finfo.Size())
	// fmt.Println("file Permission :", finfo.Mode())
	// fmt.Println("file modified at :", finfo.ModTime())






	// reading file Method 1
	
	// // creating buffer to store the file things that we want
	// buf := make([]byte, finfo.Size())

	// d, err := f.Read(buf)
	// if err != nil{
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++{

	// 	fmt.Println("data :", d, string(buf[i]))
	// }




	// // reading file Method 2
	// // use this method only when handling small size files
	// f, err := os.ReadFile("example.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	// // "../"  means going 1 level back
	// // "." means current folder


	// // reading folders
	// dir, err := os.Open("../") 
	// if err != nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// // n > 0    n file names returned
	// // n <= 0   return all file names

	// file_info, err := dir.ReadDir(-1)

	// for _, fi := range file_info{
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }







	// how to write in file 1


	// // creating file
	// f, err := os.Create("example2.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("hi! Hello")
	// // f.WriteString("My name is Udaykaran Singh")
	// // // f.WriteString works in append mode




	// // writing in file 2
	// bytes := []byte("hi! Hello")
	// f.Write(bytes)





	// read and write to another file (streaming fashion)



	// soc_file, err := os.Open("Example.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer soc_file.Close()

	// dest_file, err := os.Create("example2.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer dest_file.Close()

	// // using bufio (predefined buffer) to stream
	// reader := bufio.NewReader(soc_file)
	// writer := bufio.NewWriter(dest_file)

	// for{
	// 	b, err := reader.ReadByte() // read byte by byte
	// 	if err != nil{
	// 		if err.Error() != "EOF"{
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil{
	// 		panic(e)
	// 	}
	// }



	// // flush the writer maybe to free the memory
	// writer.Flush()

	// fmt.Println("written to new file sucessfully")






	// how to delete a file

	err := os.Remove("example2.txt")

	if err != nil{
		panic(err)
	}

	fmt.Println("file deleted sucessfully")


}