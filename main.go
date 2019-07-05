package main

import (
	"fmt"
)

// Mutex for AllStudentsRecord and TotalStudentCount
//var StudentMutex = &sync.Mutex{}
func main() {


	var reTry string = "Y"
	for reTry == "Y" || reTry == "y" {
		reTry="n"

		fmt.Print("\n( Enter n to exist). \t \t Please enter Choice. \n 1. interface implementation \n 2. Empty Interface \n 3. Type Assertion \n 4. first class function.\n 5. reflection in golang " +
			"\n 6. deepcopy vs shallow \n 7. Check two maps are same or not \n 8.copy map to other \n 9. fallthrough in golang \n 10. Tree Travers In Pre Post : ")
		var choice int
		fmt.Scanf("%d",&choice)
			switch choice {
			case 1:
				InterFaceMain(false)
				fallthrough

			case 2:
				InterFaceMain(true)
				StackMain()
				fallthrough

			case 3:
				TypeAssertion()
				fallthrough

			case 4:
				firstclassFunction()
				fallthrough

			case 5:
				CheckEqualUseReflect()
				fallthrough

			case 6:
				DeepVsShallowCopy()
				fallthrough

			case 7:
				CheckEqualUseReflect()
				fallthrough

			case 8:
				CopyMap()
				fallthrough

			case 9:
				fmt.Println("Testing fallthrough in golang")
				//break
				fallthrough
			case 91:
				fmt.Println(" fallthrough from case 9 to 91")
				//break
				fallthrough
			case 10:
				TreeTraversInPrePost()
			default:
				fmt.Println("You have not Spesified any Choice")

			}

		fmt.Printf("Do you wan to continue (y/n):")
		fmt.Scanf("%s",&reTry)
	}


	//StudentMutex.Unlock()

}

