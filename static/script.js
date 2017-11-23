const conversation = $("#list"); 
const userInput = $("#userInput");
console.log(conversation);


userInput.keypress(function (event) {
    if (event.keyCode !== 13) { 
        return; 
    }
    event.preventDefault();
    const text = userInput.val(); 

    console.log(text);
    userInput.val(""); 
    if (text.trim() == "") { 
        return;
    }

    queryParameters = {    
        "userInput": text
    }

 conversation.append("<li id='inputBox' align='right' class=\"list-group\">" + text + "<li class=\"list-group\">");
 $.get("/ask", queryParameters).done(function (resp) {
       
        setTimeout(function () { 
            conversation.append("<li id='outputBox' align='left' class=\"list-group\">" + resp + "<li class=\"list-group\">");
        }, 1000);

    }).fail(function () {
        conversation.append("<li class=\"list-group\">the bot Can't talk right now!</li class=\"list-group\">");
    });

    window.scrollTo(0, document.body.scrollHeight);
});
