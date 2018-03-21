/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   array1.go                                                       __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/15 16:18:47 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/15 16:33:57 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package	main

import ("fmt")

func		main(){
	var cheese [2]string
	cheese[0] = "Audrey"
	cheese[1] = "Winters"
	fmt.Println(cheese[0])
	fmt.Println(cheese[1])
	fmt.Println(cheese)
}
