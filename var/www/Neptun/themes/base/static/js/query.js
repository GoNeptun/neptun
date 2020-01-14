(function($) {
	$(document).ready(function() {


    $(function(){
       $('#reqnewemail').click(function(){
            reqnewemail ();
       });
    });

		$(function(){
			 $('#sysmsg').click(function(){
						$('#sysmsg').hide();
			 });
		});

		function showsysmsg() {
			$('#sysmsg').show();
		};

		$(function(){
			 $('#chpassbtn').click(function(){
						chpass ();
			 });
		});

		$(function(){
			 $('#chemailbtn').click(function(){
						chemail ();
			 });
		});




		function chpass () {
			var func = "changepass";
			$.getJSON("/user/",
		{
		param: func,
		oldpass: $("#oldpass").val(),
		newpass: $("#newpass").val(),
		 },
		function (data) {
			if (data.success != 0) {
				$("#servicemsg").html(data.answer);
			} else {
				$("#sysmsg").html(data.error);
				showsysmsg();
			}
		 });
		}

		function chemail () {
			var func = "changeemail";
			$.getJSON("/user/",
		{
		param: func,
		email: $("#email").val(),
		pass: $("#pass").val(),

		 },
		function (data) {
			if (data.success != 0) {
				$("#servicemsg").html(data.answer);
			} else {
				$("#sysmsg").html(data.error);
				showsysmsg();
			}
		 });
		}

		function reqnewemail () {
			var func = "reqnewemail";
			$.getJSON("/user/",
		{
		param: func
		 },
		function (data) {
			if (data.success != 0) {
				$("#servicemsg").html(data.answer);
			} else {
				$("#sysmsg").html(data.error);
				showsysmsg();
			}
		 });
		}




  });
  }) (jQuery);
