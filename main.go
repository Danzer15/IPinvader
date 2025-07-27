package main

import (
	"IPinvader/apiLoc"
	"IPinvader/domain"
	"IPinvader/models"
	"fmt"
)

func main() {
	var ip string
	fmt.Println("                               --- IPinvader ---                 ")
	fmt.Print("                       --- Check Ip that you got in here ---            \n\n")
	fmt.Print("IP Target = ")
	fmt.Scanln(&ip)
	fmt.Print("\n")

	Sdomain := domain.GetDomain(ip)
	Slocation := apiLoc.GetLocation(ip)

	result := models.Result{
		IP:       ip,
		Domain:   Sdomain,
		Scountry: Slocation.Scountry,
		Scity:    Slocation.Scity,
		Sisp:     Slocation.Sisp,
	}

	fmt.Print("--- Result ---\n\n")
	fmt.Println("IP:  ", result.IP)
	fmt.Println("Domain:  ", result.Domain)
	fmt.Println("Country:  ", result.Scountry)
	fmt.Println("City:  ", result.Scity)
	fmt.Println("ISP:  ", result.Sisp)
}
