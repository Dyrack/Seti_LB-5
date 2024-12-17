package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	operand1, _ := strconv.ParseFloat(r.Form.Get("operand1"), 64)
	operand2, _ := strconv.ParseFloat(r.Form.Get("operand2"), 64)
	operationSymbol := r.Form.Get("operation")

	var result float64

	switch operationSymbol {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 != 0 {
			result = operand1 / operand2
		} else {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}
	case "^":
		result = math.Pow(operand1, operand2)
	case "√":
		result = math.Pow(operand1, 1/operand2)
	case "sin":
		result = math.Sin(operand1)
	case "cos":
		result = math.Cos(operand1)
	case "tan":
		result = math.Tan(operand1)
	case "cot":
		result = 1 / math.Tan(operand1)
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%.2f", result)
}
