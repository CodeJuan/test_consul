docker run -d --name server1 -h server1 progrium/consul -server -bootstrap-expect 2

docker run -d --name server2 -h server2 progrium/consul -server -join $JOIN_IP