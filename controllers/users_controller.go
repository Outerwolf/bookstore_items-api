package controllers

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create()
}

type usersController struct{}

func (c *usersController) Create() {

}
