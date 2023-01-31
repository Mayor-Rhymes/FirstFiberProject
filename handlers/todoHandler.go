package todoHandler

import (
	"errors"
	todoModel "todobackend/models"

	"github.com/gofiber/fiber/v2"
)


var data = []todoModel.Todo{

	{Title: "Life", Content:"Yes life does go on and it can be very painful"},
	{Title:"Hello", Content: "Hello to the people who went in before me"},
	{Title:"Time", Content:"Time is not on my side but I am willing to take a huge risk"},
  }


func GetTodosHandler(c *fiber.Ctx) error {
    
         
      //data := 

	  return c.JSON(data)
}


func GetTodoHandler(c *fiber.Ctx) error{

     for _, values := range data{


		if c.Params("title") == values.Title{

			return c.JSON(values)
		} 
	 }

	 return errors.New("nothing was found")

}



func AddTodoHandler(c *fiber.Ctx) error{
    

	newTodo := new(todoModel.Todo)


	if err := c.BodyParser(newTodo); err != nil{


		return err
	} else{

		data = append(data, *newTodo)
	}

	return c.JSON(data)

	
	
	 
}



func DeleteTodoHandler(c *fiber.Ctx) error{

     var newData []todoModel.Todo

	 for _, values := range data{


		if c.Params("title") != values.Title{

			 newData = append(newData, values)
		} 

		
	 }

	 data = newData


	 return c.JSON(data)


}

