package main

import (
	"log"
	"strconv"
)

func floatToString(f *float64) *string{
	if f==nil{
		return nil
	}
	str:=strconv.FormatFloat(*f,'f',-1,64)
	return &str
}

func main() {
	var consump float64 = 41.3126548
	log.Println(*floatToString(&consump))
	log.Println(strconv.FormatInt(98764987644656,32))
}
