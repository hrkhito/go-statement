package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func getOsName() string {
	return "heyhey"
}

func foo() {
	// deferを付与した箇所はfoo関数内の処理が終わったら実行される
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogfile)
}

func thirdPartyConnectDB() {
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	// if文

	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	// result2はresult2を定義したif文でのみしか使用できない
	// fmt.Println(result2)

	num := 6
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}

	// for文

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)

	// for {
	// 	fmt.Println("hello")
	// }

	// range

	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	// switch文
	targetos := getOsName()
	switch targetos {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default!!", targetos)
	}

	// 変数osはswitch文内でのみしか使用できない
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Good night")
	}

	// defer

	foo()
	// deferを付与するとmain関数内の処理が終わった後に行われる
	defer fmt.Println("world")
	fmt.Println("Hello")

	// deferが複数あるときは一番下のdeferから上に実行されていく
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// deferの使用例 (ファイルを読み込む処理)
	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

	// log

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// Fatal~系を使用してしまったらプログラムが終了されてしまうので注意
	// log.Fatalf("%T %v", "test", "test")
	// log.Fatalln("error!!")

	// _, err := os.Open("fdafdafada")
	// if err != nil {
	// 	log.Fatalln("Exit", err)
	// }

	// LoggingSettings("test.log")

	// エラーハンドリング

	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}

	data1 := make([]byte, 100)
	// countが初めて定義されるのでこのerrはそれに併せて新しいものになる
	count, err := file.Read(data1)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	// panicとrecover

	save()

	// Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。

	ll := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var resultMin int
	resultMin = ll[0]

	for _, l := range ll {
		if resultMin > l {
			resultMin = l
		}
	}

	fmt.Println(resultMin)

	// Q2. 以下の果物の価格の合計を出力するコードを書いてください。

	mm := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var total int

	for _, v := range mm {
		total = total + v
	}

	fmt.Println(total)

}
