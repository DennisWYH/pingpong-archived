window.addEventListener('load', (event) => {
    nextLink = document.getElementById("next-button");
    nextLink.addEventListener('click', ()=> {
        addResultRequest();
    })
});

function addResultRequest() {
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
        }
    };
    xhttp.open("POST", "/addResult", true);
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    if (result==='flash'){
        result=0;
    }
    if (result==='done'){
        result=1;
    }
    if (result==='hard'){
        result=2;
    }
    const json = {
        "user": 1,
        "result": result,
        "sentence": 1,
    };
    xhttp.send(JSON.stringify(json));
}
