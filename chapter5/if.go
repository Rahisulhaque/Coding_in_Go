/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   if.go                                                           __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/04/14 13:44:27 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/04/14 14:47:29 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {

	original_input := os.Args[1:]
	fmt.Println("The type of your input is ", reflect.TypeOf(original_input), "\n")
	my_input, err := strconv.Atoi(os.Args[1:])

	fmt.Println("It has been changed to ", reflect.TypeOf(my_input), "\n")

	if my_input > 10 {
		fmt.Println("Thiis greater than 10.\n")
	} else if my_input < 10 {
		fmt.Println("This not okay! Less than 10!\n")
	} else {
		fmt.Println("Exactly!")
	}
}
