package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	localhostTime := time.Now().Round(0)
	exactNTPTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalf("Smth goes wrong. Error: \"%v\"\n", err)
	}
	//No need to do Round() due to it's already done in Local(). t.stripMono()
	fmt.Printf("current time: %s\nexact time: %s\n", localhostTime.String(), exactNTPTime.Local().String())
}
