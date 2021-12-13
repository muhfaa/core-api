
# Core Data Collaction



## Indices

* [Order](#order)

  * [Insert Order](#1-insert-order)
  * [Update Order](#2-update-order)


--------


## Order



### 1. Insert Order



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:7171/v1/order/add
```



***Body:***

```js        
{
    "kerusakan_id": 2,
    "jenis_hp": "iphone",
    "jenis_platform": "ios"
}
```



### 2. Update Order



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:7171/v1/order/update-status
```



***Body:***

```js        
{
    "id": 1,
    "version": 1
}
```

