package domain

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
		HeroIndex        int    `json:"hero_index"`
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
