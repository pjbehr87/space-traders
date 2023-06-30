package stlib

import (
	"encoding/json"
	"fmt"
	"time"
)

type Cargo struct {
	Capacity  int
	Units     int
	Inventory []InventoryItem
}

type Crew struct {
	Current  int
	Required int
	Capacity int
	Rotation string
	Morale   int
	Wages    int
}

type Engine struct {
	Symbol       string
	Name         string
	Description  string
	Condition    int
	Speed        int
	Requirements ShipRequirements
}

type FrameJson struct {
	Symbol         string
	Name           string
	Description    string
	Condition      int
	ModuleSlots    int `json:"moduleSlots"`
	MountingPoints int `json:"mountingPoints"`
	FuelCapacity   int `json:"fuelCapacity"`
	Requirements   ShipRequirements
}
type Frame struct {
	Symbol         string
	Name           string
	Description    string
	Condition      int
	ModuleSlots    int
	MountingPoints int
	FuelCapacity   int
	Requirements   ShipRequirements
}

type Fuel struct {
	Current  int
	Capacity int
	Consumed FuelConsumed
}

type FuelConsumed struct {
	Amount    int
	Timestamp string
}

type InventoryItem struct {
	Symbol      string
	Name        string
	Description string
	Units       int
}

type Module struct {
	Symbol       string
	Capacity     int
	Range        int
	Name         string
	Description  string
	Requirements ShipRequirements
}

type Mount struct {
	Symbol       string
	Name         string
	Description  string
	Strength     int
	Deposits     []string
	Requirements ShipRequirements
}

type NavJson struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          NavRouteJson
	Status         string
	FlightMode     string `json:"flightMode"`
}
type Nav struct {
	WaypointSymbol *WaypointSymbol
	Route          NavRoute
	Status         string
	FlightMode     string
}

type NavRouteJson struct {
	Destination   RoutePointJson
	Departure     RoutePointJson
	DepartureTime string
	Arrival       string
}
type NavRoute struct {
	Destination   RoutePoint
	Departure     RoutePoint
	DepartureTime time.Time
	Arrival       time.Time
}

type Reactor struct {
	Symbol       string
	Name         string
	Description  string
	Condition    int
	PowerOutput  int `json:"powerOutput"`
	Requirements ShipRequirements
}

type Registration struct {
	Name          string
	FactionSymbol string
	Role          string
}

type RoutePointJson struct {
	Symbol       string
	Type         string
	SystemSymbol string
	X            int
	Y            int
}
type RoutePoint struct {
	Symbol *WaypointSymbol
	Type   string
	X      int
	Y      int
}

type ShipRequirements struct {
	Power int
	Crew  int
	Slots int
}

type ShipJson struct {
	Symbol       string
	Registration Registration
	Nav          NavJson
	Crew         Crew
	Frame        FrameJson
	Reactor      Reactor
	Engine       Engine
	Modules      []Module
	Mounts       []Mount
	Cargo        Cargo
	Fuel         Fuel
}
type Ship struct {
	Symbol       string
	Registration Registration
	Nav          Nav
	Crew         Crew
	Frame        Frame
	Reactor      Reactor
	Engine       Engine
	Modules      []Module
	Mounts       []Mount
	Cargo        Cargo
	Fuel         Fuel
}

type ShipyardShipJson struct {
	// "type" is a reserved word
	ShipType      string `json:"type"`
	Name          string
	Description   string
	PurchasePrice int `json:"purchasePrice"`
	Frame         FrameJson
	Reactor       Reactor
	Engine        Engine
	Modules       []Module
	Mounts        []Mount
}
type ShipyardShip struct {
	ShipType      string
	Name          string
	Description   string
	PurchasePrice int
	Frame         Frame
	Reactor       Reactor
	Engine        Engine
	Modules       []Module
	Mounts        []Mount
}

type ShipTypeJson struct {
	Type string
}

type ShipyardJson struct {
	Symbol       string
	ShipTypes    []ShipTypeJson `json:"shipTypes"`
	Transactions []TransactionJson
	Ships        []ShipyardShipJson
}
type Shipyard struct {
	Symbol       string
	ShipTypes    []string
	Transactions []Transaction
	Ships        []ShipyardShip
}

type TransactionJson struct {
	WaypointSymbol string `json:"waypointSymbol"`
	ShipSymbol     string `json:"shipSymbol"`
	Price          int
	AgentSymbol    string `json:"agentSymbol"`
	Timestamp      string
}
type Transaction struct {
	WaypointSymbol *WaypointSymbol
	ShipSymbol     string
	Price          int
	AgentSymbol    string
	Timestamp      time.Time
}

type ShipData struct {
	Ship ShipJson `json:"data"`
}
type ShipsData struct {
	Ships []ShipJson `json:"data"`
	Meta  MetaPagination
}
type ShipyardData struct {
	Shipyard ShipyardJson `json:"data"`
}

func (stl *StLib) GetShipyard(system string, waypoint string) (Shipyard, error) {
	data, err := stl.GetUrl(fmt.Sprintf("systems/%s/waypoints/%s/shipyard", system, waypoint))
	if err != nil {
		return Shipyard{}, err
	}

	shipyardData := ShipyardData{}
	err = json.Unmarshal(data, &shipyardData)
	if err != nil {
		return Shipyard{}, err
	}

	sdj := shipyardData.Shipyard

	shipTypes := []string{}
	for _, sdst := range shipyardData.Shipyard.ShipTypes {
		shipTypes = append(shipTypes, sdst.Type)
	}
	transactions := []Transaction{}
	for _, stj := range shipyardData.Shipyard.Transactions {
		timestamp, err := time.Parse("2006-01-02T15:04:05Z07:00", stj.Timestamp)
		if err != nil {
			return Shipyard{}, err
		}
		transaction := Transaction{
			WaypointSymbol: getWps(stj.WaypointSymbol),
			ShipSymbol:     stj.ShipSymbol,
			Price:          stj.Price,
			AgentSymbol:    stj.AgentSymbol,
			Timestamp:      timestamp,
		}

		transactions = append(transactions, transaction)
	}
	ships := []ShipyardShip{}
	for _, ssj := range shipyardData.Shipyard.Ships {
		ships = append(ships, ShipyardShip{
			ShipType:      ssj.ShipType,
			Name:          ssj.Name,
			Description:   ssj.Description,
			PurchasePrice: ssj.PurchasePrice,
			Frame: Frame{
				Symbol:         ssj.Frame.Symbol,
				Name:           ssj.Frame.Name,
				Description:    ssj.Frame.Description,
				Condition:      ssj.Frame.Condition,
				ModuleSlots:    ssj.Frame.ModuleSlots,
				MountingPoints: ssj.Frame.MountingPoints,
				FuelCapacity:   ssj.Frame.FuelCapacity,
				Requirements:   ssj.Frame.Requirements,
			},
			Reactor: ssj.Reactor,
			Engine:  ssj.Engine,
			Modules: ssj.Modules,
			Mounts:  ssj.Mounts,
		})
	}
	return Shipyard{
		Symbol:       sdj.Symbol,
		ShipTypes:    shipTypes,
		Transactions: transactions,
		Ships:        ships,
	}, nil
}

func (stl *StLib) PurchaseShip(shipType string, waypointSymbol string) error {
	formData := []byte(fmt.Sprintf(`{
		"shipType": "%s",
		"waypointSymbol": "%s"
	}`, shipType, waypointSymbol))
	_, err := stl.PostUrl("my/ships", &formData)
	if err != nil {
		return err
	}
	return nil
}

func (sj *ShipJson) toShip() (Ship, error) {
	departureTimeT, err := time.Parse("2006-01-02T15:04:05Z07:00", sj.Nav.Route.DepartureTime)
	if err != nil {
		return Ship{}, err
	}
	arrivalT, err := time.Parse("2006-01-02T15:04:05Z07:00", sj.Nav.Route.Arrival)
	if err != nil {
		return Ship{}, err
	}
	return Ship{
		Symbol: sj.Symbol,
		Registration: Registration{
			Name:          sj.Registration.Name,
			FactionSymbol: sj.Registration.FactionSymbol,
			Role:          sj.Registration.Role,
		},
		Nav: Nav{
			WaypointSymbol: getWps(sj.Nav.WaypointSymbol),
			Route: NavRoute{
				Destination: RoutePoint{
					Symbol: getWps(sj.Nav.Route.Destination.Symbol),
					Type:   sj.Nav.Route.Destination.Type,
					X:      sj.Nav.Route.Destination.X,
					Y:      sj.Nav.Route.Destination.Y,
				},
				Departure: RoutePoint{
					Symbol: getWps(sj.Nav.Route.Destination.Symbol),
					Type:   sj.Nav.Route.Destination.Type,
					X:      sj.Nav.Route.Destination.X,
					Y:      sj.Nav.Route.Destination.Y,
				},
				DepartureTime: departureTimeT,
				Arrival:       arrivalT,
			},
		},
		Crew:    sj.Crew,
		Frame:   Frame(sj.Frame),
		Reactor: sj.Reactor,
		Engine:  sj.Engine,
		Modules: sj.Modules,
		Mounts:  sj.Mounts,
	}, nil
}

func (stl *StLib) GetShips() ([]Ship, error) {
	data, err := stl.GetUrl("my/ships")
	if err != nil {
		return []Ship{}, err
	}

	shipsData := ShipsData{}
	err = json.Unmarshal(data, &shipsData)
	if err != nil {
		return []Ship{}, err
	}

	ships := []Ship{}
	for _, sj := range shipsData.Ships {
		ship, err := sj.toShip()
		if err != nil {
			return []Ship{}, err
		}
		ships = append(ships, ship)
	}
	return ships, nil
}
func (stl *StLib) GetShip(shipSymbol string) (Ship, error) {
	data, err := stl.GetUrl(fmt.Sprintf("my/ships/%s", shipSymbol))
	if err != nil {
		return Ship{}, err
	}

	shipData := ShipData{}
	err = json.Unmarshal(data, &shipData)
	if err != nil {
		return Ship{}, err
	}

	ship, err := shipData.Ship.toShip()
	if err != nil {
		return Ship{}, err
	}

	return ship, nil
}
