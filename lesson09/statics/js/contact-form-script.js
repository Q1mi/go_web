$("#contactForm").validator().on("submit", function (event) {
    if (event.isDefaultPrevented()) {
		formError();
    } else {
        event.preventDefault();
        submitForm();
    }
});


function submitForm(){
    // Initiate Variables With Form Content
    var fname = $("#fname").val();
	var lname = $("#lname").val();
    var email = $("#email").val();
    var message = $("#message").val();


    $.ajax({
        type: "POST",
        url: "mail/php/form-process.php",
        data: "fname=" + fname + "&lname=" + lname + "&email=" + email + "&message=" + message,
        success : function(text){
            if (text == "success"){
                formSuccess();
            } else {
                formError();
                submitMSG(false,text);
            }
        }
    });
}

function formSuccess(){
    $("#contactForm")[0].reset();
    submitMSG(true, "You mail has been sent successfully!")
}
function formError(){
    $("#contactForm").removeClass().addClass('animated').one('webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend', function(){
        $(this).removeClass();
    });
}
function submitMSG(valid, msg){
    if(valid){
        var msgClasses = "text-success";
    } else {
        var msgClasses = "text-danger";
    }
    $("#msgSubmit").removeClass().addClass(msgClasses).text(msg);
}