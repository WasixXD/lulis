package main

import (
	"fmt"
	"log"
	"lulis/cpf"
	"os"
	"runtime/pprof"
	"time"
)

var sink int

func main() {
	f, _ := os.Create("cpu.prof")

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	sink = cpf.CalcRange(0, 999_999_999)
	fmt.Println(sink)
	fmt.Println("CalcRange Finished in: ", time.Since(now))

	// now = time.Now()
	// sink = cpf.DebugCPF(0, 999_999_999)
	// fmt.Println(sink)
	// fmt.Println("DebugCPF Finished in: ", time.Since(now))

	pprof.StopCPUProfile()
	f.Close()
}
