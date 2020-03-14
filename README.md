# Tyme
The library to achieve type safe datetime.

# Feature
## Datetime with resolution
Tyme provides datetime with resolution.
You can declare datetime without hour, minutes, etc.

For example, `NewLocalDate(2006, time.January, 02)` returns January 2, 2006.
It is NOT 00:00:00 January 2, 2006.
It is JUST January 2, 2006.

