$(function() {
	$("#signin").click(function (e) {
		e.preventDefault();

		var data = {username: $("#username").val(), password: $("#password").val()};
		$.post("/signin", data, function (response) {
			if (response.success) {
				window.location = "/"
			}
		})
	});
})