package domain

import (
	"sort"
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

	Agi Attribute = "agi"
	Str Attribute = "str"
	Int Attribute = "int"
	All Attribute = "all"
)

type (
	Hero struct {
		HeroIndex        string    `json:"hero_index"`
		PrimaryAttribute Attribute `json:"primary_attr"`
		NameInGame       string    `json:"localized_name"`
		Role             []Role    `json:"roles"`
		WinRates         []Rank    `json:"win_rates"`
	}

	UserPreferences struct {
		PrimaryAttribute Attribute `json:"primary_attribute" validate:"required,attribute"`
		Roles            []Role    `json:"roles" validate:"required,dive,role"`
		RankName         RankName  `json:"rank_name" validate:"required,rank_name"`
	}
	Heroes    []Hero
	Role      string
	RankName  string
	Attribute string
	Rank      struct {
		Name RankName `json:"name"`
		Rate float64  `json:"rate"`
	}
)

func (r Role) IsValid() bool {
	return r == Carry || r == Disabler || r == Durable || r == Escape ||
		r == Initiator || r == Nuker || r == Support
}

func (r RankName) IsValid() bool {
	return r == Herald || r == Guardian || r == Crusader || r == Archon ||
		r == Legend || r == Ancient
}

func (a Attribute) IsValid() bool {
	return a == Int || a == Str || a == Agi || a == All
}

func (h Heroes) SortHeroesByWinRate(rankName RankName) []Hero {
	sort.Slice(h, func(i, j int) bool {
		return h[i].GetWinRateForRank(rankName) > h[j].GetWinRateForRank(rankName)
	})
	return h
}

func (h Hero) GetWinRateForRank(rankName RankName) float64 {
	for _, winRate := range h.WinRates {
		if winRate.Name == rankName {
			return winRate.Rate
		}
	}
	return 0.0 // Return a default value if the rank name is not found
}

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
