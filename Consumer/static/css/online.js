const source = new EventSource("http://localhost:8080/rows/")
source.onmessage = (event) => {
    console.log("OnMessage Called:")
    console.log(event)
    console.log(JSON.parse(event.data))
}
