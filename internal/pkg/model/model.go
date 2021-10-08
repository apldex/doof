/*
 * @Author: Adrian Faisal
 * @Date: 08/10/21 9.02 PM
 */

package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
