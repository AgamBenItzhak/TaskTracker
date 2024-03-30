/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/AgamBenItzhak/TaskTracker/cmd"
	_ "github.com/AgamBenItzhak/TaskTracker/config"
	_ "github.com/AgamBenItzhak/TaskTracker/logger"
)

func main() {
	cmd.Execute()
}
