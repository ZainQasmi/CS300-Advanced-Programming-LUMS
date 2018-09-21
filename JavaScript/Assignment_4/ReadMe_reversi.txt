Part 1: Every "two" consecutive connections are paired together as a game. Turns implemented as a list of booleans on server side so clients can't cheat.
Valid Moves implemented. This part is complete.

Part 2: Only checks to see if all squares are occupied and player with higher count wins. 
DOES NOT CHECK "that the game can end before all boxes are filled because one player is left with no valid move"

Part 3: NOT implemented.

Part 4: Every game session is an 8x8 array matrix. And every such matrix (i.e. game session) is stored inside a list (gamesList) which is updated (only the specific relevant matrix in that list) whenever either of the corresponding player makes a move. 