

var kanrisha_method = {

	showTooltip: function (x, y, contents) {
		$('<div class="charts_tooltip">' + contents + '</div>').css( {
			position: 'absolute',
			display: 'none',
			top: y + 5,
			left: x + 5
		}).appendTo("body").fadeIn('fast');
	},

}

var km = kanrisha_method;

// Document Ready and the fun begin :) 
$(function () {


	/* Change Pattern ==================================== */

	$(".changePattern span").on("click", function(){
		var id = $(this).attr("id");
		$("body").css("background-image", "url('../Images/Textures/"+ id +".png')");
	});

	/* Opera Fix ========================================= */

	if ( $.browser['opera'] ) {
		$("aside").addClass('onlyOpera');
	}

	/* Charts ============================================ */

	if (!!$(".charts").offset() ) {
		var sin = [];
		var cos = [];

		for (var i = 0; i <= 20; i += 0.5){
			sin.push([i, Math.sin(i)]);
			cos.push([i, Math.cos(i)]);
		}

		// Display the Sin and Cos Functions
		$.plot($(".charts"), [ { label: "Cos", data: cos }, { label: "Sin", data: sin } ],
			{
				colors: ["#00AADD", "#FF6347"],

				series: {
					lines: {
						show: true,
						lineWidth: 2,
					},
					points: {show: true},
					shadowSize: 2,
				},

				grid: {
					hoverable: true,
					show: true,
					borderWidth: 0,
					tickColor: "#d2d2d2",
					labelMargin: 12,
				},

				legend: {
					show: true,
					margin: [0,-24],
					noColumns: 0,
					labelBoxBorderColor: null,
				},

				yaxis: { min: -1.2, max: 1.2},
				xaxis: {},
			});

		// Tooltip Show
		var previousPoint = null;
		$(".charts").bind("plothover", function (event, pos, item) {
			if (item) {
				if (previousPoint != item.dataIndex) {
					previousPoint = item.dataIndex;
					$(".charts_tooltip").fadeOut("fast").promise().done(function(){
						$(this).remove();
					});
					var x = item.datapoint[0].toFixed(2),
						y = item.datapoint[1].toFixed(2);
					km.showTooltip(item.pageX, item.pageY, item.series.label + " of " + x + " = " + y);
				}
			}
			else {
				$(".charts_tooltip").fadeOut("fast").promise().done(function(){
					$(this).remove();
				});
				previousPoint = null;
			}
		});

		if (!!$(".v_bars").offset() && !!$(".h_bars").offset() && !!$(".realtime_charts").offset()) {
			// Display Some Vertican bars
			$.plot($(".v_bars"), [ { data: [ [00,20], [20,50], [40,90], [60,30], [80,80], [100,60]] }, { data: [ [10,30], [30,80], [50,50], [70,10], [90,70] ] } ],
				{
					colors: ["#F7810C", "#E82E36"],

					series: {
						lines: {
							show: false,
							lineWidth: 2,
						},
						points: {show: false},
						shadowSize: 2,
						bars: {
							show: true,
							barWidth: 3,
							lineWidth: 1,
							fill: 0.8,
						}
					},

					grid: {
						hoverable: false,
						show: true,
						borderWidth: 0,
						tickColor: "#d2d2d2",
						labelMargin: 12,
					},

					legend: {
						show: false,
					},

					yaxis: { min: 0, max: 100},
					xaxis: { min: 0, max: 105},
				});

			// Display Some Vertical bars
			$.plot($(".h_bars"), [ { data: [ [20,20], [20,50], [40,00], [60,30], [80,80], [100,70]] }, { data: [ [10,10], [30,100], [50,40], [70,90], [90,60] ] } ],
				{
					colors: ["#F7810C", "#E82E36"],

					series: {
						lines: {
							show: false,
							lineWidth: 2,
						},
						points: {show: false},
						shadowSize: 2,
						bars: {
							show: true,
							barWidth: 3,
							lineWidth: 1,
							fill: 0.8,
							horizontal: true,
						}
					},

					grid: {
						hoverable: false,
						show: true,
						borderWidth: 0,
						tickColor: "#d2d2d2",
						labelMargin: 12,
					},

					legend: {
						show: false,
					},

					yaxis: { min: 0, max: 100},
					xaxis: { min: 0, max: 105},
				});

			// Display the realtime Charts
			// Generate a random data
			var data = [], totalPoints = 300;
			function  getRandomData () {
				if (data.length > 0)
					data = data.slice(1);
				// do a random walk
				while (data.length < totalPoints) {
					var prev = data.length > 0 ? data[data.length - 1] : 50;
					var y = prev + Math.random() * 10 - 5;
					if (y < 0)
						y = 0;
					if (y > 100)
						y = 100;
					data.push(y);
				}
				// zip the generated y values with the x values
				var res = [];
				for (var i = 0; i < data.length; ++i)
					res.push([i, data[i]])
				return res;
			}
			var realtime = $.plot($(".realtime_charts"), [ getRandomData() ],
				{
					colors: ["#00AADD"],

					series: {
						lines: {
							show: true,
							lineWidth: 2,
							fill: 0.65,
						},
						points: {show: false},
						shadowSize: 2,
					},

					grid: {
						show: true,
						borderWidth: 0,
						tickColor: "#d2d2d2",
						labelMargin: 12,
					},

					legend: {
						show: false,
					},

					yaxis: { min: 0, max: 105},
					xaxis: { min: 0, max: 250},
				}
			);
			function realtime_function() {
				realtime.setData([ getRandomData() ]);
				realtime.draw();
				setTimeout(realtime_function, 700);
			}
			realtime_function();
		}
	}

	// Pie Charts
	if(!!$(".pie_charts").offset()){
		$.plot($(".pie_charts"), [ { label: "iOS", data: 50 }, { label: "Android", data: 40 }, { label: "Windows", data: 30 }],
			{
				colors: ["#F7810C", "#00AADD", "#E82E36"],

				series: {
					pie: {
						show: true,
						tilt: 0.6,
						label: {
							show: true,
						}
					},
				},

				grid: {
					show: false,
				},

				legend: {
					show: true,
					margin: [0,-24],
					noColumns: 1,
					labelBoxBorderColor: null,
				},
			});

		// Donut Charts
		if(!!$(".donut_charts").offset()){
			$.plot($(".donut_charts"), [ { label: "iOS", data: 50 }, { label: "Android", data: 40 }, { label: "Windows", data: 30 }],
				{
					colors: ["#00AADD", "#F7810C", "#E82E36"],

					series: {
						pie: {
							show: true,
							innerRadius: 0.4,
						},
					},

					grid: {
						show: false,
					},

					legend: {
						show: true,
						margin: [0,-24],
						noColumns: 1,
						labelBoxBorderColor: null,
					},
				});
		}
	}

	/* Tables ============================================ */
	// Set the DataTables
	$(".datatable").dataTable({
		"sDom": "<'dtTop'<'dtShowPer'l><'dtFilter'f>><'dtTables't><'dtBottom'<'dtInfo'i><'dtPagination'p>>",
		"oLanguage": {
			"sLengthMenu": "Show entries _MENU_",
		},
		"sPaginationType": "full_numbers",
		"fnInitComplete": function(){
			$(".dtShowPer select").uniform();
			$(".dtFilter input").addClass("simple_field").css({
				"width": "auto",
				"margin-left": "15px",
			});
		}
	});

	// Table Resize-able
	$(".resizeable_tables").colResizable({
		liveDrag: true,
		minWidth: 40,
	});

	// Table with Tabs
	$( "#table_wTabs" ).tabs();

	// Check All Checkbox
	$(".tMainC").click(function(){
		var checked = $(this).prop("checked");
		var parent = $(this).closest(".twCheckbox");

		parent.find(".checker").each(function(){
			if (checked){
				$(this).find("span").addClass("checked");
				$(this).find("input").prop("checked", true);
			}else{
				$(this).find("span").removeClass("checked");
				$(this).find("input").prop("checked", false);
			}
		})
	});

	/* Forms ============================================= */
	$(".simple_form").uniform(); // Style The Checkbox and Radio



	/* Slider ============================================ */
	$(".sSimple").slider();

	$(".swMin").slider({
		range: "min",
		value: 80,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swmLabel" ).html( "$" + ui.value );
		}
	});

	$(".swMin-1").slider({
		range: "min",
		value: 120,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swmLabel" ).html( "$" + ui.value );
		}
	});

	$(".swMin-2").slider({
		range: "min",
		value: 220,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swmLabel" ).html( "$" + ui.value );
		}
	});

	$(".swMin-3").slider({
		range: "min",
		value: 350,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swmLabel" ).html( "$" + ui.value );
		}
	});

	$(".swMin-4").slider({
		range: "min",
		value: 450,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swmLabel" ).html( "$" + ui.value );
		}
	});

	$(".swMin-5").slider({
		range: "min",
		value: 600,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swmLabel" ).html( "$" + ui.value );
		}
	});

	$(".swMax").slider({
		range: "max",
		value: 600,
		min: 1,
		max: 700,
		slide: function( event, ui ) {
			$( ".swnLabel" ).html( "$" + ui.value );
		}
	});

	$( ".swRange" ).slider({
		range: true,
		min: 0,
		max: 500,
		values: [ 75, 300 ],
		slide: function( event, ui ) {
			$( ".swrLabel" ).html( "$" + ui.values[ 0 ] + " - $" + ui.values[ 1 ] );
		}
	});

	$( "#swVer-1" ).slider({
		orientation: "vertical",
		range: "min",
		min: 0,
		max: 100,
		value: 60,
	});

	$( "#swVer-2" ).slider({
		orientation: "vertical",
		range: "min",
		min: 0,
		max: 100,
		value: 40,
	});

	$( "#swVer-3" ).slider({
		orientation: "vertical",
		range: "min",
		min: 0,
		max: 100,
		value: 30,
	});

	$( "#swVer-4" ).slider({
		orientation: "vertical",
		range: "min",
		min: 0,
		max: 100,
		value: 15,
	});

	$( "#swVer-5" ).slider({
		orientation: "vertical",
		range: "min",
		min: 0,
		max: 100,
		value: 40,
	});

	$( "#swVer-6" ).slider({
		orientation: "vertical",
		range: "min",
		min: 0,
		max: 100,
		value: 80,
	});

	/* Progress ========================================== */

	$(".sProgress").progressbar({
		value: 40
	});

	$(".pwAnimate").progressbar({
		value: 1,
		create: function() {
			$(".pwAnimate .ui-progressbar-value").animate({"width":"100%"},{
				duration: 10000,
				step: function(now){
					$(".paValue").html(parseInt(now)+"%");
				},
				easing: "linear"
			})
		}
	});

	$(".pwuAnimate").progressbar({
		value: 1,
		create: function() {
			$(".pwuAnimate .ui-progressbar-value").animate({"width":"100%"},{
				duration: 30000,
				easing: 'linear',
				step: function(now){
					$(".pauValue").html(parseInt(now*10.24)+" Mb");
				},
				complete: function(){
					$(".pwuAnimate + .field_notice").html("<span class='must'>Upload Finished</span>");
				}
			})
		}
	});

	/* Tab Toggle ======================================== */

	$(".cwhToggle").click(function(){
		// Get Height
		var wC = $(this).parents().eq(0).find('.widget_contents');
		var wH = $(this).find('.widget_header_title');
		var h = wC.height();

		if (h == 0) {
			wH.addClass("i_16_downT").removeClass("i_16_cHorizontal");
			wC.css('height','auto').removeClass('noPadding');
		}else{
			wH.addClass("i_16_cHorizontal").removeClass("i_16_downT");
			wC.css('height','0').addClass('noPadding');
		}
	})

	/* Dialog ============================================ */

	$.fx.speeds._default = 400; // Adjust the dialog animation speed

	$(".bDialog").dialog({
		autoOpen: false,
		show: "fadeIn",
		modal: true,
	});

	$(".dConf").dialog({
		autoOpen: false,
		show: "fadeIn",
		modal: true,
		buttons: {
			"Yeah!": function() {
				$( this ).dialog( "close" );
			},
			"Never": function() {
				$( this ).dialog( "close" );
			}
		}
	});

	$(".bdC").live("click", function(){ /* change click to live */
		$(".bDialog").dialog( "open" );
		return false;
	});

	$(".bdcC").live("click", function(){ /* change click to live */
		$(".dConf").dialog( "open" );
		return false;
	});



	/* Drop Menu ========================================= */

	$(".drop_menu").parent().on("click", function(){
		var status = $(this).find(".drop_menu").css("display");
		if (status == "block"){
			$(this).find(".drop_menu").css("display", "none");
		}else{
			$(this).find(".drop_menu").css("display", "block");
		}
	});

	$(".top_tooltip").parent().on("hover", function(){
		var status = $(this).find(".top_tooltip").css("display");
		if (status == "block"){
			$(this).find(".top_tooltip").css("display", "none");
		}else{
			$(this).find(".top_tooltip").css("display", "block");
		}
	});

	/* Inline Dialog ===================================== */

	$(".iDialog").on("click", function(){
		$(this).fadeOut("slow").promise().done(function(){
			$(this).parent().remove();
		});
	});


	/* Internet 设置  selector===================================== */

		$("select#MTU_selector").change(function(){

			if ($(this).val() == "手动"){
				$("#MTU_size").attr("disabled",false);
				$("#MTU_size").removeClass("lightgray");
			}
			if ($(this).val() == "自动"){
				$("#MTU_size").attr("disabled",true);
				$("#MTU_size").addClass("lightgray");
			}
		});

	/* Internet 设置  mac克隆 checkbox===================================== */

	$("#mac_cb").change(function(){

		if ($(this).attr("checked")) {
			for (var i = 1; i <= 6; i++) {
				$("#mac" + i).attr("disabled", false);
				$("#mac" + i).removeClass("lightgray");
			}
			$("#mac_btn").removeClass("bg_lightgray");
			$("#mac_btn").attr("disabled", false);
		}
		if (!$(this).attr("checked")) {
			for (var i = 1; i <= 6; i++) {
				$("#mac" + i).attr("disabled", true);
				$("#mac" + i).addClass("lightgray");
			}
			$("#mac_btn").addClass("bg_lightgray");
			$("#mac_btn").attr("disabled", true);
		}
	});


	<!-- 基本页面 路由器设置 编辑 -->
	$("#basicEdit").click(function() {
		if ($(this).hasClass("blue")){
			$(this).removeClass("blue");
			$(this).addClass("edit_active");
			$("#basicInfo").css("display","none");
			$("#basicSetting").css("display","block");
		}
	});

	<!--  基本页面 路由器密码 编辑 -->
	$("#basicRouterEdit").click(function() {
		if ($(this).hasClass("blue")){
			$(this).removeClass("blue");
			$(this).addClass("edit_active");
			$("#basicRouterInfo").css("display","none");
			$("#basicRouterSetting").css("display","block");
		}
	});

	<!--  基本页面 路由器密码设置对话框 关闭按钮点击 -->
	$("#closeSetPassword").click(function() {
		$("#basicRouterEdit").removeClass("edit_active");
		$("#basicRouterEdit").addClass("blue");
		$("#basicRouterSetting").css("display","none")
		$("#basicRouterInfo").css("display","block");
	});


	<!-- 本地网络 路由器设置 编辑 -->
	$("#routerEdit").click(function() {
		if ($(this).hasClass("blue")){
			$(this).removeClass("blue");
			$(this).addClass("edit_active");
			$("#routerInfo").css("display","none");
			$("#routerSetting").css("display","block");
		}
	});

	<!-- 连接类型选择 -->
	$("select#linkType").change(function(){

		<!-- 连接类型选择为 静态IP时，显示静态IP配置项 -->
		if ($(this).val() == "0"){
			$("#staticIP").css("display","block");
			$("#PPPoE").css("display","none");
		}
		<!-- 连接类型选择为 PPPoE时，显示PPPoE配置项 -->
		if ($(this).val() == "1"){
			$("#staticIP").css("display","none");
			$("#PPPoE").css("display","block");
		}
	});


	<!-- 连接模式选择 -->
	$("input[name=PPPoE_link]").change(function(){
		var PPPoE_link = $(this).val();
		if(PPPoE_link=="0"){
			$("#maxTime").attr("disabled",false);
			$("#maxTime").removeClass("lightgray");
			$("#repeatDuration").addClass("lightgray");
			$("#repeatDuration").attr("disabled",true);
		}
		if(PPPoE_link=="1"){
			$("#repeatDuration").attr("disabled",false);
			$("#repeatDuration").removeClass("lightgray");
			$("#maxTime").addClass("lightgray");
			$("#maxTime").attr("disabled",true);
		}
	});


	<!-- 计算IP访问地址 Begin -->

	//初始化
	$("#beginIP").text("192.168." + $("#beginIPThird").val() + "." + $("#beginIPFourth").val());
	var lastIP = $("#beginIPFourth").val() * 1 + $("#maxUser").val() * 1 - 1 * 1;
	$("#endIP").text("192.168." + $("#beginIPThird").val() + "." + lastIP);

	$("#beginIPThird").change(function() {
		$("#beginIP").text("192.168." + $("#beginIPThird").val() + "." + $("#beginIPFourth").val());
		var lastIP = $("#beginIPFourth").val() * 1 + $("#maxUser").val() * 1 - 1 * 1;
		$("#endIP").text("192.168." + $("#beginIPThird").val() + "." + lastIP);
	});

	$("#beginIPFourth").change(function() {
		$("#beginIP").text("192.168." + $("#beginIPThird").val() + "." + $("#beginIPFourth").val());
		var lastIP = $("#beginIPFourth").val() * 1 + $("#maxUser").val() * 1 - 1 * 1;
		$("#endIP").text("192.168." + $("#beginIPThird").val() + "." + lastIP);
	});

	$("#maxUser").change(function() {
		$("#beginIP").text("192.168." + $("#beginIPThird").val() + "." + $("#beginIPFourth").val());
		var lastIP = $("#beginIPFourth").val() * 1 + $("#maxUser").val() * 1 - 1 * 1;
		$("#endIP").text("192.168." + $("#beginIPThird").val() + "." + lastIP);
	});
	<!-- 计算IP访问地址 End -->


	/* 管理页面  UPnP 启用 checkbox===================================== */

	$("#UPnP_cb").change(function(){

		if ($(this).attr("checked")) {

			$("#magAllowUser").attr("disabled", false);
			$("#magNoInter").attr("disabled", false);
			$("#magAllowUser").parent(".checked").css("opacity","1");
			$("#magNoInter").parent(".checked").css("opacity","1");

		}
		if (!$(this).attr("checked")) {
			$("#magAllowUser").attr("disabled", true);
			$("#magNoInter").attr("disabled", true);
			$("#magAllowUser").parent(".checked").css("opacity","0.2");
			$("#magNoInter").parent(".checked").css("opacity","0.2");
		}
	});


	/* 本地网络 DHCP服务器 启用 checkbox===================================== */

	$("#DCHP_cb").change(function(){

		if ($(this).attr("checked")) {
			$("#DHCPActive").find("input").attr('disabled', false);
			$("#DHCPActive").find("input").css('color', '#000');
		}
		if (!$(this).attr("checked")) {
			$("#DHCPActive").find("input").attr('disabled', true);
			$("#DHCPActive").find("input").css('color', '#aaa');
		}
	});



	/* 本地网络 DHCP服务器 启用 checkbox===================================== */

	$("#DCHP_cb").change(function(){

		if ($(this).attr("checked")) {
			$("#DHCPActive").find("input").attr('disabled', false);
			$("#DHCPActive").find("input").css('color', '#000');
		}
		if (!$(this).attr("checked")) {
			$("#DHCPActive").find("input").attr('disabled', true);
			$("#DHCPActive").find("input").css('color', '#aaa');
		}
	});


	$('input[type=radio][name=NAT]').change(function() {

		if (this.value == '0') {
			$("#OSPFarea").attr('disabled', true);
			$("#OSPFlabel").css('color', 'lightgray');
		}
		else if (this.value == '1') {
			$("#OSPFarea").attr('disabled', true);
			$("#OSPFlabel").css('color', 'lightgray');
		}
		else if (this.value == '2') {
			$("#OSPFarea").attr('disabled', false);
			$("#OSPFlabel").css('color', '#000');
		}
	});




});


