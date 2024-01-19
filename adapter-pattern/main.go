package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdaptor struct {
	windows *Windows
}

func (w *WindowsAdaptor) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windows.insertIntoUSBPort()
}

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsAdaptor := &WindowsAdaptor{
		windows: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsAdaptor)
	fmt.Println("---------------------------------------------")
	mySQL := &MySQL{}
	dbClient := &dbClient{}

	dbClient.EstablishDBConnection(mySQL)

	mongo := &MongoDB{}
	mongoAdptor := &MongoDBAdaptor{
		mongoDB: mongo,
	}

	dbClient.EstablishDBConnection(mongoAdptor)

}

type Database interface {
	ConnectDB()
}

type MySQL struct{}

func (m *MySQL) ConnectDB() {
	fmt.Println("Connected to MySQL DB.")
}

type MongoDB struct{}

func (m *MongoDB) connectToMongoDB() {
	fmt.Println("Connected to MONGO DB.")
}

type MongoDBAdaptor struct {
	mongoDB *MongoDB
}

func (m *MongoDBAdaptor) ConnectDB() {
	fmt.Println("Converting MYSQL connection to Mongo connection. ")
	m.mongoDB.connectToMongoDB()
}

type dbClient struct{}

func (c *dbClient) EstablishDBConnection(db Database) {
	fmt.Println("Client is trying to connect to MySQL Database.")
	db.ConnectDB()
}
