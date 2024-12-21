package calculation

func Calc(expression string) (float64, error) {
	var operations []rune
	var numbers []float64
	var trash float64
	var count int

	// Проверяем правильность расставления скобок
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '(':
			count++
		case ')':
			count--
		}
		if count < 0 {
			return 0.0, ErrInvalidParentheses
		}
	}
	if count != 0 {
		return 0.0, ErrInvalidParentheses
	}

	// Анализируем вводимые данные:
	count = 0

	for i := 0; i < len(expression); i++ {

		// числа и арифметические знаки записываем в списки
		if expression[i] == '*' || expression[i] == '/' || expression[i] == '+' || expression[i] == '-' {
			operations = append(operations, []rune(expression)[i])

		} else if (expression[i] - '0') <= 9 {
			numbers = append(numbers, float64(expression[i]-'0'))

			// для скобок вызываем рекурсию и их результат записываем в список к числам
		} else if expression[i] == '(' {
			for j := i; ; j++ {
				if expression[j] == '(' {
					count++
				} else if expression[j] == ')' {
					count--
				}
				if expression[j] == ')' && count == 0 {
					trash, _ = Calc(expression[i+1 : j])
					numbers = append(numbers, trash)
					i = j
					break
				}
			}
			// если попадаются другие символы - ошибка
		} else {
			return 0.0, ErrInvalidExpression
		}
	}

	// проверяем количество чисел и знаков
	if len(numbers) != len(operations)+1 {
		return 0.0, ErrInvalidExpression
	}

	// вычисляем преоритетные операции
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case '*':
			numbers = append(append(numbers[:i], numbers[i]*numbers[i+1]), numbers[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--

		case '/':
			if numbers[i+1] == 0 {
				return 0.0, ErrDivisionByZero
			}

			numbers = append(append(numbers[:i], numbers[i]/numbers[i+1]), numbers[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--
		}
	}

	// вычисляем менее преоритетные операции
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case '+':
			numbers = append(append(numbers[:i], numbers[i]+numbers[i+1]), numbers[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--

		case '-':
			numbers = append(append(numbers[:i], numbers[i]-numbers[i+1]), numbers[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--

		}
	}

	return numbers[0], nil

}
