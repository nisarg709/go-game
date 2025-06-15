package service

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Continents []Continent `json:"continents"`
	Games      []Game      `json:"games"`
	Progress   []Milestone `json:"milestones"`
}

type Continent struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
	City []struct {
		Uid      string `json:"uid"`
		Name     string `json:"name"`
		Requires int    `json:"requires"`
		Level    struct {
			Entries    int `json:"entries"`
			Step       int `json:"step"`
			Difficulty int `json:"difficulty"`
		} `json:"level"`
		Next struct {
			Proceed bool   `json:"proceed"`
			City    string `json:"city"`
			Miles   int    `json:"miles"`
		} `json:"next"`
		Games []string `json:"games"`
	} `json:"cities" bson:"cities"`
}

type Game struct {
	Uid        string `json:"uid" bson:"_id"`
	Name       string `json:"name"`
	HelpCost   int    `json:"help_cost" bson:"help_cost"`
	ResumeCost int    `json:"resume_cost" bson:"resume_cost"`
	ForceCost  int    `json:"force_cost" bson:"force_cost"`
	Levels     []struct {
		Difficulty int `json:"difficulty"`
		Tries      int `json:"tries"`
		Rewards    []struct {
			Type string `json:"type"`
			Qty  int    `json:"qty"`
		} `json:"rewards"`
		Config interface{} `json:"config"`
	} `json:"levels"`
}

type Milestone struct {
	Amount  int `json:"amount"`
	Rewards []struct {
		Type string `json:"type"`
		Qty  int    `json:"qty"`
	} `json:"rewards"`
}
