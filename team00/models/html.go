package models

type SessionsModel struct {
	Count    int64
	Sessions []Session
}

type AnomalyModel struct {
	Count     int64
	Session   Session
	Anomalies []Anomaly
}
