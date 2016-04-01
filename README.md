# test_consul
test consul for service register/discovery

# Running a real Consul cluster in a production environment
```
docker run -d -h node1 -v /mnt:/data \
    -p 10.0.1.1:8300:8300 \
    -p 10.0.1.1:8301:8301 \
    -p 10.0.1.1:8301:8301/udp \
    -p 10.0.1.1:8302:8302 \
    -p 10.0.1.1:8302:8302/udp \
    -p 10.0.1.1:8400:8400 \
    -p 10.0.1.1:8500:8500 \
    -p 172.17.42.1:53:53/udp \
    progrium/consul -server -advertise 10.0.1.1 -bootstrap-expect 3
```
```
docker run -d -h node2 -v /mnt:/data  \
    -p 10.0.1.2:8300:8300 \
    -p 10.0.1.2:8301:8301 \
    -p 10.0.1.2:8301:8301/udp \
    -p 10.0.1.2:8302:8302 \
    -p 10.0.1.2:8302:8302/udp \
    -p 10.0.1.2:8400:8400 \
    -p 10.0.1.2:8500:8500 \
    -p 172.17.42.1:53:53/udp \
    progrium/consul -server -advertise 10.0.1.2 -join 10.0.1.1
```
```
docker run -d -h node3 -v /mnt:/data  \
    -p 10.0.1.3:8300:8300 \
    -p 10.0.1.3:8301:8301 \
    -p 10.0.1.3:8301:8301/udp \
    -p 10.0.1.3:8302:8302 \
    -p 10.0.1.3:8302:8302/udp \
    -p 10.0.1.3:8400:8400 \
    -p 10.0.1.3:8500:8500 \
    -p 172.17.42.1:53:53/udp \
    progrium/consul -server -advertise 10.0.1.3 -join 10.0.1.1
```
