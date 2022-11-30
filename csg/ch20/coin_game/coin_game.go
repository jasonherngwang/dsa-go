/*
Problem
------------------------------------------
Coin Game
- Pile of coins
- Two players, taking turns
- Each player must take 1 or 2 coins on their turn.
- If a player takes the last coin, they lose.

Inputs:
- Number of coins in pile (integer)
- Current player
Outputs: Winning player (Current or Opponent)

Rules, Requirements, Definitions
------------------------------------------
- At least 1 coin in starting pile
- Cannot skip turn

Examples, Test Cases, Edge Cases
------------------------------------------
Case: 1 coin in pile
- Input: 1, Current
- Output: Opponent wins, since Current must take 1 coin

Case: 2 coins in pile
- Input: 2, Current
- Output: Current wins, since Current can take 1 coin and force Opponent to take
  the last coin (and lose)

Case: 3 coins in pile
- Input: 3, Current
- Output: Current wins, since Current can take 2 coins and force Opponent to
  take the last coin (and lose)

Case: 4 coins in pile
- Input: 4 or greater, Current
- Output: Opponent wins.
  - If Current takes 1, Opponent takes 2 and forces Current to lose.
  - If Current takes 2, Opponent takes 1 and forces Current to lose.

Data Structure
------------------------------------------
None

Algorithm
------------------------------------------
Approach: Recursive
Time: O(N^2) since each function call makes 2 recursive calls
Space: O(N^2) calls on call stack

Steps
- Base case: <=0 coins remaining (can be negative if we start with 1 and take 2)
	- The opponent has taken all remaining coins, so current player wins.
- Else
  - Determine who goes next, using strings "you" or "them".
  - The current player can take 1 or 2 coins. Model these scenarios by
    performing a recursive call using args coins-1 and coins-2, checking if
	either returns a win for the current player.
  - If so, return current player's name. Else return opponent's name.


Approach 2: Finding a Pattern
Time: O(1)
Space: O(1)

Coins Winner
1     Them
2     You
3     You
4     Them
5     You
6     You
7     Them

If numCoins - 1 is divisible by 3, winner is "Them".
*/

package csg

func coingameWinnerRecursive(numCoins int, player string) string {
	if numCoins <= 0 {
		return player
	}

	nextPlayer := "them"
	if player == "them" {
		nextPlayer = "you"
	}

	if coingameWinnerRecursive(numCoins-1, nextPlayer) == player || coingameWinnerRecursive(numCoins-2, nextPlayer) == player {
		return player
	} else {
		return nextPlayer
	}
}

func coingameWinner(numCoins int, player string) string {
	if (numCoins-1)%3 == 0 {
		return "them"
	}
	return "you"
}
