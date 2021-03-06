package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	instnList := getIpListFromFilename(filename)
	// Part 1
	interpreter := createInterpreter(instnList)
	/*
		interpreter.varEnv["a"] = 1
		interpreter.varEnv["d"] = 2
		interpreter.varEnv["e"] = 106500
		interpreter.varEnv["g"] = 0
		interpreter.varEnv["b"] = 106500
		interpreter.varEnv["c"] = 123500
		interpreter.varEnv["f"] = 0
		interpreter.instnIndex = 20
		//mulFreq := interpreter.mulFreq
		//fmt.Println("mul freq - ", mulFreq)
	*/
	interpreter.run()
	fmt.Println(interpreter.varEnv)
	fmt.Printf("")

	/*
		interpreter2 := createInterpreter(instnList)
		interpreter2.varEnv["a"] = 1
		interpreter2.run()
		fmt.Println("Val in h - ", interpreter2.varEnv["h"])
	*/
}

func getIpListFromFilename(filename string) []Instruction {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	var instnList []Instruction
	for _, ipStr := range ipStrList {
		ipStr = strings.TrimSpace(ipStr)
		ipComponents := strings.Split(ipStr, " ")
		cmdName := ipComponents[0]
		arg1 := ipComponents[1]
		arg2 := ""
		if len(ipComponents) > 2 {
			arg2 = ipComponents[2]
		}

		instn := Instruction{
			cmdName,
			arg1,
			arg2,
		}
		instnList = append(instnList, instn)
	}
	return instnList
}
