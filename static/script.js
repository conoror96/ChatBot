const conversation = $("#list"); // $ means jquery, it won't get the item without the #
const userInput = $("#userInput");
console.log(conversation);


userInput.keypress(function (event) {
    if (event.keyCode !== 13) { // User presses enter
        return; 
    }
    event.preventDefault(); // prevents page refresh
    const text = userInput.val(); 

    console.log(text);
    userInput.val(""); // set userInput to nothing

    if (text.trim() == "") { //if statement for trimming spaces
        return;
    }

    queryParameters = {     // a query parameter user-input is expected
        "userInput": text
    }

    conversation.append("<li id='inputBox' align='right' class=\"list-group\">" + text + "<li class=\"list-group\">");
    //http://localhost:8080
    $.get("/ask", queryParameters).done(function (resp) {
        // this code will execute when the request gets a response.
        setTimeout(function () { // wait 1 second then add element.
            conversation.append("<li id='outputBox' align='left' class=\"list-group\">" + resp + "<li class=\"list-group\">");
        }, 1000);

    }).fail(function () { // error / fail message 
        conversation.append("<li class=\"list-group\">The bot Can't talk right now!</li class=\"list-group\">");
    });

    window.scrollTo(0, document.body.scrollHeight);
});
