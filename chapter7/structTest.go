/* ***************************************************************************************** */
/*                                                                                           */
/*                                                                                           */
/*   structTest.go                                                   __    _            __   */
/*                                                       _________ _/ /_  (_)______  __/ /   */
/*   By: rahisul <rahisul@icloud.com>                   / ___/ __ `/ __ \/ / ___/ / / / /    */
/*                                                     / /  / /_/ / / / / (__  ) /_/ / /     */
/*   Created: 2018/03/15 17:31:43 by rahisul          /_/   \__,_/_/ /_/_/____/___,_/_/      */
/*   Updated: 2018/03/15 17:46:02 by rahisul                                                 */
/*                                                                                           */
/* ***************************************************************************************** */
package	main
import	("fmt")


type Movie struct
{
	Name 	string
	Rating	float32
}

func		main(){
	m := Movie{
		Name: "Scarface",
		Rating: 8.2,
		}
	fmt.Println(m.Name, m.Rating)
}


//There shouldn't be any space between function body and parameters
