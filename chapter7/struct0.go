/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   struct0.go                                                      __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/15 17:46:48 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/15 17:59:07 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */
package main
import	("fmt")

type City struct{
	Name	string
	County	string
	State	string
	zipcode int	
}
func	main(){
	myCity := City{
		Name: "Santa Monica",
		County:"Los Angeles",
		State: "California",
		zipcode: 90404}

	fmt.Println(myCity.Name, myCity.County, myCity.State, myCity.zipcode)
}
