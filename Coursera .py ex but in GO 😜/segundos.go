package main

import "fmt"

func main() {
	var segundos int
	fmt.Println("Por favor, entre com o número de segundos que deseja converter:")
	fmt.Scan(&segundos)

	dia := segundos / 86400
	sobraDia := segundos % 86400
	hora := sobraDia / 3600
	sobraHora := sobraDia % 3600
	minuto := sobraHora / 60
	sobraMinuto := sobraHora % 60
	segundo := sobraMinuto

	fmt.Println(dia, "dias,", hora, "horas,", minuto, "minutos e", segundo, "segundos.")
}
