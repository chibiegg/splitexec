package main

import (
    "fmt"
    "flag"
    "io"
    "os"
    "os/exec"
    "log"
    "strings"
)


const BUFFER_SIZE = 134217728

var cmd *exec.Cmd
var writer io.WriteCloser

func main(){
    s := flag.Int("s", 0, "size")
    flag.Parse()

    buf := make([]byte, BUFFER_SIZE)
    remain := *s
    counter := 0

    for {

      if cmd == nil {

        args := []string{}
        for _, a := range flag.Args() {
          if strings.Contains(a, "%") {
            args = append(args, fmt.Sprintf(a, counter))
          } else {
            args = append(args, a)
          }
        }
        log.Printf("exec %#v", args)
        cmd = exec.Command(args[0], args[1:]...)
        writer, _ = cmd.StdinPipe()
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Start()
        counter += 1
      }

      readBuf := buf

      if len(buf) > remain {
        readBuf = make([]byte, remain)
      }

      n, err := os.Stdin.Read(readBuf)

      if err == io.EOF {
        if n == 0 {
          break
        }
      } else if err != nil {
        log.Fatal("Read()", err)
      }

      _, err = writer.Write(readBuf[:n])

      if err != nil {
        log.Fatal("Write()", err)
      }

      remain -= n

      if remain <= 0 {
        writer.Close()
        cmd.Wait()
        remain = *s
        cmd = nil
      }
    }
}
