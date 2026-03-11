package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/health", healthCheck)

	http.ListenAndServe("0.0.0.0:8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {

	hostname, _ := os.Hostname()
	port := os.Getenv("PORT")
	instance := os.Getenv("INSTANCE_NUMBER")
	now := time.Now().Format(time.RFC1123)

	fmt.Fprintf(
		w, 
<!DOCTYPE html>
<html>
<head>
<title>Go on Clever Cloud</title>
<style>
body {
	font-family: Arial, sans-serif;
	background:#0f172a;
	color:white;
	margin:0;
	padding:40px;
}

.container{
	max-width:900px;
	margin:auto;
}

.card{
	background:#1e293b;
	padding:25px;
	border-radius:10px;
	margin-bottom:20px;
}

h1{
	color:#22c55e;
}

.grid{
	display:grid;
	grid-template-columns:1fr 1fr;
	gap:15px;
}

.box{
	background:#020617;
	padding:15px;
	border-radius:6px;
	font-family:monospace;
}

a{
	color:#38bdf8;
}
</style>
</head>

<body>

<div class="container">

<div class="card">
<h1>🚀 Go running on Clever Cloud</h1>
<p>This application is deployed automatically on Clever Cloud.</p>
</div>

<div class="card">
<h2>Runtime information</h2>

<div class="grid">

<div class="box">
Hostname<br>
<b>%s</b>
</div>

<div class="box">
Instance<br>
<b>%s</b>
</div>

<div class="box">
Port<br>
<b>%s</b>
</div>

<div class="box">
Server Time<br>
<b>%s</b>
</div>

</div>

</div>

<div class="card">
<h2>Healthcheck</h2>
<p>Endpoint available at:</p>
<p><a href="/health">/health</a></p>
</div>

<div class="card">
<p>Try it yourself on <a href="https://www.clever.cloud">Clever Cloud</a> ☁️</p>
</div>

</div>

</body>
</html>
,
		hostname,
		instance,
		port,
		now,
	)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
