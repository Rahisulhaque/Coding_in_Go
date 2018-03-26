/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   varError.go                                                     __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/26 12:58:14 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/26 13:01:06 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import (
	"fmt"
)

func main() {

	var i int
	i = "one"
	fmt.Println(i)
}

//# command-line-arguments
//./varError.go:22:6: cannot use "one" (type string) as type int in assignment
