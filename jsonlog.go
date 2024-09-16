package toolkits

/*
	Logging examples
	----------------
	1. config.CLog("VISIT_INDEX_PAGE", map[string]interface{}{"data1": "test1", "data2": "test2"}).Info("Visit Index Page")
	2. config.CLog("INDEX_WARNING", map[string]interface{}{"data1": "t1", "data2": "t2"}).Warning("Index warning")
	3. config.CLog("INDEX_ERROR", map[string]interface{}{"data1": "t1", "data2": "t2"}).Error("Index Error")
*/

import (
	log "github.com/sirupsen/logrus"
)

// CLog to display JSON log in termanl
func CLog(actionCode string, context map[string]interface{}) *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})

	data := log.Fields{
		"action":  actionCode,
		"context": context,
	}

	return log.WithFields(data)
}
