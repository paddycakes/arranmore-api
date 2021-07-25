package danalto

type SensorDataList []SensorDataPoint

type SensorDataPoint struct {
	Timestamp   int64    `json:"timestamp"`
	Sf          string   `json:"sf"`
	Rssi        []float32    `json:"rssi"`
	Gateways    []string `json:"gateways"`
	MaxRssi     string   `json:"max_rssi"`
	PrefGateway string   `json:"pref_gateway"`
	Payload     struct {
		Temperature string `json:"temperature"`
		Humidity    string `json:"humidity"`
		Motion      string `json:"motion"`
		Battery     string `json:"battery"`
		Light       string `json:"light"`
	} `json:"payload"`
}


// TODO: This should be camelCase. But when sent via api needs to be snake_case.

type DeviceDataOpts struct {
	StartTimestamp int64
	EndTimestamp int64
}
