package card

import "bank/pkg/bank/types"

func Total(cards []types.Card) types.Money {
	sum:=types.Money(0)
	//общая сумма средств на всех картах

	for _, card := range cards{
		
		if card.Balance<0 {
			continue
		}
		
		if !card.Active {
			continue
		}
		sum+=card.Balance
	}
	return sum
}
