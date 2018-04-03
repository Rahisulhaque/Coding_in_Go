/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   ifelse.go                                                       __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/31 16:08:59 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/31 16:25:54 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import "fmt"

func main() {
	fmt.Println("This is go lang! pretty amazing! Huh?")

	i := false
	if true {
		fmt.Println("Ya! That's true!")
	}
	ifelse(i)
}

func ifelse(x bool) {
	if x {
		fmt.Println("The input is true!")
	} else {
		fmt.Println("The input is false!")
	}
}

//make sure that the { curly braces are in the same line with if word!
//There will be curly braces
