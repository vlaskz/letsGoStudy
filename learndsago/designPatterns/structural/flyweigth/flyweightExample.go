package flyweigth

import "fmt"

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO type: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

type DataTransferObject interface {
	getId() string
}
type Customer struct {
	id   string
	name string
	ssn  string
}

func (customer Customer) getId() string {
	return customer.id
}

type Employee struct {
	id   string
	name string
}

func (employee Employee) getId() string {
	return employee.id
}

type Manager struct {
	id   string
	name string
	dept string
}

func (manager Manager) getId() string {
	return manager.id
}

type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (address Address) getId() string {
	return address.id
}

func Example() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}

	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer: ", customer.getId())

	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee: ", employee.getId())

	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager: ", manager.getId())

	var address DataTransferObject = factory.getDataTransferObject("address")
	fmt.Println("Address: ", address.getId())
}
