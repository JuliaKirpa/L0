function onLoaded() {
    var source = new EventSource("http://localhost:8080/as");
    source.onmessage = function (event) {
        console.dir(event);
        document.getElementById("counter").innerHTML = event.data;
    }
}

