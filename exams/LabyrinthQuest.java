/*
Labyrinth Explorer's Quest

In a magical labyrinth, an adventurer named Max ventures forth. The labyrinth comprises rooms interconnected by unidirectional doors. Each time Max traverses a door, the labyrinth's keeper rewards with a coin.

The objective is to maximize the score. However, a score of 0 is incurred when not using any doors. Max can traverse through the doors multiple times.

Given the number of rooms 'n' and doors 'm', along with their descriptions (door connections from one room to another with the associated coin value), assist Max in determining the maximum score achievable.

Input Format:

The first line contains two integers 'n' and 'm', denoting the number of rooms and doors respectively.
The next 'm' lines describe the doors. Each line contains three integers 'v', 'u', and 'w':
'v' represents the room Max is currently in.
'u' represents the room Max can reach through the door.
'w' denotes the coin value received when moving from room 'v' to room 'u' through this door.

Constraints:

2 ≤ n ≤ 3 × 10^5
1 ≤ m ≤ 3 × 10^5
1 ≤ v, u ≤ n
0 ≤ w ≤ 10^9

Output Format:

Output a single integer representing the maximum score Max can achieve.

Sample Input:

6 8
1 2 9
2 3 7
3 1 1
2 4 5
3 4 4
4 5 6
5 6 3
6 4 10

Sample Output:

9

Explanation:

One of the optimal paths is: 2 -> 3 -> 1 -> 2 -> 4 -> 5 -> 6 -> 4, resulting in a score of 10 - 1 = 9.
*/
