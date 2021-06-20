package _7_flyweight

const (
	TeamA = iota
	TeamB
)

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

func (t *teamFlyweightFactory) GetTeam(name int) *Team {
	if t.createdTeams[name] != nil {
		return t.createdTeams[name]
	}

	team := getTeamFactory(name)
	t.createdTeams[name] = &team
	return t.createdTeams[name]
}

func getTeamFactory(team int) Team {
	switch team {
	case TeamB:
		return Team{
			ID:   2,
			Name: TeamB,
		}
	default:
		return Team{
			ID:   1,
			Name: TeamA,
		}
	}
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
