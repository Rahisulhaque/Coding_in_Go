/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   conversion.go                                                   __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/04/14 14:57:07 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/04/14 15:46:29 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import "fmt"     //This is for printing
import "strconv" //This is for converting
import "reflect" //This is for type checking
import "os"      //This is for command line arg
func main() {
	var my_var string = os.Args[1]

	var my_str string = "21"

	fmt.Println("The type of my age is ", reflect.TypeOf(my_str))

	my_age, err := strconv.Atoi(my_str)
	if err != nil {
		my_age = 32
	}

	myMindAge, err := strconv.Atoi(my_var)
	if err != nil {
		fmt.Println("That sucks!")
	}
	fmt.Println("My age is ", myMindAge)
	fmt.Println("The type of my_age is ", reflect.TypeOf(my_age))

	fmt.Println(reflect.TypeOf(my_var))
	fmt.Println("***** END *****")
}
