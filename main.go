package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Planner struct {
    TODO    []string
}

func (p *Planner) AddTask(task string){
    p.TODO = append(p.TODO, task)
}

func (p *Planner) DeleteTask(task_index int){
    if (len(p.TODO) == 1){
        p.TODO = []string{}
    } else{
        p.TODO = append(p.TODO[:task_index-1], p.TODO[task_index])
    }
}

func (p *Planner) PrettyPrint(){
    for i, task := range(p.TODO){
        fmt.Println(i+1, "- " ,task)
    }
}

func (p *Planner) WriteChanges(){
    dir, err := os.UserHomeDir()
    check(err)
    todoJson, err := json.Marshal(p.TODO)
    check(err)
    err = os.WriteFile(dir+"/.local/tasks.json", todoJson, 0666)
    check(err)
}

func (p *Planner) ReadState(){
    dir, err := os.UserHomeDir()
    check(err)
    todoJson, err := os.ReadFile(dir+"/.local/tasks.json")
    if (err != nil){
        err = os.WriteFile(dir+"/.local/tasks.json",[]byte("[]"),0755)
        check(err)
        todoJson, err = os.ReadFile(dir+"/.local/tasks.json")
        check(err)
    }
    check(err)
    var todo []string
    err = json.Unmarshal(todoJson, &todo)
    check(err)
    p.TODO = todo

}

func (p *Planner) EditTask(index int){
    list := p.TODO
    var editedTask string = getUserAnswer()
    p.TODO = []string{}
    for i, task := range(list){
        if i == index{
            p.TODO = append(p.TODO, editedTask)
        } else{
            p.TODO = append(p.TODO, task)
        }
    }
    p.WriteChanges()
}

func (p *Planner) Menu(){
    p.ReadState()
    p.PrettyPrint()
    fmt.Println("")
    fmt.Println("")
    fmt.Println("(1) - Write task")
    fmt.Println("(2) - Delete Task")
    var choice string = getUserAnswer()
    choice = strings.Trim(choice, "\n ")
    if (choice == "1"){
        fmt.Print("Type the task: ")
        var newTask string = getUserAnswer()
        p.AddTask(newTask)
        p.WriteChanges()
    } else if choice == "2"{
        fmt.Print("Type the task index: ")
        delTask, err := strconv.Atoi(getUserAnswer())
        check(err)
        p.DeleteTask(delTask)
        p.WriteChanges()
    } else if choice == "3"{
        if len(p.TODO) == 0{
            fmt.Println("Write a task first!!")
        }
        fmt.Print("Type the task index: ")
        editIndex, err := strconv.Atoi(
            strings.Trim(getUserAnswer(), "\n "))
        check(err)
        fmt.Printf("Editing task %s", p.TODO[editIndex-1])
        p.EditTask(editIndex-1)
    }else {
        fmt.Println("Erro! Entrada inválida")
        fmt.Print("Entrada: ")
        fmt.Print(choice)
    }
}

func getUserAnswer() string{
    in := bufio.NewReader(os.Stdin)
    var choice string
    var err error
    choice, err = in.ReadString('\n')
    check(err)
    return choice
}

func main(){
    var p Planner
    p.Menu()
}

func check(err error){
    if (err != nil){
        fmt.Println(err)
        os.Exit(4)
    }
}
