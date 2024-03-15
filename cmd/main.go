package main

import (
	"log"

	calc "github.com/hidori/go-div"
)

func main() {
	for dividend := 0; dividend <= 0x0ffff; dividend++ {
		if dividend%0x0100 == 0 {
			log.Printf("dividend=%02x??\n", dividend/0x0100)
		}

		for divisor := 1; divisor <= 0x0ffff; divisor++ {
			quotient, remainder, err := calc.Div(uint16(dividend), uint16(divisor))
			if err != nil {
				log.Fatalf("fail to Div(%04x/%04x) err=%v", dividend, divisor, err)
			}

			if quotient != uint16(dividend/divisor) {
				log.Fatalf("Div(%04x/%04x) returns invalid quotient=%04x", dividend, divisor, quotient)
			}

			if remainder != uint16(dividend%divisor) {
				log.Fatalf("Div(%04x/%04x) returns invalid remainder=%04x", dividend, divisor, remainder)
			}
		}
	}
}
