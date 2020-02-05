(function($) {
	$(document).ready(function() {


    $(function(){
       $('#maint').click(function(){
            setmaint ();
       });
    });

		 $(function(){
       $('#ureg').click(function(){
            setreg ();
       });
    });

		$(function(){
			 $('#emc').click(function(){
						setmail ();
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

		function setmaint () {
			var func = "maintenance";
			$.getJSON("/adminfuncs/",
		{
		param: func
		 },
		function (data) {
			if (data.success != 0) {
				$("#mm").html(data.answer.status);
				$("#mm").css("color", data.answer.style);
				$("#maintbtn").html(data.answer.button);
			} else {
				$("#sysmsg").html(data.error);
				showsysmsg();
			}
		 });
		}

		function setreg () {
			var func = "regstatus";
			$.getJSON("/adminfuncs/",
		{
		param: func
		 },
		function (data) {
			if (data.success != 0) {
				$("#ur").html(data.answer.status);
				$("#ur").css("color", data.answer.style);
				$("#uregbtn").html(data.answer.button);
			} else {
				$("#sysmsg").html(data.error);
				showsysmsg();
			}
		 });
		}

		function setmail () {
			var func = "emailcheck";
			$.getJSON("/adminfuncs/",
		{
		param: func
		 },
		function (data) {
			if (data.success != 0) {
				$("#ec").html(data.answer.status);
				$("#ec").css("color", data.answer.style);
				$("#emcbtn").html(data.answer.button);
			} else {
				$("#sysmsg").html(data.error);
				showsysmsg();
			}
		 });
		}



  });
  }) (jQuery);
