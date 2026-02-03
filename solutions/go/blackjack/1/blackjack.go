package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
        case "two":
        	return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
		case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
        	return 9
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
        return "P"
    }
    cardSum := ParseCard(card1) + ParseCard(card2)
    switch {
    	case cardSum == 21:
    		if shouldWin(dealerCard) {
                return "W"
            } else {
                return "S"
            }
    	case cardSum > 16 && cardSum < 21:
    		return "S"
    	case cardSum > 11 && cardSum < 17:
    		if ParseCard(dealerCard) >= 7 {
                return "H"
            } else {
                return "S"
            }
        case cardSum <= 11:
            return "H"
		}
    return "Unknown"
}

func shouldWin(dealerCard string) bool {
    return (dealerCard != "ace") && !isFace(dealerCard) && (dealerCard != "10")
}

func isFace(card string) bool {
    switch card {
        case "jack", "queen", "king":
        	return true
    }
    return false
}
