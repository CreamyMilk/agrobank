
## Add Product
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
"categoryID"    :  1,    
"productname"   : "pridname", 
"image"         : "http://image",
"imagelarge"    : "http://large",
"description"   : "this descrioption is too long",
"packingtype"   : "bottles",
"stock"         :  100,
"price"         :  2000.34
}' http://localhost:3000/store/add

```

## Edit Product
```json
curl --header "Content-Type: application/json"   --request PUT   --data '{
"productID"     :  10,
"categoryID"    :  1,    
"productname"   : "newName", 
"image"         : "http://image",
"imagelarge"    : "http://large",
"description"   : "this descrioption is too long",
"packingtype"   : "bottles",
"stock"         :  100,
"price"         :  2000.34
}' http://localhost:3000/store/update

```