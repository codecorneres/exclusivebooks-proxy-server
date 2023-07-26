package routes

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
)

type NewMember struct {
	XMLName       xml.Name `xml:"NewMember"`
	CustomerId    string   `xml:"CustomerId"`
	Title         string   `xml:"Title"`
	FirstName     string   `xml:"FirstName"`
	LastName      string   `xml:"LastName"`
	ContactNumber string   `xml:"ContactNumber"`
	DateOfBirth   string   `xml:"DateOfBirth"`
	OptIn         string   `xml:"OptIn"`
	CommsPref     string   `xml:"CommsPref"`
	VodacomID     string   `xml:"vodacomID"`
	MemberIdNum   string   `xml:"MemberIdNum"`
	EmailAddress  string   `xml:"EmailAddress"`
}

type JoinFanaticsRequest struct {
	XMLName   xml.Name `xml:"JoinFanaticsRequest"`
	Namespace string   `xml:"xmlns:tns,attr"`
	Message   string   `xml:"xmlns:soap,attr"`
	Parts     JoinFanaticsParts
}

type JoinFanaticsParts struct {
	XMLName   xml.Name `xml:"parts"`
	NewMember NewMember
}

type MergeFanaticsCustomerRequest struct {
	XMLName   xml.Name `xml:"MergeFanaticsCustomer"`
	Namespace string   `xml:"xmlns:tns,attr"`
	Message   string   `xml:"xmlns:soap,attr"`
	Parts     MergeFanaticsCustomerParts
}

type MergeFanaticsCustomerParts struct {
	XMLName        xml.Name `xml:"parts"`
	CustomerId     string   `xml:"CustomerId"`
	FanaticsNumber string   `xml:"FanaticsNumber"`
}

// Define your route handlers here
func JoinFanaticsHandler(c *fiber.Ctx) error {
	const (
		soapURL    = "http://herakles.exclusivebooks.co.za/exclusive/CustomerManagement.php"
		soapAction = "urn:WebsiteCustomer#JoinFanatics"
	)
	// Build the SOAP request payload
	requestPayload := JoinFanaticsRequest{
		Namespace: "urn:WebsiteCustomer",
		Message:   "JoinFanaticsRequest",
		Parts: JoinFanaticsParts{
			NewMember: NewMember{
				CustomerId:    "123456",
				Title:         "Mr",
				FirstName:     "John",
				LastName:      "Doe",
				ContactNumber: "1234567890",
				DateOfBirth:   "1990-01-01",
				OptIn:         "true",
				CommsPref:     "email",
				VodacomID:     "V123456",
				MemberIdNum:   "M123456",
				EmailAddress:  "john.doe@example.com",
			},
		},
	}

	// Convert the SOAP request to XML
	requestBody, err := xml.MarshalIndent(requestPayload, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling SOAP request:", err)
		return c.SendString("Error marshaling SOAP request:")
	}

	// Create the HTTP POST request
	fmt.Println("Requeest params:")
	fmt.Println(string(requestBody))
	req, err := http.NewRequest("POST", soapURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return c.SendString("Error creating HTTP request:")
	}

	// Set SOAP-specific headers
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", soapAction)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending SOAP request:", err)
		return c.SendString("Error sending SOAP request:")
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading SOAP response:", err)
		return c.SendString("Error reading SOAP response:")
	}

	// Print the SOAP response
	fmt.Println("Successfull SOAP response:", resp)
	fmt.Println(string(responseBody))
	return c.SendString(string(responseBody))
}

func MergeFanaticsCustomerHandler(c *fiber.Ctx) error {
	const (
		soapURL    = "http://herakles.exclusivebooks.co.za/exclusive/CustomerManagement.php"
		soapAction = "urn:WebsiteCustomer#MergeFanaticsCustomer"
	)
	// Build the SOAP request payload
	requestPayload := MergeFanaticsCustomerRequest{
		Namespace: "urn:WebsiteCustomer",
		Message:   "MergeFanaticsCustomerRequest",
		Parts: MergeFanaticsCustomerParts{
			CustomerId:     "123456",
			FanaticsNumber: "313122",
		},
	}

	// Convert the SOAP request to XML
	requestBody, err := xml.MarshalIndent(requestPayload, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling SOAP request:", err)
		return c.SendString("Error marshaling SOAP request:")
	}

	// Create the HTTP POST request
	fmt.Println("Requeest params:")
	fmt.Println(string(requestBody))
	req, err := http.NewRequest("POST", soapURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return c.SendString("Error creating HTTP request:")
	}

	// Set SOAP-specific headers
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", soapAction)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending SOAP request:", err)
		return c.SendString("Error sending SOAP request:")
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading SOAP response:", err)
		return c.SendString("Error reading SOAP response:")
	}

	// Print the SOAP response
	fmt.Println("Successfull SOAP response:", resp)
	fmt.Println(string(responseBody))
	return c.SendString(string(responseBody))
}
