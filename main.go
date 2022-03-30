/*
Copyright © 2022 Drake Axelrod <drake@draxel.io>
*/

package main

import (
	"github.com/DrakeAxelrod/annoyme/data"
	"github.com/DrakeAxelrod/annoyme/cmd"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
