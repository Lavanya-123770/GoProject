package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	inputFiles := []string{"server1.log", "server2.log", "server3.log"}
	err := ProcessLogs(inputFiles, "errors.log")
	if err != nil {
		log.Fatal("Error Detected :", err)

	}
}

func ProcessLogs(input []string, output string) error {
	errorchannel := make(chan string)
	var wg sync.WaitGroup
	var Writewg sync.WaitGroup
	errorfiles, err := os.Create(output)
	if err != nil {
		fmt.Println("Errors Occured :", err)
		return err
	}

	defer errorfiles.Close()
	Writewg.Add(1)
	go func() {
		defer Writewg.Done()
		for lineinfiles := range errorchannel {
			_, err := errorfiles.WriteString(lineinfiles + "\n")
			if err != nil {
				log.Print("Error to write the files:", err)
			}
		}
	}()
	for _, fileName := range input {
		wg.Add(1)

		go func(filname string) {
			defer wg.Done()
			file, err := os.Open(filname)
			if err != nil {
				log.Println("Error was found opening file :", filname, err)
				return

			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "ERROR") {
					formattedLines := fmt.Sprintf("[%s]%s", filname, line)
					fmt.Println("Sending error:", formattedLines)
					errorchannel <- formattedLines
				}
			}
			if err := scanner.Err(); err != nil {
				log.Println("Scanner Error ", filname, ":", err)
			}
			fmt.Println("Finished file : ", filname)
		}(fileName)
	}
	go func() {
		wg.Wait()
		close(errorchannel)
	}()
	Writewg.Wait()

	return nil
}
