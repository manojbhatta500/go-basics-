package main

import "fmt"

func main() {
	fmt.Println("Welcome to manoj Secret Application")
	fmt.Println("hint: highschool")

	var secrets = map[string]string{
		"money": "my salary is 15,000",
		"email": "my email address is something@gmail.com",
		"study": "i have completed Bacholers in computer applications",
		"watch": "i don't wear watch",
	}

	fmt.Print("Passkey:")

	var hintKey string = "lba"

	var userInputKey string

	fmt.Scan(&userInputKey)

	if hintKey == userInputKey {
		fmt.Println("Logged in successfull")

		secretFunctionality(secrets)

	} else {
		fmt.Println("passkey dosen't matches ")
		fmt.Println("try again")

		main()

	}

}

func showSecretMenu(secrets map[string]string) {
	var i int = 1
	for index, _ := range secrets {

		fmt.Println(i, ".", index)
		i++
	}
}

func secretFunctionality(secrets map[string]string) {

	showSecretMenu(secrets)
	var userInput string
	fmt.Print("now please input words  to see the secrets:")
	fmt.Scan(&userInput)

	currentSecret, err := secrets[userInput]
	if !err {
		fmt.Println("sorry the string dosen't matches please try again:")
		main()

	}
	fmt.Println(currentSecret)

	main()

}
