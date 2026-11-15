package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает активность, которую можно разобрать из строки
// и по которой можно сформировать информационное сообщение.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает каждую строку dataset и выводит информацию об активности.
// Ошибки разбора и формирования информации логируются, обработка продолжается.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println("ошибка разбора данных:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка формирования информации:", err)
			continue
		}
		fmt.Println(info)
	}
}
