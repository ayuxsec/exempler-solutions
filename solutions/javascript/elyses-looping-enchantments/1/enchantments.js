// @ts-check

/**
 * Determine how many cards of a certain type there are in the deck
 *
 * @param {number[]} stack
 * @param {number} card
 *
 * @returns {number} number of cards of a single type there are in the deck
 */
export function cardTypeCheck(stack, card) {
  let num = 0;
  stack.forEach((c) => {
    if (c == card) {
      num += 1
    }
  });
  return num
}

/**
 * Determine how many cards are odd or even
 *
 * @param {number[]} stack
 * @param {boolean} type the type of value to check for - odd or even
 * @returns {number} number of cards that are either odd or even (depending on `type`)
 */
export function determineOddEvenCards(stack, type) {
  let numOdd = 0;
  let numEven = 0;
  
  stack.forEach((number) => {
    if (number % 2 == 0) {
      numEven += 1;
    } else {
      numOdd += 1;
    }
  });

  if (type) {
    return numEven
  } else {
    return numOdd
  }
}
