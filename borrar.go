package main

import (
	"log"
	"os/exec"
	"github.com/fatih/color"
)


func main(){
	cmd := exec.Command("./lib/main.exe", "-i",  "-o", "war.bin", "-l", "es")
	err := cmd.Run()
	if err != nil {
		color.Red("\narchivo no encontrado")
		log.Fatal(err)
		}
}
