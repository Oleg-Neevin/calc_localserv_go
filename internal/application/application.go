package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Oleg-Neevin/calc_localserv_go/pkg/calculation"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

// Функция запуска приложения
// тут будем читать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("application was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed with error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
	} else {
		fmt.Fprintf(w, "result: %f", result)
	}

}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":8080", nil)
}
