# Problem

1. Only whole packs can be sent. Packs cannot be open. 
2. Given rule 1. Send out no more items than necessary to fulfil the order.
3. Given rules 1 & 2 above, send out as few packs as possible to fulfill each order.

The following pack sizes we offer are:

- 250 items
- 500 items
- 1000 items
- 2000 items
- 5000 items

## Approaches

### Greedy Approach

If you sort the pack sizes by largest first and then iterate through them, selecting the largest pack that fits within the remaining order, you'll get a solution that works for many cases. This is known as the greedy approach. Here's how it works:

1. Sort the pack sizes in descending order: [5000, 2000, 1000, 500, 250]
2. Start with the full order amount
3. For each pack size:
   - If the pack size is less than or equal to the remaining order:
     - Add this pack to the solution
     - Subtract the pack size from the remaining order
   - Repeat until the remaining order is 0

This approach is simple and efficient, working well for many scenarios. However, it doesn't always produce the optimal solution - since it only optimises for the least amount of items, not the fewest amount of total packs. For example:

- For an order of 251 items, the greedy approach would select [250, 250], using 2 packs
- The optimal solution is [500, 1], using only 1 pack

Despite this limitation, the greedy approach can be a good starting point or a quick approximation method. 


