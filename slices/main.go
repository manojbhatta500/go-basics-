package main

import "fmt"

func main() {
	var routes = []string{"nepalgunz", "kohalpur", "chisapani", "attarya"}
	displayRoutes(1, routes)

	routes = append(routes, "mahendrangar", "banbasa")

	displayRoutes(2, routes)

	visitedRoutes(2, routes)

}

func visitedRoutes(i int, route []string) {
	fmt.Println()
	unvisitedroutes := route[i:]
	newVisitedRoutes := route[0:i]

	fmt.Println("Visited Locations:")

	for i := 0; i < len(newVisitedRoutes); i++ {
		fmt.Println(i+1, ".", newVisitedRoutes[i])
	}

	fmt.Println()

	fmt.Println("Remaining Locations:")

	for i := 0; i < len(unvisitedroutes); i++ {
		fmt.Println(i+1, ".", unvisitedroutes[i])
	}

}

func displayRoutes(routeNumber int, route []string) {

	fmt.Println()

	fmt.Println("RouteNumber:", routeNumber)
	fmt.Println("_____________________", routeNumber, "______________________-")

	for i := 0; i < len(route); i++ {
		fmt.Println(i+1, ".", route[i])
	}
}
