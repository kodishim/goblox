# 🧱 Goblox

A simple & effective wrapper for the Roblox API, Rolimons API and Rblx.Trade API.

## 🔍 Testing

All tests rely on external APIs. As a result test data is needed.

### 🦸 rouser

#### .env

```
COOKIE=[.ROBLOSECURITY Cookie]
SECRET=[Roblox TFA Secret]
```

#### data/trades.json

```
{
  "testTrade": {
    "offers": [
      {
        "userId": 25277066,
        "userAssetIds": [48918241760],
        "robux": 0
      },
      {
        "userId": 560349464,
        "userAssetIds": [29791249875],
        "robux": 10
      }
    ]
  }
}
```

### 🌐 roscraper

#### data/proxies.txt

```
USER:PASS@HOST:PORT
```

### 🤝 rolimons

#### .env

```
COOKIE=[Rolimons Cookie]
```

#### data/rolimons.json

```
{
  "userId": 25277066,
  "testTradeAd": {
    "offer": [55610781],
    "request": [1048037],
    "requestTags": ["demand"]
  }
}
```
