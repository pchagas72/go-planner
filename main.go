package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
    dir, oldErr := os.UserHomeDir()
    check(oldErr)
    todoJson, err := json.Marshal(p.TODO)
    check(err)
    os.WriteFile(dir+"/.local/tasks.json", todoJson, 0666)
}

func (p *Planner) ReadState(){
    dir, oldErr := os.UserHomeDir()
    check(oldErr)
    todoJson, err := os.ReadFile(dir+"/.local/tasks.json")
    check(err)
    var todo []string
    NewErr := json.Unmarshal(todoJson, &todo)
    check(NewErr)
    p.TODO = todo

}

func (p *Planner) Menu(){
    p.ReadState()
    p.PrettyPrint()
    fmt.Println("")
    fmt.Println("")
    fmt.Println("(1) - Write task")
    fmt.Println("(2) - Delete Task")
    var choice string = getUserAnswer()
    if (choice == "1"){
        fmt.Print("Type the task: ")
        var newTask string = getUserAnswer()
        p.AddTask(newTask)
        p.WriteChanges()
    } else{
        fmt.Print("Type the task index: ")
        delTask, err := strconv.Atoi(getUserAnswer())
        check(err)
        p.DeleteTask(delTask)
        p.WriteChanges()
    }
}

func getUserAnswer() string{
    var choice string;
    _, err := fmt.Scanln(&choice)
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
