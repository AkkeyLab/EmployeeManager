/**
---------------------
1. 登録
2. 参照
3. 更新
4. 削除
5. 終了
---------------------
登録情報: 名前、誕生日、性別、給与。登録後にIDが自動生成される。
ちゃんとバリデーションする。
>> イメージ <<
$ ./shaintouroku
---------------------
1. 登録
2. 参照
3. 更新
4. 削除
5. 終了
---------------------
メニュー選択してね> 1(1はユーザーからのinput
名前は？
hoge (ユーザー入力
誕生日は？ (YYYY/MM/DD)
2020/01/あ (ユーザー入力
不正な誕生日です。
誕生日は？ (YYYY/MM/DD)
2020/01/01 (ユーザー入力
給与は？
11111 (ユーザー入力
登録しました。
---------------------
1. 登録
2. 参照
3. 更新
4. 削除
5. 終了
---------------------
メニュー選択してね> 2(1はユーザーからのinput
ID name age salary
------------------
1 hoge 20 1,000
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Employee struct {
	id       string
	name     string
	birthday string
	sex      string
	salary   int
}

func main() {
	menu([]Employee{})
}

func menu(employees []Employee) {
	fmt.Print("メニューを選択してください\n1. 登録\n2. 参照\n3. 更新\n4. 削除\n5. 終了\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "1":
		array := append(employees, *update(new(Employee)))
		menu(array)
	case "2":
		print(employees)
		menu(employees)
	case "3":
		fmt.Print("更新")
	case "4":
		menu([]Employee{})
	case "5":
		os.Exit(0)
	default:
		fmt.Print("不正な入力です")
	}
}

func update(data *Employee) *Employee {
	data.name = validateInput("[^\\s　]", "名前")
	data.birthday = validateInput("^[0-9]{4}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])$", "生年月日（YYYY/MM/DD）")
	data.sex = validateInput("Male|Female", "性別（Male or Female）")
	var salary int
	salary, _ = strconv.Atoi(validateInput("[0-9]", "給与"))
	data.salary = salary
	return data
}

func validateInput(rules string, request string) string {
	data := inputRequest(request + "：")
	if regexp.MustCompile(rules).Match([]byte(data)) {
		return data
	} else {
		fmt.Print("不正な入力です\n")
		return validateInput(rules, request)
	}
}

func print(employees []Employee) {
	for _, employee := range employees {
		fmt.Printf("名前：%s\n生年月日：%s\n性別：%s\n給与：%d\n", employee.name, employee.birthday, employee.sex, employee.salary)
	}
}

func inputRequest(name string) string {
	fmt.Print(name)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
