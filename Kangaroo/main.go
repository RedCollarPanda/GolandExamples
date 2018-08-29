package main

import (
	"os"
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
	"strconv"
	"math"
)

type Config struct {
	Basic struct {
		MinValue int
		MaxValue  int
	}
}


func main() {

	var cfg Config
	//Лучше передавать путь до конфига, но для простоты можно и так
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]

	if len(args) != 4 {
		fmt.Println("Ожидается 4 (четыре) параметра!")
		os.Exit(1)
	}

	var confirmed []int64

	for _, value := range args {

		res, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		if res < cfg.Basic.MinValue || res > cfg.Basic.MaxValue {
			log.Fatalf("Input data %d is out of range %d <-> %d ." , res, cfg.Basic.MinValue, cfg.Basic.MaxValue)
			return
		}

		confirmed = append(confirmed, int64(res))

	}

	x1, v1, x2, v2 := confirmed[0], confirmed[1], confirmed[2], confirmed[3]

	//либо задний кенгуру догоняет переднего, либо нам тут нечего делать.
	if !isClosing(x1, v1, x2, v2)  {
		log.Println("No")
		return
	}

	//точка пересечения x1 + i * v1 == x2 + i * v2. <==> x2 - x1 == (v1 - v2)* i.
	// <==>
	if math.Mod(math.Abs(float64(x2-x1)), math.Abs(float64(v1-v2))) != 0 {
		log.Println("No")
		return
	}


	log.Println("Yes")
}

func isClosing(x1, v1, x2, v2 int64) bool {

	return (x2-x1 > 0 && v2-v1 < 0) || (x2-x1 < 0 && v2-v1 > 0)

}
