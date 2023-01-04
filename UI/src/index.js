var url = "http://127.0.0.1:3250/api/generateText"

function generateText() {
    document.getElementById("text").className = "notification";
    document.getElementById("text").innerHTML = "Loading...";
    var text = document.getElementById("input").value
    console.log(text)
    const Http = new XMLHttpRequest();
    Http.open("POST", url);
    Http.send(text);
    Http.onreadystatechange = (e) => {
        document.getElementById("text").innerHTML = Http.responseText
        console.log(Http.responseText)
    }
}