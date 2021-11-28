package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

//OpenInputFiles opens the input file
func OpenInputFile(fileName string) []string{
	Seq, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: Cannot open the input files")
		os.Exit(3)
	}

	// create the variable to hold the lines
	var lines []string = make([]string, 0)

	scanner := bufio.NewScanner(Seq)

	for scanner.Scan() {
		// append it to the lines slice
		lines = append(lines, scanner.Text())
		//fmt.Println("The current line is:", scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(3)
	}

	// close the file and return the lines
	Seq.Close()
	return lines
}

func GetMapAccessionID(fileData [] string) []string{
	var AccessID []string = make([]string, 0)

	//fmt.Println(len(fileData))

	for i:=0;i<len(fileData);i++{
	    StartTag := strings.Contains(fileData[i], ">")
	    if StartTag == true {
	    	//fmt.Println(fileData[i])
	    	var Data []string = strings.Split(fileData[i], " ")
	    	//fmt.Println(Data)
	    	//fmt.Println(Data[0])
   			AccessID = append(AccessID, Data[0])
	    }
	}
	return AccessID
}

func GetNucleotideSequence(fileData [] string) []string{
	var NucleotideSequence []string = make([]string, 0)
	var SingleNucleotideSequence string

	for i:=0;i<len(fileData);i++{
	    StartTag := strings.Contains(fileData[i], ">")
	    if StartTag == false {
	    	if fileData[i] == "" {
				NucleotideSequence = append(NucleotideSequence,SingleNucleotideSequence)
				SingleNucleotideSequence = ""
	    	} else {
	    		SingleNucleotideSequence = SingleNucleotideSequence + fileData[i]
	    	}
	    } 
	}
	NucleotideSequence = append(NucleotideSequence,SingleNucleotideSequence)
	SingleNucleotideSequence = ""

	fmt.Println(len(NucleotideSequence))
	return NucleotideSequence
}

func Mapper(ListOfAccessionID,NucleotideSequence [] string) map[string]string{

	MappedIdToNucleotideSeq := make(map[string]string)

	AccessionIDLen := len(ListOfAccessionID)
	NucleotideSequenceLen := len(NucleotideSequence)

	if AccessionIDLen != NucleotideSequenceLen {
		fmt.Println("Error: there is some kind of error in the file data")
		os.Exit(3)
	} else {
		for i:=0;i<NucleotideSequenceLen;i++{
			ID := ListOfAccessionID[i]
			ID = ID[1:len(ID)]			
			MappedIdToNucleotideSeq[ID] = NucleotideSequence[i]
		}
	}

	return MappedIdToNucleotideSeq
}

func main() {

	filename := os.Args[1]               //Read the user given filename
	FileData := OpenInputFile(filename) //Open the file with given filename and return Data
	//fmt.Println(openFile)
	ListOfAccessionID := GetMapAccessionID(FileData)
	//fmt.Println(ListOfAccessionID)
	NucleotideSequence := GetNucleotideSequence(FileData)
	//fmt.Println(NucleotideSequence)
	fmt.Println(Mapper(ListOfAccessionID,NucleotideSequence))
}