package observer

import "strconv"

type WeatherStation struct {
	subscribers []*Subscriber
	temprature  float32
}

func (ws *WeatherStation) Subsribe(subscriber Subscriber) {
	ws.subscribers = append(ws.subscribers, &subscriber)
}
func (ws *WeatherStation) Unsubscribe(subscriber Subscriber) {
	for i, v := range ws.subscribers {
		if *v == subscriber {
			ws.subscribers = append(ws.subscribers[:i], ws.subscribers[i+1:]...)
			return
		}
	}
}
func (ws *WeatherStation) Notify() {

	temp := strconv.FormatFloat(float64(ws.temprature), 'f', -1, 32)
	for _, s := range ws.subscribers {
		(*s).update(temp)
		(*s).display(temp)
	}

}
func (ws *WeatherStation) GetTemprature() float32 {
	return ws.temprature
}

func (ws *WeatherStation) SetTemprature(temp float32) {
	ws.temprature = temp
}
