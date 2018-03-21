/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   bool2.go                                                        __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/21 12:36:43 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/21 12:47:12 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	var answer bool
	fmt.Println(answer)
	answer = true
	fmt.Println(answer)
}

//by default go will assign "false" to your uninitialized boolean
