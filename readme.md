# MQTT Broker using Docker Compose

This mail contains a Docker Compose file (`docker-compose.yaml`) and a Mosquitto configuration file (`mosquitto.conf`) to set up an MQTT broker using Docker.

## Prerequisites

Before you begin, ensure that you have the following installed on your machine:

- Docker
- Docker Compose

## Setup Instructions

1. **Download the Files:**
   - Download the `docker-compose.yaml` and `mosquitto.conf` files into a directory of your choice.

2. **Navigate to the Directory:**
   - Open a terminal and navigate to the directory where you saved the `docker-compose.yaml` and `mosquitto.conf` files.

3. **Start the MQTT Broker:**
   - Run the following command in the terminal:
     ```bash
     docker-compose up
     ```
   This command will build the necessary Docker containers and start the MQTT broker.

4. **Access the MQTT Broker:**
   - Once the broker is up and running, you can access it using the default MQTT port (1883) on your localhost or the IP address of your machine.

5. **Stopping the Broker:**
   - To stop the broker and remove the Docker containers, press `Ctrl + C` in the terminal where the `docker-compose up` command is running.


# Run The Application 

1. **Run Publisher & Subscriber:**
   - To run publisher `go run publisher.go constants.go`
   - To run subscriber `go run subscriber.go constants.go`