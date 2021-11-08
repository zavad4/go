package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	lab2 "github.com/zavad4/go/tree/main/lab2"
)

func main() {
	var inConsole = flag.String("e", "", "input from console - compulsory input string")
	var inFile = flag.String("f", "", "input from file - compulsory input file")
	var outFile = flag.String("o", "", "output to file - compulsory output file")
	var err error
	var handler lab2.ComputeHandler

	flag.Parse()
	if *inConsole == "" {
		if *inFile != "" {
			handler = lab2.ComputeHandler{
				Input:  bytes.NewBufferString(*inConsole),
				Output: os.Stdout,
			}
			if *outFile != "" {
				outfile, _ := os.Open(*outFile)
				handler = lab2.ComputeHandler{
					Input:  bytes.NewBufferString(*inConsole),
					Output: outfile,
				}
			}
		} else {
			//err = "Error: Wrong input format."
			//fmt.Fprintf(os.StdErr, "Error: Wrong input format.")
		}
	} else {
		if *inFile == "" {
			handler = lab2.ComputeHandler{
				Input:  bytes.NewBufferString(*inConsole),
				Output: os.Stdout,
			}
			if *outFile != "" {
				outfile, _ := os.Open(*outFile)
				handler = lab2.ComputeHandler{
					Input:  bytes.NewBufferString(*inConsole),
					Output: outfile,
				}
			}
		} else {
			//err = "Error: Wrong input format."
			//fmt.Fprintf(os.StdErr, "Error: Wrong input format.")
		}
	}
	err = handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors: %s", err)
	} else {
		fmt.Fprintf(os.Stderr, "Errors: wrong arguments")
	}
	// if *fFlag != "" && *eFlag == "" && *oFlag == "" {
	// 	isCorrectArgs = true
	// 	readfile, _ := os.Open(*fFlag)
	// 	handler = lab2.ComputeHandler{
	// 		In:  readfile,
	// 		Out: os.Stdout,
	// 	}
	// }
	// if *fFlag == "" && *eFlag != "" && *oFlag == "" {
	// 	isCorrectArgs = true
	// 	handler = lab2.ComputeHandler{
	// 		In:  bytes.NewBufferString(*eFlag),
	// 		Out: os.Stdout,
	// 	}
	// }
	// if *fFlag != "" && *eFlag == "" && *oFlag != "" {
	// 	isCorrectArgs = true
	// 	readfile, _ := os.Open(*fFlag)
	// 	writefile, _ := os.Open(*oFlag)
	// 	handler = lab2.ComputeHandler{
	// 		In:  readfile,
	// 		Out: writefile,
	// 	}
	// }
	// if *fFlag == "" && *eFlag != "" && *oFlag != "" {
	// 	isCorrectArgs = true
	// 	writefile, _ := os.Open(*oFlag)
	// 	handler = lab2.ComputeHandler{
	// 		In:  bytes.NewBufferString(*eFlag),
	// 		Out: writefile,
	// 	}
	// }
	// if isCorrectArgs == true {
	// 	err = handler.Compute()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Errors: %s", err)
	// 	}
	// } else {
	// 	fmt.Fprintf(os.Stderr, "Errors: wrong arguments")
	// }
}
