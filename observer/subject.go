package main

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type WeatherData struct {
	observers   []Observer
	temperature float64
	Humidity    float64
	Pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	var indexToRemove int
	// we search for an element with an index.
	for i, observer := range w.observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	w.observers = append(w.observers[:indexToRemove], w.observers[indexToRemove+1:]...)
}

func (w *WeatherData) NotifyObservers() {
	for _, o := range w.observers {
		o.Update(w.temperature, w.Humidity, w.Pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurement(temp, humidity, pressure float64) {
	w.temperature = temp
	w.Humidity = humidity
	w.Pressure = pressure

	w.MeasurementsChanged()
}
