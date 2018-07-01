/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   Types.go                                                        __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/06/25 13:01:59 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/06/25 13:45:07 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import "fmt"

func main() {
	fmt.Println("    *******      *********   ")
    fmt.Println("  ***    **     ***     ***  ")
	fmt.Println(" ***            ***     *** ")
	fmt.Println(" ***   *****    ***     *** ")
	fmt.Println(" ***     ****   ***     *** ")
	fmt.Println("  **********     *********  ")


	var i	int
	var f	float 
	var b	bool 
	var s	string
	var crushes	[4]string 
    
	i = 2125
	f = 9.5
	b = true 
	s = "Audrey Winters"
	crushes[] = ["Audrey", "Jordyn", "Carley", "Sophie"]


	fmt.Println("Here is my variables:\n")
	fmt.Println("My cadet number is %v.\n", i)
	fmt.Println("Here is an example of flaot %v.\n", f)
	fmt.Println("It worked! %v. \n", b)
	fmt.Println("I used to like %v.\n", s)
	fmt.Println("%v was soooo cute!\n", crushes [3])
	}
