window.addEventListener('load', (event) => {
    nextLink = document.getElementById("next-button");
    nextLink.addEventListener('click', ()=> {
        myFunction();
    })
});

function myFunction() {
    radioButtonGroup = document.getElementsByName("btnradio");
    for (let i = 0; i < radioButtonGroup.length; i++) {
        if (radioButtonGroup[i].checked == true) {
            var selectedRadioButton = radioButtonGroup[i];
            result = selectedRadioButton.value
            makeRequest(result);
        };
    }
}

function makeRequest(result){
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            // Typical action to be performed when the document is ready:
            alert("response is 200");
        }
    };
    xhttp.open("POST", "http://127.0.0.1:8080/addResult/"+result, true);
    console.log("what's in xhttp, ", xhttp);
    xhttp.send();
}

