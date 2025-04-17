package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		log.Println("Предупреждение: передан пустой набор данных")
		return
	}

	for i, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга (строка %d): %v\n", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации (строка %d): %v\n", i+1, err)
			continue
		}

		// Убираем лишнее форматирование, оставляем только info
		fmt.Println(info)
	}
}
