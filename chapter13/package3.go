/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   package3.go                                                     __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/06/10 10:29:18 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/06/10 10:52:32 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {

	s := ".ti esrever dna ti plif nwod gniht ym tup I"

	fmt.Println("We have imported a packqge from github")
	fmt.Println(stringutil.Reverse(s))
}
