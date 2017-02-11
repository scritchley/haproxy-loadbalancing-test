# HAProxy Custom Load Balancing

Start with:

    docker-compose up

Test by making a POST using cURL. 

    curl -XPOST -d '{"value":"123"}' http://127.0.0.1:5000

The backend server should be selected based on the value passed in the 'value' property of the JSON request body.
