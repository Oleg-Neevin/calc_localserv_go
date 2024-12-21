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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\n  \"error\": \"%s\"\n}", err)
		return
	}
	if request.Expression == "Hello world!" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\n  \"error\": \"It's not a bug. It's a feature\"\n}")
		return
	}
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "{\n  \"error\": \"%s\"\n}", err.Error())
	} else {
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "{\n  \"result\": %f\n}", result)
	}

}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}
