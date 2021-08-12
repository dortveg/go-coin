# somename

<> is a terminal cryptocurrency price and volume ticker/tracker written in Go!

## Usage

Simply type in your desired coin pairs without any spaces, dashes or slashes (ex. "BTCUSDT") and turn it on.\
<> will show the current price and 24hr change as a ticker. It will also update and show the volume flow for every set interval (default 1 minute), and compare it to the average volume flow for the past 3 hours.\
Whenever the current volume flow is higher than the average volume flow by your desired amount (default 100%), <> will play a sound alert and highlight the coin.\
\
Currently it receives the price and volume data from Binance.com.

## License
[GNU AGPL](https://choosealicense.com/licenses/agpl-3.0/)
