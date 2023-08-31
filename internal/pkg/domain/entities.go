package domain

import (
	"strconv"
	"strings"
)

const (
	Carry     Role = "Carry"
	Disabler  Role = "Disabler"
	Durable   Role = "Durable"
	Escape    Role = "Escape"
	Initiator Role = "Initiator"
	Nuker     Role = "Nuker"
	Support   Role = "Support"

	Herald   RankName = "herald"
	Guardian RankName = "guardian"
	Crusader RankName = "crusader"
	Archon   RankName = "archon"
	Legend   RankName = "legend"
	Ancient  RankName = "ancient"
)

type (
	Hero struct {
		HeroIndex        string `json:"hero_index"`
		PrimaryAttribute string `json:"primary_attr"`
		NameInGame       string `json:"localized_name"`
		Role             []Role `json:"roles"`
		WinRates         []Rank `json:"win_rates"`
	}

	UserPreferences struct {
		PrimaryAttribute string   `json:"primary_attribute"`
		Roles            []Role   `json:"roles"`
		RankName         RankName `json:"rank"`
	}

	Role     string
	RankName string
	Rank     struct {
		Name RankName `json:"name"`
		Rate float64  `json:"rate"`
	}
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

func BuildWinRates(h []string) ([]Rank, error) {
	heraldWinRate, err := strconv.ParseFloat(h[6], 64)
	if err != nil {
		return nil, err
	}
	guardianWinRate, err := strconv.ParseFloat(h[9], 64)
	if err != nil {
		return nil, err
	}
	crusaderWinRate, err := strconv.ParseFloat(h[12], 64)
	if err != nil {
		return nil, err
	}
	archonWinRate, err := strconv.ParseFloat(h[15], 64)
	if err != nil {
		return nil, err
	}
	legendWinRate, err := strconv.ParseFloat(h[18], 64)
	if err != nil {
		return nil, err
	}
	ancientWinRate, err := strconv.ParseFloat(h[21], 64)
	if err != nil {
		return nil, err
	}

	return []Rank{
		{
			Name: Herald,
			Rate: heraldWinRate,
		},
		{
			Name: Guardian,
			Rate: guardianWinRate,
		},
		{
			Name: Crusader,
			Rate: crusaderWinRate,
		},
		{
			Name: Archon,
			Rate: archonWinRate,
		},
		{
			Name: Legend,
			Rate: legendWinRate,
		},
		{
			Name: Ancient,
			Rate: ancientWinRate,
		},
	}, nil
}
