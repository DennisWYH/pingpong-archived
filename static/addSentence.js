function addNewSentence() {
    var chinese = document.getElementById("sentence-input-chinese").value;
    var english = document.getElementById("sentence-input-english").value;

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            // Typical action to be performed when the document is ready:
        }
    };
    xhttp.open("POST", "http://localhost:8080/addSentence", true);
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    var sentence = {"chinese":chinese, "pinyin": "", "english": english}
    xhttp.send(JSON.stringify(sentence));
}


