# Assignment 2 - Scalable Web Service with Golang - Hacktiv8 - FGA Digitalent

## Build Rest API in GO

By: Muhammad Agi Sahriza Daan Nur (149368582101-159)

### List API

1. Create Orders

   - Method: POST
   - URL: /orders
   - Body Request:

   ```json
   {
     "customerName": "Agi Sahriza",
     "items": [
       {
         "itemCode": "i1",
         "description": "Iphone XR",
         "quantity": 1
       },
       {
         "itemCode": "i2",
         "description": "Iphone 11",
         "quantity": 2
       }
     ]
   }
   ```

2. List Orders

   - Method: GET
   - URL: /orders

3. Update Order

   - Method: PUT
   - URL: /orders/{orderID}
   - Body Request:

   ```json
   {
     "customerName": "Agi Corbuzer",
     "items": [
       {
         "itemCode": "i1",
         "description": "Iphone XR",
         "quantity": 10
       },
       {
         "itemCode": "i2",
         "description": "Iphone 11",
         "quantity": 20
       }
     ]
   }
   ```

4. Delete Order
   - Method: Delete
   - URL: /orders/{orderID}
