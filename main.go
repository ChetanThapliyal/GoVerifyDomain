package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Initialize a scanner to read input from the user
	scanner := bufio.NewScanner(os.Stdin)

	// Loop to repeatedly ask for input until "exit" is entered
	for {
		fmt.Print("Enter a domain name to verify (or type 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		// Exit the loop if user enters "exit"
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Call the function to check the domain/email
		checkDomain(input)
	}

	// Check for any errors while reading input
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Lookup MX records for the domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
	}

	// Check if MX records were found
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Lookup TXT records for the domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	}

	// Loop through TXT records to find SPF record
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Lookup TXT records for DMARC
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
	}

	// Loop through DMARC records to find DMARC record
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Print the results
	fmt.Println(domain)
	fmt.Printf("hasMX : %v\n", hasMX)
	fmt.Printf("hasSPF : %v\n", hasSPF)
	fmt.Printf("spfRecord : %s\n", spfRecord)
	fmt.Printf("hasDMARC : %v\n", hasDMARC)
	fmt.Printf("dmarcRecord : %s\n", dmarcRecord)
	fmt.Println()
}
