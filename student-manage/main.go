package main
import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type Student struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Age        int     `json:"age"`
	CGPA       float64 `json:"cgpa"`
}

type Payment struct {
	ID         int     `json:"id"`
	Amount     float64 `json:"amount"`
	Time       string  `json:"time"`
	Department string  `json:"department"`
}

var students []Student
var payments []Payment
var err error;

func main(){
	app := fiber.New();

	// routes
	app.Get("/students", getStudents)
	app.Post("/students", createStudent)
	app.Get("/students/:id", getStudent)
	app.Put("/students/:id", updateStudent)
	app.Delete("/students/:id", deleteStudent)

	app.Get("/payments", getPayments)
	app.Post("/payments", createPayment)
	app.Get("/payments/:id", getPayment)
	app.Put("/payments/:id", updatePayment)
	app.Delete("/payments/:id", deletePayment)

	app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Student Management API is running!")
    })

	app.Listen(":3000")
	
}

func getStudents(c *fiber.Ctx) error {
	return c.JSON(students)
}

func createStudent(c *fiber.Ctx) error{
	var s Student
	if err := c.BodyParser(&s);

	err != nil {
		return c.Status(400).JSON("Invalid request body Wehen creating User")
	}

	students = append(students , s)
	return c.Status(200).JSON(s)
}

func updateStudent(c *fiber.Ctx) error{
  id , _ := strconv.Atoi(c.Params("id"))
  
  var body Student
  if err != nil {
	return  c.Status(400).JSON("Error found")
  }

  for i, s := range students {
	if s.ID == id {
      body.ID = s.ID
	  students[i] = body
	  return c.JSON(body)
	}
  }

  return c.Status(401).JSON(fiber.Map{"message": "Not found"})
}

func getStudent(c *fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))

	for _,s := range students {
		if s.ID ==id {
			return  c.JSON(s)
		}
	}

	return c.Status(404).JSON(fiber.Map{"message": "Student not found"})
}

func deleteStudent(c *fiber.Ctx) error {
	id , _ := strconv.Atoi(c.Params("id"))

	for i,s := range students {
		if s.ID == id {
			students = append(students[:i] , students[i+1:]...)
			return c.Status(200).JSON(fiber.Map{"message": "Student deleted successfully"})
		}
	}
	return c.Status(404).JSON(fiber.Map{"message": "Student not found"})
}

// payments route 
func getPayments (c *fiber.Ctx) error {
  return c.JSON(payments)
}

func createPayment(c *fiber.Ctx) error {
	var p Payment 
    if err := c.BodyParser(&p); 

    err != nil {
		return c.Status(400).JSON("Invalid request body Wehen creating User")
	}

	payments = append(payments, p)
	return  c.Status(200).JSON(p)
}

func getPayment (c *fiber.Ctx) error {
	id , _ := strconv.Atoi(c.Params("id"))
	for _,p := range payments {
		if p.ID == id {
			return  c.JSON(p)
		}

	}
	return c.Status(404).JSON(fiber.Map{"message": "Payment not found"})
}

func updatePayment (c *fiber.Ctx)error {
  id,_ := strconv.Atoi(c.Params("id"))
  var body Payment

  if err != nil {
	return  c.Status(400).JSON("Error found")
  }

  for i,p := range payments {
	if p.ID == id {
		body.ID = p.ID
		payments[i] = body
		return c.JSON(body)
	}
  }

  return c.Status(404).JSON(fiber.Map{"message": "Payment not found"});
}

func deletePayment(c *fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	for i,p := range payments {
		if p.ID == id {
			payments = append(payments[:i], payments[i+1:]... )
		}
	}
	return c.Status(200).JSON(fiber.Map{"message": "Payment deleted successfully"})
}