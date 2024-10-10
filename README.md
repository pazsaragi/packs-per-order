# Packs Per Order

> We want to build an API and UI that allows customers to place orders of any size. We are constrained by the following rules:

1. Only whole packs can be sent. Packs cannot be open. 
2. Given rule 1. Send out no more items than necessary to fulfil the order.
3. Given rules 1 & 2 above, send out as few packs as possible to fulfill each order.

The following pack sices we offer are:

- 250 items
- 500 items
- 1000 items
- 2000 items
- 5000 items

## Examples

| items ordered | Correct number of packs | Incorrect number of packs |
| --- | --- | --- |
| 1 | 1 x 250 | 1 x 500 - more items than necessary |
| 250 | 1 x 250 | 1 x 500 - more items than necessary |
| 251 | 1 x 500 | 2 x 250 - more packs than necessary |
| 501 | 1 x 500, 1 x 250 | 1 x 1000 - more packs than necessary, 3 x 250 more packs than necessary |
| 12001 | 2 x 5000, 1 x 2000, 1 x 250 | 3 x 5000 - more items than necessary |

This application should calculate the ideal number of packs to ship to the customer. 

**Note:** Optionally, the application should be extensible and should still work if packs can be added or removed.