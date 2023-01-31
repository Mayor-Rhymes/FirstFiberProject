package main

import (
	"github.com/gofiber/fiber/v2"
	
	todoRouter "todobackend/routers"
	
)




func main(){

   app := fiber.New()


   


   //api/todo micro
   app.Mount("/api/todo", todoRouter.TodoMicro)

   // Path[api/todo/]
   //Method[Get]
   todoRouter.GetAll()

   //Path[api/todo/:title]
   //Method[Get]
   todoRouter.GetOne()

   //Path[api/todo/]
   //Method[Post]
   todoRouter.AddOne()

   //Path[api/todo/:title]
   //Method[Delete]
   todoRouter.DeleteOne()




  


   app.Get("/", func (c *fiber.Ctx) error {


	    return c.SendString("Todo Application with fiber")
   })
   


   
   app.Listen(":3000")

}