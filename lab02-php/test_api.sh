#!/bin/bash

BASE_URL="http://localhost:8081"

echo "=== Testing Products ==="
echo "Creating product..."
curl -X POST $BASE_URL/products -H "Content-Type: application/json" -d '{"name":"Laptop","price":"999.99"}'
echo -e "\nGetting all products..."
curl $BASE_URL/products
echo -e "\nUpdating product 1..."
curl -X PUT $BASE_URL/products/1 -H "Content-Type: application/json" -d '{"name":"Gaming Laptop","price":"1299.99"}'
echo -e "\nGetting product 1 after update..."
curl $BASE_URL/products/1
echo -e "\nDeleting product 1..."
curl -X DELETE $BASE_URL/products/1
echo -e "\nGetting all products after delete..."
curl $BASE_URL/products

echo -e "\n\n=== Testing Categories ==="
echo "Creating category..."
curl -X POST $BASE_URL/categories -H "Content-Type: application/json" -d '{"name":"Electronics","description":"Devices"}'
echo -e "\nGetting all categories..."
curl $BASE_URL/categories
echo -e "\nUpdating category 1..."
curl -X PUT $BASE_URL/categories/1 -H "Content-Type: application/json" -d '{"name":"Premium Electronics","description":"High-end devices"}'
echo -e "\nGetting category 1 after update..."
curl $BASE_URL/categories/1
echo -e "\nDeleting category 1..."
curl -X DELETE $BASE_URL/categories/1
echo -e "\nGetting all categories after delete..."
curl $BASE_URL/categories

echo -e "\n\n=== Testing Orders ==="
echo "Creating order..."
curl -X POST $BASE_URL/orders -H "Content-Type: application/json" -d '{"customer_name":"John","total_amount":"1000.00","status":"pending"}'
echo -e "\nGetting all orders..."
curl $BASE_URL/orders
echo -e "\nUpdating order 1..."
curl -X PUT $BASE_URL/orders/1 -H "Content-Type: application/json" -d '{"customer_name":"John","total_amount":"1000.00","status":"shipped"}'
echo -e "\nGetting order 1 after update..."
curl $BASE_URL/orders/1
echo -e "\nDeleting order 1..."
curl -X DELETE $BASE_URL/orders/1
echo -e "\nGetting all orders after delete..."
curl $BASE_URL/orders

echo -e "\n\n=== Test Complete ==="
