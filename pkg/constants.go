package pkg

var (
	Competitions = map[string]string{
		SixNations:              SixNationsId,
		RugbyWorldCup:           RugbyWorldCupId,
		Premiership:             PremiershipId,
		Top14:                   Top14Id,
		UnitedRugbyChampionship: UnitedRugbyChampionshipId,
		RugbyChampionship:       RugbyChampionshipId,
	}
)

const (
	CrawlBaseUrl   = "https://www.espn.co.uk/rugby/"
	CrawlCompField = "table/_/league/"

	SixNations   = "Six Nations"
	SixNationsId = "180659"

	RugbyWorldCup   = "Rugby World Cup"
	RugbyWorldCupId = "164205"

	Premiership   = "Premiership"
	PremiershipId = "267979"

	Top14   = "Top 14"
	Top14Id = "270559"

	UnitedRugbyChampionship   = "United Rugby Championship"
	UnitedRugbyChampionshipId = "270557"

	RugbyChampionship   = "Rugby Championship"
	RugbyChampionshipId = "244293"
)
