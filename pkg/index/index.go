package index

import (
	"awesomeProject/pkg/crawler"
	"strings"
)

// ф-я делит полученную строку на слова и возвращает их в виде массива
func split(s string) []string {
	arr := strings.Fields(s)
	return arr
}

func Indx(m []crawler.Document) map[string]Elem {
	index := map[string]Elem{}

	// i - перебор по ссылкам
	for i := range m {
		words := split(m[i].Title)
		for j := range words { // j- перебор по полю title текущей ссылки
			//index[words[i]] = elem{}
			_, exist := index[words[j]]
			//добавление нового ID если слово существует
			if exist {
				nums := index[words[j]]
				// проверка на содержание аналогичного id в структуре
				s := Contains(nums.Num, m[i].ID)
				if s == false {
					nums.Num = append(nums.Num, m[i].ID)
					index[words[j]] = nums
				}
			}
			//добавление нового элемента если слово не существует
			if !exist {
				newElem := &Elem{}

				newElem.Num = append(newElem.Num, m[i].ID)
				newElem.Count = 1
				index[words[j]] = *newElem
			}
		}
	}
	return index
}

// Contains проверка на содержание элемента в слайсе
func Contains(a []int, id int) bool {
	for _, elem := range a {
		if elem == id {
			return true
		}
	}
	return false
}

// Elem элемент хэша
type Elem struct {
	Num   []int
	Count int
}
