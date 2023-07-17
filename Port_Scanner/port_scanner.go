//SHEO, 09.07.2023, TCP-PORT-SCANNER, GOLANG

package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

// Global variable to save amount of ports ¯\_(ツ)_/¯
var totalPorts = 0

// A worker-func to run inside goroutines
func worker(
	ports chan int,
	address string,
	results chan int) {

	for p := range ports {
		address := fmt.Sprintf("%s:%d", address, p)
		conn, err := net.Dial("tcp", address)
		///Add the number of the port into the chan
		///In case the connection was successful
		if err != nil {
			results <- 0
			continue
		}
		//End the connxtion
		conn.Close()
		results <- p
	}
}

// Nothin' special: unique name generator for files
func GenName(ad string) string {
	dt := time.Now()
	var a string = dt.Format("01-02-2006")
	return ad + "_" + a + ".txt"
}

// Nothin' special either: file-writer
func WriteLog(pts []int, fname, address string) {
	file, err := os.Create(fname)
	defer file.Close()
	//useless handler, not shit-code tho
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, address+"\n")
	//loop through every port
	for i := range pts {
		fmt.Fprintf(file, "Port %d opened!\n", pts[i])
	}
	fmt.Println(fname)
}

func main() {
	fmt.Println("Input the address of the website to scan: ")
	var address string

	fmt.Scanln(&address)

	fmt.Println("Input the ports range: ")

	var lb, ub int
	fmt.Fscan(os.Stdin, &lb)
	fmt.Fscan(os.Stdin, &ub)
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	//loop to make some workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, address, results)
	}
	//Insertion of all the ports into the channels
	//Run result gathering when actual scanning begins
	go func() {
		for i := lb; i <= ub; i++ {
			ports <- i
		}
	}()

	for i := lb; i < ub; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	fmt.Printf("\nTotal ports: %d\n-----\n", len(openports))

	z := 1
	for _, port := range openports {
		fmt.Printf("%d) Port %d opened!\n", z, port)
		z++
	}
	fmt.Printf(`Ports had been scanned.
Do you want to create a log file? (Y|N) `)

	var exit string
	fmt.Scan(&exit)
	if exit == "y" || exit == "Y" {
		WriteLog(openports, GenName(address), address)
	}
	close(ports)
	close(results)
}
