package main

import (
	"fmt"
	rolemanager "github.com/sorooshjaberi/GoRoleManager/RoleManager"
)

func main() {


	roles := new(rolemanager.Role).Init()
	perms := roles.GetPermisions(5)

	fmt.Printf("perms: %v\n", perms)

}
