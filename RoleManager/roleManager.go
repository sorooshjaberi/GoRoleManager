package rolemanager

type Role struct {
	roles map[int]string
}

var RolesMap = map[int]string{
	4: "READ",
	2: "WRITE",
	1: "EXECUT",
}


func (role *Role) Init() Role {
	roles := RolesMap

	role.roles = roles

	return *role
}

func (role Role) GetRoleScores() []int {
	scores := []int{}

	for key := range role.roles {
		scores = append(scores, key)
	}

	return scores

}

// 6 -> [2 , 4]
func (role Role) GetSumFromFormula(score int) []int {

	formula := []int{}

	if score < 1 {
		return nil
	}

	scores := role.GetRoleScores()

	for _, scoreValue := range scores {
		resultScores := role.GetSumFromFormula(score - scoreValue)

		sum := 0

		for _, calculatedScore := range resultScores {
			sum += calculatedScore
		}

		resultOfSums := scoreValue + sum

		if resultOfSums == score {
			return append(resultScores, scoreValue)
		}

	}

	return formula

}

func (role Role) GetPermisions(score int) []string {
	perms := []string{}

	scores := role.GetSumFromFormula(score)

	for _, score := range scores {
		perms = append(perms, role.roles[score])
	}

	return perms
}
