# discards
pokemontcg.io data caching server

### Required
- API key from https://dev.pokemontcg.io/

### Usage
```
git clone https://github.com/pidgy/discards.git
cd discards
cat <pokemontcg.io-api-key> .api
go run main.go
```

#### GET random card 
```
curl localhost:8080/card
```

#### GET specific card
```
curl localhost:8080/card?id=swsh8-3
```

#### GET all aets
```
curl localhost:8080/sets
```
