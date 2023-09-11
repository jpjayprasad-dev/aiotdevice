# AIOT Sensor/Device

This is a golang application which mocks data points that simulate the IoT sensors. These simulators will read data from the provided CSV file and make them available via an HTTP API.
Also the application allows receiving controls to the controllers of IoT devices via HTTP POST.

## Application Specs

Our objective is to build a light weight go lang app which mocks the behaviour of IoT sensors or devices installed in the room.

1. IoT sensor data in csv format will be stored into (./data/) directory
2. The room configuration will be stored into ./config.yaml. The room configuration file will have the following format
```
room:
  id: <unique_id of the room>
  host: <specify the ip address which the device/sensor runs>
  port: <port which the device/sensor runs>
  infile: <path to the iot simulation data>
  outfile: <path to store the control data recieved>
```
## Build Service
1. Clone this git repo.
```
git clone https://github.com/jpjayprasad-dev/aiotdevice/
```
2. Build go application
```
go build
```
## Run Service
```
./aiotdevice
```
## Usage
1. To get sensor/device data
   ```
   curl -X GET http://<host>:<port>/<device_id>/data
   ```
2. To post controls to the device controller
```
curl -X POST -d '{"room_id" : <unique_id_of_the_room>, "device_id" : <type_of_the_device>, "controlpoint" : <control_parameter>, "value": <control_value>}' http://<host>:<port>/control
```

## Examples [Screenshots]
To get data from sensor of type life_being

<img width="1216" alt="image" src="https://github.com/jpjayprasad-dev/aiotdevice/assets/73153441/96007c53-26af-45ba-8ba0-9ae046182a89">

To post temperature controls to the controller of the device type aircon

<img width="1501" alt="image" src="https://github.com/jpjayprasad-dev/aiotdevice/assets/73153441/bd10c311-2f2b-4901-bd54-61bef2d5ba55">

