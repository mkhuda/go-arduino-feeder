package models

// Movie is used to represent movie profile data
type Temperature struct {
	DeviceSerial string `json:"deviceSerial" bson:"deviceserial"`
	Temperature  int    `json:"temperature" bson:"temp"`
	Humidity     int    `json:"humidity" bson:"hum"`
	Heat         int    `json:"heat" bson:"heat"`
	CreatedOn    string `json:"createdOn" bson:"createdon"`
}
