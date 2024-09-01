package helper

import (
    "fmt"
    "os"
    "bufio"
)

func Check(err error){
    if (err != nil){
        fmt.Println(err)
        os.Exit(4)
    }
}


func GetUserAnswer() string{
    in := bufio.NewReader(os.Stdin)
    var choice string
    var err error
    fmt.Print(">>> ")
    choice, err = in.ReadString('\n')
    Check(err)
    return choice
}



