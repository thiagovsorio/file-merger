package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define source and destination filenames
	sourceFileOne := flag.String("source1", "", "Path to the source file")
	sourceFileTwo := flag.String("source2", "", "Path to the source file")
	destinationFile := flag.String("destination", "output.txt", "Path to the destination file (output.txt by default)")

	flag.Parse()

	if *sourceFileOne == "" || *sourceFileTwo == "" || *destinationFile == "" {
		flag.PrintDefaults()
		return
	}

	// Open the two source files
	source, err := os.Open(*sourceFileOne)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	source2, err := os.Open(*sourceFileTwo)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}

	defer source.Close()
	defer source2.Close()

	destination, err := os.Create(*destinationFile)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}

	defer destination.Close()

	scanner1 := bufio.NewScanner(source)
	scanner2 := bufio.NewScanner(source2)

	writer := bufio.NewWriter(destination)

  for {
    // Read line from first source
    if !scanner1.Scan() {
      break
    }
    line := scanner1.Text()
    fmt.Fprintln(writer, line)

    // Read line from second source
    if !scanner2.Scan() {
      break
    }
    line = scanner2.Text()
    fmt.Fprintln(writer, line)
  }
  
  if err := scanner1.Err(); err != nil {
    fmt.Println("Error reading first source file:", err)
  }
  if err := scanner2.Err(); err != nil {
    fmt.Println("Error reading second source file:", err)
  }

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error writing to destination file:", err)
	}
	fmt.Println("Successfully copied lines to", *destinationFile)
}
