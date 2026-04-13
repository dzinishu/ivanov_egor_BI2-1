package main

import "fmt"

func main() {

	// 1. Создаем срезы (списки)
	subjects := []string{"Математика", "Физика", "Информатика", "История"}
	surnames := []string{"Иванов", "Петров", "Сидоров", "Кузнецов"}
	names := []string{"Алексей", "Мария", "Дмитрий", "Ольга"}
	dates := []string{"01.03.2026", "02.03.2026", "03.03.2026", "04.03.2026"}

	// 2. Что хотим отфильтровать
	filterSubject := "Математика"
	filterDate := "01.03.2026"

	// 3. Перебор каждого с каждым
	for _, subject := range subjects {
		for _, surname := range surnames {
			for _, name := range names {
				for _, date := range dates {

					// 4. Фильтр
					if subject == filterSubject && date == filterDate {
						fmt.Println(subject, surname, name, date)
					}
				}
			}
		}
	}
}
