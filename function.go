package main

import "fmt"
                              //example-1
/*
func add(x int, y int) int{
r := x*y
return r
}
*/
                              //example-2
/*                          
func add(x, y int) int{
r := x*y
return r
}
*/                         
                              //example-3   name returns values                  
/*                          
func add(x, y int) (r int){
r= x*y
return r
}
*/      
                               //example-4 
/*
func add(x, y int) (r int){
r= x*y
return
}
*/ 
                             //returning multiple values
/*                                 
func rectangle(y int, b int) (area int, parameter int) {
parameter = 2 * (y + b)
area = y * b
return
}     
*/
                             //pointer
/*
func update(a *int, t *string) {
*a = *a + 5
*t = *t + "rahoman"
return
} 
*/
func main(){
                              //example-1
/*
x:=add(4, 5)
fmt.Println(x)
*/
                             //example-2    
/*
x:=add(4, 5)
fmt.Println(x)
*/
                              //example-3 
/*
x:=add(4, 5)
fmt.Println(x)
*/
                              //example-4

/*
x:=add(4, 5)
fmt.Println(x)
*/

                              //returning multiple values
/*                              
a,p := rectangle(10, 20)
fmt.Println(a, p)
*/
                              //pointer   
/*
a:=10
t:="ziaur"
update(&a, &t)
fmt.Println(a, t)
*/

                             //anonymous function
a := func(x, y int) (r int){
r=x*y
return
}
fmt.Println(a(10,10))

}