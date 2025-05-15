package main

//Пошук по тексту
//Опис программи:
//программа аналізує текст введений користувачем
//вона дозволяе:
// - визначити скільки разів зустрічаеться слово у тексті
// - підрахувати кількість слів в тексті
// - вивести найдовше слова і його довжину
// - знайти перше слово яке починається на вказану літеру
//пілся аналізу коричтувачу дадуть вибір на повторний аналіз

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функція для пошуку слів що починаються на задану літеру
func SWord(text, word string) {
	//видалення пробілів и преведення до нижнього регістру text та word
	text = strings.ToLower(strings.TrimSpace(text))
	word = strings.ToLower(strings.TrimSpace(word))

	//підрахунок кількості слів які починаються на літеру користувача
	wordCount := strings.Count(text, word)

	//вивід інформації користувачу
	if wordCount == 0 {
		fmt.Printf("Слово \"%s\" немає в тексті.\n", word)
	} else if wordCount == 1 {
		fmt.Printf("Слово \"%s\" зустрічається %d раз.\n", word, wordCount)
	} else {
		fmt.Printf("Слово \"%s\" зустрічається %d рази.\n", word, wordCount)
	}
}

// функція для підрахунку кількості слів в тексті
func allWord(text string) {
	//розбиття тексту
	words := strings.Fields(text)

	//вивід інформації користувачу
	fmt.Printf("У тексті всього %d слів.\n", len(words))
}

// функція для пошуку першого слова яке починається с заданої літери
func searchWord(text, word string) {
	//розбиття тексту
	allwords := strings.Split(text, " ")

	//створення змінной для шуканого Слова
	var sWord string

	//цикл який находить перше слово яке починаеться с заданої літери
	for _, w := range allwords {
		if strings.HasPrefix(strings.ToLower(w), strings.ToLower(word)) {
			sWord = w
			break
		}
	}

	//вивід інформації
	if sWord != "" {
		fmt.Printf("Перше слово, що починається на \"%s\": %s \n", word, sWord)
	} else {
		fmt.Printf("Слів які містять \"%s\" не знайдено.\n", word)
	}
}

// функція для підрахунку найдовшого слова в тексті
func longWord(text string) {
	//розбіття текста
	words := strings.Fields(text)

	//цикл для знаходження найдовшого слова
	long := words[0]
	for _, w := range words {
		if len([]rune(w)) > len([]rune(long)) {
			long = w
		}
	}

	//вивід інформації
	fmt.Printf("Найдовще слово в тексті: %s його довжина: %d \n", long, len([]rune(long)))
}

// основна функція
func main() {
	//створення змінних input, word, sWords, choise
	var input, word, sWords string
	var choise string

	//створення консолього читачя
	reader := bufio.NewReader(os.Stdin)

	//бесконечний цикл
	for {
		//цикл для перевірки на пустий текст
		for {
			fmt.Println("Введіть текст для аналізу: ")
			input, _ = reader.ReadString('\n')
			if strings.TrimSpace(input) == "" {
				fmt.Println("Текст не повинен бути пустим!")
				continue
			} else {
				break
			}
		}

		fmt.Println("\nВведіть слово для пошуку: ")
		fmt.Scanln(&word)
		SWord(input, word)

		allWord(input)

		longWord(input)

		fmt.Println("\nВведіть літеру для пошуку першого слова: ")
		fmt.Scanln(&sWords)
		searchWord(input, sWords)

		//цикл для повтореного аналізу тексту
		for {
			fmt.Println("Хочете проаналізувати інший текст? (yes/no)")
			fmt.Scanln(&choise)

			switch choise {
			case "no":
				fmt.Println("До побачення!")
				return
			case "yes":
				break
			default:
				fmt.Println("Такого Варіанта немає!")
				continue
			}
			break
		}
	}

}
