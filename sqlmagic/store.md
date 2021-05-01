
## Add Product
```json
curl --header "Content-Type: application/json"   --request POST   --data '{
"categoryID"    :  1,    
"ownerID"       :  1,
"productname"   : "Product name", 
"image"         : "https://burgerfarms.com/wp-content/gallery/fertilizers-plant-food/Organic-Fertilizers.JPG",
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
"ownerID"       :  1,
"productID"     :  12,
"categoryID"    :  1,    
"productname"   : "Another New Name", 
"image"         : "http://image",
"imagelarge"    : "http://large",
"description"   : "this descrioption is too long",
"packingtype"   : "bottles",
"stock"         :  100,
"price"         :  2000.34
}' http://localhost:3000/store/update

```

## Get All Product By category
```json
curl --header "Content-Type: application/json"   --request POST   --data '{

"categoryid"       :  1

}' http://localhost:3000/store/products

```

## Get Current Owners STOCK
```json
curl --header "Content-Type: application/json"   --request POST   --data '{

"ownerid"       :  1

}' http://localhost:3000/store/stock

```

## Get All Categories 
```json
curl --header "Content-Type: application/json"   http://localhost:3000/store/categories

```



