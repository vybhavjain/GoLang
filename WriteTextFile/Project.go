package main

import (
	"fmt"
	"os"
	//"bufio"
	//"strings"
)

func CreateFile(FileName string) {
	// Open the file to write the frequency map
	outFile, err := os.Create(FileName)
	if err != nil {
		fmt.Println("Sorry: could not create the file")
	}
    outFile.Close()
}


func WriteDataToFile(FileName,AccessionID string, freqMap map[string]int) {
	outFile, err := os.OpenFile(FileName ,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//    f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error: Cannot open the input files")
		os.Exit(3)
	}

	fmt.Fprintf(outFile, "Data for accession ID is: %v \n",AccessionID)
		
	for sequence := range freqMap {
		fmt.Fprintf(outFile, "Sequence is: %v and Value is %d \n",sequence,freqMap[sequence])	
	}
	fmt.Fprintf(outFile, "\n \n",)	
	//fmt.Fprintln(outFile)
}



func main() {

	Value := make(map[string]int)
	Value["AAA"] = 2
	Value["CAC"] = 23
	Value["ATA"] = 21
	Value["CGA"] = 4
	fmt.Println(Value)
	AccessionID1 := "AJODBU20.6"
	AccessionID2 := "AJODBU20.6"
	CreateFile("output.txt")
	WriteDataToFile("output.txt",AccessionID1,Value)
	WriteDataToFile("output.txt",AccessionID2,Value)

}