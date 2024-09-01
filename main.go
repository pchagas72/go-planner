package main

import (
	"flag"

	"github.com/pchagas72/go-planner/planner"
)
func main(){
    var p planner.Planner
    edit := flag.Bool("edit", false, "Want do edit?")
    flag.Parse()
    if (*edit){
        p.Menu()
    } else{
        p.PrettyPrint()
    }
}

