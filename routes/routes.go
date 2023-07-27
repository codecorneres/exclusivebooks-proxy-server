package routes

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
)

func JoinFanaticsHandler(c *fiber.Ctx) error {
	postRequest := struct {
		CustomerId    string `json:"CustomerId"`
		Title         string `json:"Title"`
		FirstName     string `json:"FirstName"`
		LastName      string `json:"LastName"`
		ContactNumber string `json:"ContactNumber"`
		DateOfBirth   string `json:"DateOfBirth"`
		OptIn         string `json:"OptIn"`
		CommsPref     string `json:"CommsPref"`
		VodacomID     string `json:"VodacomID"`
		MemberIdNum   string `json:"MemberIdNum"`
		EmailAddress  string `json:"EmailAddress"`
	}{}

	if err := c.BodyParser(&postRequest); err != nil {
		return err
	}
	// URL of the SOAP endpoint
	soapURL := "http://staging.herakles.exclusivebooks.co.za/exclusive/CustomerManagement.php"

	// SOAP Action header
	soapAction := "urn:WebsiteCustomer#JoinFanatics"

	// SOAP request body (replace with your actual SOAP envelope)
	soapRequest := `
			<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:WebsiteCustomer">
				<soap:Body>
					<tns:JoinFanaticsRequest>
						<NewMember>
			                <CustomerId>` + postRequest.CustomerId + `</CustomerId>
			                <Title>` + postRequest.Title + `</Title>
			                <FirstName>` + postRequest.FirstName + `</FirstName>
			                <LastName>` + postRequest.LastName + `</LastName>
			                <ContactNumber>` + postRequest.ContactNumber + `</ContactNumber>
			                <DateOfBirth>` + postRequest.DateOfBirth + `</DateOfBirth>
			                <OptIn>` + postRequest.OptIn + `</OptIn>
			                <CommsPref>` + postRequest.CommsPref + `</CommsPref>
			                <VodacomID>` + postRequest.VodacomID + `</VodacomID>
			                <MemberIdNum>` + postRequest.MemberIdNum + `</MemberIdNum>
			                <EmailAddress>` + postRequest.EmailAddress + `</EmailAddress>
			            </NewMember>
					</tns:JoinFanaticsRequest>
				</soap:Body>
			</soap:Envelope>`

	// Create the HTTP request
	req, err := http.NewRequest("POST", soapURL, bytes.NewBufferString(soapRequest))
	if err != nil {
		fmt.Println("Error creating SOAP request:", err)
		return c.SendString("Error creating SOAP request:")
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)

	// Make the request to the SOAP endpoint
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making SOAP request:", err)
		return c.SendString("Error marshaling SOAP request:")
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading SOAP response:", err)
		return c.SendString("Error reading SOAP response:")
	}

	// Print the SOAP response
	fmt.Println("SOAP Response:")
	fmt.Println(string(body))
	return c.SendString(string(body))
}

func MergeFanaticsCustomerHandler(c *fiber.Ctx) error {
	postRequest := struct {
		CustomerId     string `json:"CustomerId"`
		FanaticsNumber string `json:"FanaticsNumber"`
	}{}

	if err := c.BodyParser(&postRequest); err != nil {
		return err
	}
	// URL of the SOAP endpoint
	soapURL := "http://staging.herakles.exclusivebooks.co.za/exclusive/CustomerManagement.php"

	// SOAP Action header
	soapAction := "urn:WebsiteCustomer#MergeFanaticsCustomer"

	// SOAP request body (replace with your actual SOAP envelope)
	soapRequest := `
			<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tns="urn:WebsiteCustomer">
				<soap:Body>
					<tns:MergeFanaticsCustomerRequest>
			                <CustomerId>` + postRequest.CustomerId + `</CustomerId>
			                <FanaticsNumber>` + postRequest.FanaticsNumber + `</FanaticsNumber>
					</tns:MergeFanaticsCustomerRequest>
				</soap:Body>
			</soap:Envelope>`

	// Create the HTTP request
	req, err := http.NewRequest("POST", soapURL, bytes.NewBufferString(soapRequest))
	if err != nil {
		fmt.Println("Error creating SOAP request:", err)
		return c.SendString("Error creating SOAP request:")
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)

	// Make the request to the SOAP endpoint
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making SOAP request:", err)
		return c.SendString("Error marshaling SOAP request:")
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading SOAP response:", err)
		return c.SendString("Error reading SOAP response:")
	}

	// Print the SOAP response
	fmt.Println("SOAP Response:")
	fmt.Println(string(body))
	return c.SendString(string(body))

}
