package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ttacon/libphonenumber"
)

type nresult struct { // here we'll fill in the results
	Nnumber string `json:"number"`
	Ntype   string `json:"type"`
}

func getnumber(number *gin.Context) {

	num, _ := number.GetQuery("norm")

	num2, err := libphonenumber.Parse(num, "DE")
	if err != nil {
		log.Fatalln(time.Now(), err)
	}

	var returntype string

	formattedNum := libphonenumber.Format(num2, libphonenumber.INTERNATIONAL)

	i := libphonenumber.GetNumberType(num2)

	switch i {
	case libphonenumber.FIXED_LINE:
		returntype = "Fixed Line"
	case libphonenumber.MOBILE:
		returntype = "Mobile"
	case libphonenumber.FIXED_LINE_OR_MOBILE:
		returntype = "Fixed Line or Mobile"
	case libphonenumber.TOLL_FREE:
		returntype = "Toll Free"
	case 4:
		returntype = "Premium Rate"
	case 5:
		returntype = "shared cost"
	case 6:
		returntype = "VoIP"
	case 7:
		returntype = "Personal Number"
	case 8:
		returntype = "Pager"
	case 9:
		returntype = "UAN"
	case 10:
		returntype = "Voice Mail"
	case 11:
		returntype = "Unknown"
	}

	fmt.Println(formattedNum)
	fmt.Println(returntype)

	result := nresult{formattedNum, returntype}
	number.IndentedJSON(200, result)

}

func main() {

	router := gin.Default()
	// router.LoadHTMLGlob("*.html")

	//	router.GET("/", usage)

	router.GET("/getnumber/", getnumber)
	log.Println(time.Now(), "Number Lookup started")
	router.Run()
}
