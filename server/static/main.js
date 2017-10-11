function step1() {
    sendStep1();
    //fadeToNewStep("step-1", "step-2");
}

function fadeToNewStep(from, to) {
    $("#" + from).fadeOut(
        function()
        {
            $("#" + to).fadeIn();
        }
    );
}

function sendStep1() {
    $.ajax({
        type: "POST",
        url: "/api/",
        dataType: "json",
        contentType: "application/json; charset=utf-8",
        success: function() {console.log("success")},
        data: JSON.stringify({"step": "step1", "data": "test"})
    });
}