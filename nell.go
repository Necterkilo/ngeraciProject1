/*
* Necter Kilo's Shell = Nell
*/

package main

import (
        "os"
        "os/exec"
        "log"
        "fmt"
        "strings"
        "bufio"
)

func main() {
        prompt := "[user@domain]-> "
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print(prompt)

        for scanner.Scan() {
                userInput := scanner.Text()
                userInputArray := strings.Split(userInput, " ")
                userCmd := userInputArray[0]
                userArgs := userInputArray[1:]

                switch userCmd {
                case "exit":
                        os.Exit(0)
                case "cd":
                        path := ""
                        for _, args := range userArgs {
                                path += args
                        }
                        _, pathErr := os.Stat(path)
                        if pathErr != nil {
                                log.Fatal("Error!", pathErr)
                        }
                        os.Chdir(path)
                case "whoami":
                        fmt.Println("Nicholas A. Geraci - ngeraci")
                default:
                        cmd := exec.Command(userCmd, userArgs...)
                        cmd.Stdout = os.Stdout
                        execErr := cmd.Start()
                        if execErr != nil {
                                log.Fatal("Error!", execErr)
                        }
                        exitErr := cmd.Wait()
                        if exitErr != nil {
                                log.Fatal("Error!", exitErr)
                        }
                }
                fmt.Print(prompt)
        }
}
