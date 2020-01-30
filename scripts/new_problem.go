// +build newProblem

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt1 := promptui.Prompt{Label: "序号"}
	seq, _ := prompt1.Run()

	prompt2 := promptui.Prompt{Label: "项目"}
	proj, _ := prompt2.Run()

	prompt3 := promptui.Prompt{Label: "题目"}
	problem, _ := prompt3.Run()

	prompt4 := promptui.Prompt{Label: "解答"}
	solution, _ := prompt4.Run()

	prompt5 := promptui.Prompt{Label: "函数名"}
	function, _ := prompt5.Run()

	folder := fmt.Sprintf("problems/%s_%s", seq, proj)
	createFolder := exec.Command("mkdir", folder)
	createFolder.CombinedOutput()

	data := []byte(fmt.Sprintf(`/*

题目: %s

解答: %s

*/

func %s() {
}

func main() {
}`, problem, solution, function))
	ioutil.WriteFile(folder+"/"+"main.go", data, 0644)
}
