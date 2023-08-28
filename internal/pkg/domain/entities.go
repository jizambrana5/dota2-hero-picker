package domain

import "strings"

const (
	Carry     Role = "Carry"
	Disabler  Role = "Disabler"
	Durable   Role = "Durable"
	Escape    Role = "Escape"
	Initiator Role = "Initiator"
	Nuker     Role = "Nuker"
	Support   Role = "Support"
)

type (
	Hero struct {
		HeroIndex        string `json:"hero_index"`
		PrimaryAttribute string `json:"primary_attr"`
		NameInGame       string `json:"localized_name"`
		Role             []Role `json:"roles"`
	}

	UserPreferences struct {
		PrimaryAttribute string `json:"primary_attribute"`
		Roles            []Role `json:"roles"`
	}

	Role string
)

func BuildRoles(s string) []Role {
	roleStr := strings.Replace(s, " ", "", -1)
	singleRoles := strings.Split(roleStr, ",")
	roles := make([]Role, 0)
	for _, v := range singleRoles {
		roles = append(roles, Role(v))
	}
	return roles
}
