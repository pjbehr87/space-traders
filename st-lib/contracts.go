package stlib

import (
	"encoding/json"
	"time"
)

type ContractJson struct {
	Id               string
	FactionSymbol    string
	Type             string
	Terms            TermsJson
	Accepted         bool
	Fulfilled        bool
	Expiration       string
	DeadlineToAccept string
}
type Contract struct {
	Id               string
	FactionSymbol    string
	Type             string
	Terms            Terms
	Accepted         bool
	Fulfilled        bool
	Expiration       time.Time
	DeadlineToAccept time.Time
}

type ContractsData struct {
	Contracts []ContractJson `json:"data"`
}
type ContractData struct {
	Contract ContractJson `json:"data"`
}

type TermsJson struct {
	Deadline string
	Payment  Payment
	Deliver  []DeliverJson
}
type Terms struct {
	Deadline   time.Time
	Payment    Payment
	Deliveries []Deliver
}

type Payment struct {
	OnAccepted  int64 `json:"onAccepted"`
	OnFulfilled int64 `json:"onFulfilled"`
}

type DeliverJson struct {
	TradeSymbol       string
	DestinationSymbol string
	UnitsRequired     int64
	UnitsFulfilled    int64
}
type Deliver struct {
	TradeSymbol       string
	DestinationSymbol *WaypointSymbol
	UnitsRequired     int64
	UnitsFulfilled    int64
}

func (stl *StLib) AcceptContract(contractId string) error {
	_, err := stl.PostUrl("my/contracts/" + contractId + "/accept")
	if err != nil {
		return err
	}
	return nil
}

func (stl *StLib) ListContracts() ([]Contract, error) {
	resp, err := stl.GetUrl("my/contracts")
	if err != nil {
		return []Contract{}, err
	}

	contractsData := ContractsData{}
	err = json.Unmarshal(resp, &contractsData)
	if err != nil {
		return []Contract{}, err
	}

	contracts := []Contract{}
	for _, cd := range contractsData.Contracts {
		deadlineT, err := time.Parse("2006-01-02T15:04:05Z07:00", cd.Terms.Deadline)
		if err != nil {
			return []Contract{}, err
		}
		deliveries := []Deliver{}
		for _, cdtd := range cd.Terms.Deliver {
			deliveries = append(deliveries, Deliver{
				TradeSymbol:       cdtd.TradeSymbol,
				DestinationSymbol: getWps(cdtd.DestinationSymbol),
				UnitsRequired:     cdtd.UnitsRequired,
				UnitsFulfilled:    cdtd.UnitsFulfilled,
			})
		}

		expirationT, err := time.Parse("2006-01-02T15:04:05Z07:00", cd.Expiration)
		if err != nil {
			return []Contract{}, err
		}
		deadlineToAcceptT, err := time.Parse("2006-01-02T15:04:05Z07:00", cd.DeadlineToAccept)
		if err != nil {
			return []Contract{}, err
		}
		c := Contract{
			Id:            cd.Id,
			FactionSymbol: cd.FactionSymbol,
			Type:          cd.Type,
			Terms: Terms{
				Deadline: deadlineT,
				Payment: Payment{
					OnAccepted:  cd.Terms.Payment.OnAccepted,
					OnFulfilled: cd.Terms.Payment.OnFulfilled,
				},
				Deliveries: deliveries,
			},
			Accepted:         cd.Accepted,
			Fulfilled:        cd.Fulfilled,
			Expiration:       expirationT,
			DeadlineToAccept: deadlineToAcceptT,
		}

		contracts = append(contracts, c)
	}
	return contracts, nil
}

func (stl *StLib) GetContract(contractId string) (Contract, error) {
	resp, err := stl.GetUrl("my/contracts/" + contractId)
	if err != nil {
		return Contract{}, err
	}

	contractData := ContractData{}
	err = json.Unmarshal(resp, &contractData)
	if err != nil {
		return Contract{}, err
	}

	cj := contractData.Contract
	deadlineT, err := time.Parse("2006-01-02T15:04:05Z07:00", cj.Terms.Deadline)
	if err != nil {
		return Contract{}, err
	}
	deliveries := []Deliver{}
	for _, cjtd := range cj.Terms.Deliver {
		deliveries = append(deliveries, Deliver{
			TradeSymbol:       cjtd.TradeSymbol,
			DestinationSymbol: getWps(cjtd.DestinationSymbol),
			UnitsRequired:     cjtd.UnitsRequired,
			UnitsFulfilled:    cjtd.UnitsFulfilled,
		})
	}

	expirationT, err := time.Parse("2006-01-02T15:04:05Z07:00", cj.Expiration)
	if err != nil {
		return Contract{}, err
	}
	deadlineToAcceptT, err := time.Parse("2006-01-02T15:04:05Z07:00", cj.DeadlineToAccept)
	if err != nil {
		return Contract{}, err
	}

	return Contract{
		Id:            cj.Id,
		FactionSymbol: cj.FactionSymbol,
		Type:          cj.Type,
		Terms: Terms{
			Deadline: deadlineT,
			Payment: Payment{
				OnAccepted:  cj.Terms.Payment.OnAccepted,
				OnFulfilled: cj.Terms.Payment.OnFulfilled,
			},
			Deliveries: deliveries,
		},
		Accepted:         cj.Accepted,
		Fulfilled:        cj.Fulfilled,
		Expiration:       expirationT,
		DeadlineToAccept: deadlineToAcceptT,
	}, nil
}
