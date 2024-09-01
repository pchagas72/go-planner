package planner

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pchagas72/go-planner/helper"
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
    } else if (task_index == len(p.TODO)){
        p.TODO = p.TODO[:len(p.TODO)-1] 
    }else{
        p.TODO = append(p.TODO[:task_index-1], p.TODO[task_index])
    }
}

func (p *Planner) PrettyPrint(){
    p.ReadState()
    fmt.Println("")
    for i, task := range(p.TODO){
        fmt.Print(i+1, "- " ,task)
    }
    fmt.Println("")
}

func (p *Planner) WriteChanges(){
    dir, err := os.UserHomeDir()
    helper.Check(err)
    todoJson, err := json.Marshal(p.TODO)
    helper.Check(err)
    err = os.WriteFile(dir+"/.local/tasks.json", todoJson, 0666)
    helper.Check(err)
}

func (p *Planner) ReadState(){
    dir, err := os.UserHomeDir()
    helper.Check(err)
    todoJson, err := os.ReadFile(dir+"/.local/tasks.json")
    if (err != nil){
        err = os.WriteFile(dir+"/.local/tasks.json",[]byte("[]"),0755)
        helper.Check(err)
        todoJson, err = os.ReadFile(dir+"/.local/tasks.json")
        helper.Check(err)
    }
    helper.Check(err)
    var todo []string
    err = json.Unmarshal(todoJson, &todo)
    helper.Check(err)
    p.TODO = todo

}

func (p *Planner) EditTask(index int){
    list := p.TODO
    var editedTask string = helper.GetUserAnswer()
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
    fmt.Println("(1) - Write task")
    fmt.Println("(2) - Delete Task")
    fmt.Println("(3) - Edit Task")
    fmt.Println("")
    var choice string = helper.GetUserAnswer()
    choice = strings.Trim(choice, "\n ")
    if (choice == "1"){
        fmt.Print("Type the task: ")
        var newTask string = helper.GetUserAnswer()
        p.AddTask(newTask)
        p.WriteChanges()
    } else if choice == "2"{
        fmt.Print("Type the task index: ")
        delTask, err := strconv.Atoi(strings.Trim(helper.GetUserAnswer(),"\n "))
        helper.Check(err)
        p.DeleteTask(delTask)
        p.WriteChanges()
    } else if choice == "3"{
        if len(p.TODO) == 0{
            fmt.Println("Write a task first!!")
        }
        fmt.Print("Type the task index: ")
        editIndex, err := strconv.Atoi(
            strings.Trim(helper.GetUserAnswer(), "\n "))
        helper.Check(err)
        fmt.Printf("Editing task %s", p.TODO[editIndex-1])
        p.EditTask(editIndex-1)
    }else {
        fmt.Println("Erro! Entrada inválida")
        fmt.Print("Entrada: ")
        fmt.Print(choice)
    }
}

