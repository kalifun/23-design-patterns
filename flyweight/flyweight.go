package flyweight

import "fmt"

// player 
type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

// newPlayer
//  @param playerType
//  @param dressType
//  @return *player
func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

// newLocation
//  @receiver p
//  @param lat
//  @param long
func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// dress
type dress interface {
	getColor() string
}

// const 
//  @param //TerroristDressType terrorist dress type 
const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

// var 
//  @param dressFactorySingleInstance 
//  @param } 
var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

// dressFactory 
type dressFactory struct {
	dressMap map[string]dress
}

// getDressByType 
//  @receiver d 
//  @param dressType 
//  @return dress 
//  @return error 
func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}

	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

// getDressFactorySingleInstance 
//  @return *dressFactory 
func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

// terroristDress
type terroristDress struct {
	color string
}

// getColor
//  @receiver t
//  @return string
func (t *terroristDress) getColor() string {
	return t.color
}

// newTerroristDress
//  @return *terroristDress
func newTerroristDress() *terroristDress {
	return &terroristDress{
		color: "green",
	}
}

// counterTerroristDress 
type counterTerroristDress struct {
	color string
}

// getColor 
//  @receiver t 
//  @return string 
func (t *counterTerroristDress) getColor() string {
	return t.color
}

// newCounterTerroristDress 
//  @return *counterTerroristDress 
func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{
		color: "red",
	}
}

// game 
type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

// newGame 
//  @return *game 
func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

// addTerrorist 
//  @receiver c 
//  @param dressType 
func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

// addCounterTerrorist 
//  @receiver c 
//  @param dressType 
func (c *game) addCounterTerrorist(dressType string) {
    player := newPlayer("CT", dressType)
    c.counterTerrorists = append(c.counterTerrorists, player)
    return
}
