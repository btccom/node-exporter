##  Example

### http



### stratum server

```
# TYPE btc_stratum_server gauge
btc_stratum_server_height{pool="f2pool.com", username="dubuqingfeng", last_hash=""} 650000
```


### public peer

```
# TYPE btc_public_node gauge
btc_public_node_height{address="f2pool.com", coin="btc", last_hash=""} 650000
```