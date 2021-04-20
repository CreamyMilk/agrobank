## Deposit Funds
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
"walletname"        :"Peter",    
"phonenumber"       :"254797678252", 
"fcmtoken"          :"FCMTOKEN",
"amount"            :"10"
}' http://localhost:3000/wallet/deposit

```

## Send Money
```json
curl --header "Content-Type: application/json"   --request POST   --data '
{
"from"     : "JOB",
"to"       : "ALICE", 
"amount"   : 10
}' http://localhost:3000/wallet/sendmoney

```

## Get Balance 

```json
curl --header "Content-Type: application/json"   --request POST   --data '
{
"walletname"     : "JOB"
}' http://localhost:3000/wallet/balance

```

## Verify receiver

```json
curl --header "Content-Type: application/json"   --request POST   --data '
{
"phonenumber"     : "254797678251",
"amount"          : 10000
}' http://localhost:3000/wallet/verify

```


## Get Transactions

```json
curl --header "Content-Type: application/json"   --request POST   --data '
{
"walletname"     : "254797678251"
}' http://localhost:3000/wallet/transactions

```

