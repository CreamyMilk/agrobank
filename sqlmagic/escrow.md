
## Add Product
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
"walletname"    :  "JOB",    
"productid"     :  0,    
"quantity"      :  1,    
"passwordHash"  :  "3423432",
"acceptDelivery":  0
}' http://localhost:3000/invoice/create

```
