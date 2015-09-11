package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"sync"
)

func exe_cmd(cmd []string, wg *sync.WaitGroup) {
	fmt.Println(cmd)
	for _, str := range cmd {
		fmt.Println(str)
	}
	out, err := exec.Command(cmd[0], cmd[1], cmd[2], cmd[3]).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
}

func main() {
	response, err := http.Get("http://gatry.com")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		r := regexp.MustCompile("(?s)article.*?article")
		for _, result := range r.FindAllString(string(contents), -1) {
			//fmt.Printf("%d %s\n", i, result)
			regexName1 := regexp.MustCompile("(?s)h3.*?</h3>")
			fmt.Printf("NAME1: %s\n", regexName1.FindString(result))
			regexName2 := regexp.MustCompile("(?s)target=\"_blank\">(.*)?</a>")
			productName := regexName2.FindStringSubmatch(regexName1.FindString(result))[1]
			fmt.Printf("NAME2: %s\n", productName)
			regexValor1 := regexp.MustCompile("(?s)\"price\">.*?</span>")
			fmt.Printf("VALOR1: %s\n", regexValor1.FindString(result))
			regexValor2 := regexp.MustCompile(">(.*)<")
			productPrice := regexValor2.FindStringSubmatch(regexValor1.FindString(result))[1]
			fmt.Printf("VALOR2: %s\n", productPrice)
			wg := new(sync.WaitGroup)
			commands := []string{"notify-send", "--expire-time=30000", productName, "R$" + productPrice}
			wg.Add(1)
			go exe_cmd(commands, wg)
			wg.Wait()
			break
		}
		//fmt.Printf("%s\n", string(contents))
	}

}
