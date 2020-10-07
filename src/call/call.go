package main
 
import "fmt"
import "github.com/carlescere/scheduler"
import "runtime"
 
func main() {
    // 5秒に1回 success!! と出力させる
    scheduler.Every(5).Seconds().Run(printSuccess)
    runtime.Goexit()
}
 
func printSuccess() {
    fmt.Printf("success!! \n")
}