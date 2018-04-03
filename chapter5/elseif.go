/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   elseif.go                                                       __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/31 16:39:43 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/31 16:54:00 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import (
	"fmt"
)

func main() {

	b := "audrey"

	if b == "jordyn" {
		fmt.Println("Nope")
	} else if b == "adiana" {
		fmt.Println("No no no!")
	} else if b == "jaquilyn" {
		fmt.Println("Nope")
	} else {
		fmt.Println("That was Audrey!")
	}
}
