package main

import (
	// "bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func init() {
	cmd1 := exec.Command("go", "env")
	cmd2 := exec.Command("grep", "GOROOT")
	studio1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Println("stdout err", err)
		return
	}
	if err := cmd1.Start(); err != nil {
		fmt.Println("Command err", err)
		return
	}
	var outputBufo1 bytes.Buffer
	for {
		tempOutput := make([]byte, 4096)
		n, err := studio1.Read(tempOutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("read err ", err)
			}

		}
		if n > 0 {
			outputBufo1.Write(tempOutput[:n])
		}
	}
	fmt.Println(1, outputBufo1.String(), 1)
	// outputBufo1 := bufio.NewReader(studio1)
	// outputo, _, err := outputBufo1.ReadLine()
	// for{
	// 	outputo, _, err := outputBufo1.ReadLine()
	// }
	// if err != nil {
	// 	fmt.Println("read err")
	// }
	// fmt.Println("1", string(outputo), " 1")

	//-----
	studio2, err := cmd2.StdinPipe()
	if err != nil {
		fmt.Println("cmd2 pipe err ", err)
		return
	}
	outputBufo1.WriteTo(studio2)

	var outputBufo2 bytes.Buffer
	cmd2.Stdout = &outputBufo2
	if err := cmd2.Start(); err != nil {
		fmt.Println("cmd2 start err", err)
		return
	}
	err = studio2.Close() //关闭 通道   因为  部分命令 会等待 结束 所以手动关闭
	if err != nil {
		fmt.Println("close cmd2 err ", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Println("wait err ", err)
		return
	}
	fmt.Println(outputBufo2.String())
}
func main() {

}
