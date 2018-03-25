package main
//https://youtu.be/Utf-A4rODH8?list=WL

func main(){
  switch os.Args[1]{
  case "run":
    run()
  default:
    panic("what??")
  }
}

func run(){
  fmt.Printf("running %v\n", os.Args[2:])

  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  must(cmd.Run())
}

func must(err error){
  if err != nil{
    panic(err)
  }
}
