package main
import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID uint `json:"id" gorm:"primarykey"`
	Title string `json:"title"`
	Description string `json:"description"`
	UserId string `json:"user_id"`
	Task string `json:"task"`
	CreatedAt time.Time `json:"created_at"`
}

var users []User
var tasks []Task
var userID = 1
var taskID = 1

var jwtSecret = []byte("secret")

func main(){
  app:= fiber.New();

  app.Post("/tasks" , createTask);
  app.Get("/tasks" , getTask);
  app.Get("/tasks/:id" , getSingleTask);
  app.Patch("/tasks/:id" , updateTask);
  app.Delete("/tasks/:id" , deleteTask);


  app.Listen(" :3000");

}

func createTask(c *fiber.Ctx) error{
	// userID := c.Locals("user_id").(int)
	 var task Task 
	 c.BodyParser(&task)
	 tasks =append(tasks, task)
	 return c.JSON(task)
}

func getTask(c *fiber.Ctx) error{
	 var result []Task
	 
	 return c.JSON(tasks)
}
func getSingleTask(c *fiber.Ctx) error{
	 return c.JSON(tasks)
}

func updateTask(c *fiber.Ctx) error{
	// Implementation for updating a task
	return nil
}
func deleteTask(c *fiber.Ctx) error{
	// Implementation for deleting a task
	return nil
}