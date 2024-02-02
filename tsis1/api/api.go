package api 

type Team struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Chassis    string   `json:"chassis"`
	Engine     string   `json:"engine"`
	TeamChief  string   `json:"teamChief"`
	Established int      `json:"established"`
	Drivers    []string `json:"drivers"`
}

var Teams = []Team{
	{1, "Mercedes", "Mercedes", "Mercedes", "Toto Wolff", 2010, []string{"Lewis Hamilton", "George Russel"}},
	{2, "Red Bull", "Red Bull", "Honda", "Christian Horner", 2005, []string{"Max Verstappen", "Sergio Perez"}},
	{3, "Scuderia Ferrari", "Ferrari", "Ferrari", "Frederic Vasseur", 1950, []string{"Charles Leclerc", "Carlos Sainz"}},
	{4, "McLaren", "McLaren", "Mercedes", "Andrea Stella", 1966, []string{"Oscar Piastri", "Lando Norris"}},
	{5, "Aston Martin", "Aston Martin", "Mercedes", "Otmar Szafnauer", 2018, []string{"Fernando Alonso", "Lance Stroll"}},
	{6, "Alpine", "Alpine", "Renault", "Bruno Famin", 2021, []string{"Pierre Gasly", "Esteban Ocon"}},
	{7, "AlphaTauri", "AlphaTauri", "Honda", "Franz Tost", 2006, []string{"Daniel Ricciardo", "Yuki Tsunoda"}},
	{8, "Alfa Romeo", "Alfa Romeo", "Ferrari", "Frederic Vasseur", 1950, []string{"Valtteri Bottas", "Zhou Guanyu"}},
	{9, "Haas", "Haas", "Ferrari", "Guenther Steiner", 2016, []string{"Nicolas Hulkenberg", "Kevin Magnussen"}},
	{10, "Williams Racing", "Williams", "Mercedes", "James Vowles", 1977, []string{"Logan Sargeant", "Alex Albon"}},
}
func GetTeamByName(name string) *Team {
	for i := range Teams {
		if Teams[i].Name == name {
			return &Teams[i]
		}
	}
	return nil
}