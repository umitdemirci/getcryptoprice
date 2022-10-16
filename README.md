# Get Crypto Price
Get crypto price from Binance via Binance API.(For now just supported currency is USDT)

# Build
```bash
go build .
```

# Usage

```bash
# usage 
getcryptoprice BTC
# output
# BTC price is 19137.36 USDT
```

Changing the name of the executable file for easy use: 
```bash
mv getcryptoprice price
```

### TODO:
- [ ] support other vendors
- [ ] support other currencies