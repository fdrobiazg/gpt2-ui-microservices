var api = "http://127.0.0.1:3250"

$(document).ready(function() {
    $("#btn").click(function() {
        $.ajax({
            url: api + "/api/generateText",
            type: "GET",
            dataType: "text/plain",
            timeout: 3000,
            success: function(data) {
                $("#quote").html("Loading...") 
                $("#quote").addClass('is-link')
                $( "#quote" ).html(data + '</br><b>'); 
            },
            error: function(xmlhttprequest, textstatus, message) {
                $("#quote").removeClass('is-link')
                $("#quote").addClass('is-danger')
                if(textstatus==="timeout") {
                    $( "#quote" ).html("got timeout");
                } else {
                    $( "#quote" ).html(message);
                }
            }
        })
    })
})