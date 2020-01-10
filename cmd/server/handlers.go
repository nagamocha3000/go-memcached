package main

import "net/http"

import "strconv"

import "fmt"

func (api *httpAPI) home(w http.ResponseWriter, r *http.Request) {
	var err error = nil
	if err != nil {
		api.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	w.Write([]byte("Hello go-memcached"))
}

func (api *httpAPI) getValue(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":key"))
	if err != nil {
		api.notFound(w)
		return
	}
	w.Write([]byte(fmt.Sprintf("val for key %d", id)))
}

func (api *httpAPI) handleSet(w http.ResponseWriter, r *http.Request) {
	var key, val, exptimeStr string
	reply := api.cache.Set(key, val, exptimeStr)
	//return only reply
	w.Write([]byte(reply))
}

func (api *httpAPI) handleAdd(w http.ResponseWriter, r *http.Request) {
	var key, val, exptimeStr string
	reply := api.cache.Add(key, val, exptimeStr)
	//return only reply
	w.Write([]byte(reply))
}

func (api *httpAPI) handleReplace(w http.ResponseWriter, r *http.Request) {
	var key, val, exptimeStr string
	reply := api.cache.Replace(key, val, exptimeStr)
	//return only reply
	w.Write([]byte(reply))
}

func (api *httpAPI) handleAppend(w http.ResponseWriter, r *http.Request) {
	var key, val, exptimeStr string
	reply := api.cache.Append(key, val, exptimeStr)
	//return only reply
	w.Write([]byte(reply))
}

func (api *httpAPI) handlePrepend(w http.ResponseWriter, r *http.Request) {
	var key, val, exptimeStr string
	reply := api.cache.Prepend(key, val, exptimeStr)
	//return only reply
	w.Write([]byte(reply))
}

func (api *httpAPI) handleIncrement(w http.ResponseWriter, r *http.Request) {
	var key, numStr string
	reply := api.cache.Increment(key, numStr)
	w.Write([]byte(reply))
}

func (api *httpAPI) handleDecrement(w http.ResponseWriter, r *http.Request) {
	var key, numStr string
	reply := api.cache.Decrement(key, numStr)
	w.Write([]byte(reply))
}

func (api *httpAPI) handleCompareAndSwap(w http.ResponseWriter, r *http.Request) {
	var key, val, exptimeStr string
	reply := api.cache.CompareAndSwap(key, val, exptimeStr, "")
	w.Write([]byte(reply))
}

func (api *httpAPI) handleGet(w http.ResponseWriter, r *http.Request) {
	var key string
	val, reply := api.cache.Get(key)
	w.Write([]byte(reply))
	w.Write([]byte(val))
}

func (api *httpAPI) handleGetEntryPlusToken(w http.ResponseWriter, r *http.Request) {
	var key string
	val, reply, token := api.cache.GetEntryPlusToken(key)
	w.Write([]byte(reply))
	w.Write([]byte(val))
	w.Write([]byte(token))
}

func (api *httpAPI) handleDelete(w http.ResponseWriter, r *http.Request) {
	var key string
	reply := api.cache.Delete(key)
	w.Write([]byte(reply))
}

func (api *httpAPI) handleClear(w http.ResponseWriter, r *http.Request) {
	reply := api.cache.Clear()
	w.Write([]byte(reply))
}

func (api *httpAPI) handleStats(w http.ResponseWriter, r *http.Request) {
	reply := api.cache.Stats()
	w.Write([]byte(reply))
}
