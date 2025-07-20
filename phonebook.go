// создан репозиторий в github
package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name      string
	Surname   string
	Tel       string
	TabNumber string
}

// глобальная переменная data - срез структур
var data = []Entry{}

// функция фыполняет поиск по заданной фамилии и возвращает полную запись
func searchBySurname(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func searchByTabNumber(key string) *Entry {
	for i, v := range data {
		if v.TabNumber == key {
			return &data[i]
		}
	}
	return nil
}

// фукция выводит все доступные записи
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Владимир", "Шман", "8(926)595-67-47", "1832"})
	data = append(data, Entry{"Максим", "Муравьев", "8(926)112-35-35", "5000"})
	data = append(data, Entry{"Алексей", "Жучков", "8(967)241-66-96", "2018"})

	//различаем команды
	switch arguments[1] {
	//команда поиска
	case "searchBySurname":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := searchBySurname(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)

	case "searchByTabNumber":
		if len(arguments) != 3 {
			fmt.Println("Usage: search TabNumber")
			return
		}
		result := searchByTabNumber(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)

		//команда списка
	case "list":
		list()
		//ответ на все остальное
	default:
		fmt.Println("Not a valid option")
	}

}
