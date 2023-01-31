package todoRouter



import (

	"github.com/gofiber/fiber/v2"
	todoHandler "todobackend/handlers"
)


var TodoMicro = fiber.New()

func GetAll(){

   TodoMicro.Get("/", todoHandler.GetTodosHandler)

}



func GetOne(){


	TodoMicro.Get("/:title", todoHandler.GetTodoHandler)
}


func AddOne(){


	TodoMicro.Post("/", todoHandler.AddTodoHandler)
}

func DeleteOne(){


	TodoMicro.Delete("/:title", todoHandler.DeleteTodoHandler)
}