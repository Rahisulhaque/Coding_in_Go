/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   typeof.go                                                       __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/26 13:35:30 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/26 14:28:37 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i int
	var f float32
	var str string
	var arr [3]int
	var ans bool

	//	We will use refelct module to check the data type

	fmt.Println(reflect.TypeoOf(i))
	fmt.Println(reflect.TypeoOf(f))
	fmt.Println(reflect.TypeoOf(str))
	fmt.Println(reflect.TypeoOf(ans))
	fmt.Println(reflect.TypeoOf(arr))
}
