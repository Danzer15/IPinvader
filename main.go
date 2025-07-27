package main

import (
	"IPinvader/apiLoc"
	"IPinvader/domain"
	"IPinvader/models"
	"fmt"
)

func main() {
	var ip string
	fmt.Println("                --- IPinvader ---                 ")
	fmt.Println("            --- IP To Domain Tools ---            ")
	fmt.Print("IP Target = ")
	fmt.Scanln(&ip)

	Sdomain := domain.GetDomain(ip)
	Slocation := apiLoc.GetLocation(ip)

	result := models.Result{
		IP:       ip,
		Domain:   Sdomain,
		Scountry: Slocation.Scountry,
		Scity:    Slocation.Scity,
		Sisp:     Slocation.Sisp,
	}

	fmt.Println("--- Result ---")
	fmt.Println("IP:  ", result.IP)
	fmt.Println("Domain:  ", result.Domain)
	fmt.Println("Country:  ", result.Scountry)
	fmt.Println("City:  ", result.Scity)
	fmt.Println("ISP:  ", result.Sisp)
}
