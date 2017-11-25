package main

import (
  "flag"
  "os"
  "fmt"
  "encoding/csv"
)

func main(){
  csvFileName := flag.String("csv", "problems.csv", "a csv file in the form of 'question,answer' on each line")
  flag.Parse()

  file, err := os.Open(*csvFileName)
  if err != nil {
    exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
  }

  r := csv.NewReader(file)
  lines, err := r.ReadAll()
  if err != nil{
    exit("Error parsing the CSV file")
  }

  fmt.Println(lines)
}

func exit(msg string){
  fmt.Println(msg)
  os.Exit(1)
}
