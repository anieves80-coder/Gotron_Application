$(document).ready(function () {
	let ws = new WebSocket('ws://localhost:' + global.backendPort + '/web/app/events');
	const dt = new Date();
	const date = dt.getMonth() + 1 + '/' + dt.getDate() + '/' + dt.getFullYear();
	let prev;
	$('#dateInput').val(date);	

	$('#frm').on('submit', function (e) {
		e.preventDefault();
		const data = {
			rma: $('#rmaInput').val(),
			sn1: $('#sn1Input').val(),
			sn2: $('#sn2Input').val(),
			frmDate: $('#dateInput').val(),
			comment: $('#msgTextarea').val()
		};
		ws.send(
			JSON.stringify({
				event: 'add-one',
				data
			})
		);
	});

	$('#searchBtn').on('click', function (e) {
		e.preventDefault();
		const data = {
			rma: $('#rmaInput').val().trim(),
			sn1: $('#sn1Input').val().trim(),
			frmDate: $('#dateInput').val().trim(),
			// update: false
		};
		
		ws.send(
			JSON.stringify({
				event: 'get-searchBy'
			})
		);
		 
	});

	$('#tableResults').on('click','.resultRows', function (e) {	
		prev = $(this).attr("id");				
		ws.send(
			JSON.stringify({
				event: 'get-searchBy',
				data: { rma: $(this).attr("id"), update: "true" }
			})
		);
	});

	$('#modifyBtn').on('click', function (e) {
		e.preventDefault();
		const data = {
			rma: $('#rmaInput').val().trim(),
			sn1: $('#sn1Input').val().trim(),
			sn2: $('#sn2Input').val().trim(),
			frmDate: $('#dateInput').val().trim(),
			comment: $('#msgTextarea').val().trim(),
			prev
		};
		console.log(data);
		ws.send(
			JSON.stringify({
				event: 'update-one',
				data
			})
		);
	});

	$('#frm').on('reset', function (e) {
		setTimeout(function () {
			$('#dateInput').val(date);
		});
	});

	$('input[type=radio]').change(function () {
		const opt = this.value;

		switch (opt) {
			case 'modify':
				setModify();
				break;
			case 'search':
				setSearch();
				break;
			default:
				setAdd();
		}
	});

	function setSearch() {
		$('#addBtn, #modifyBtn').addClass('btnHide');
		$('#searchBtn').removeClass('btnHide');
		$('#dateInput, #sn2Input').val('');
		$('#sn2Input').attr('disabled', 'disabled');
	}
	function setAdd() {
		$('#searchBtn, #modifyBtn').addClass('btnHide');
		$('#addBtn').removeClass('btnHide');
		$('#dateInput').val(date);
		$('#sn2Input').removeAttr('disabled');
	}
	function setModify() {
		$('#addBtn, #searchBtn').addClass('btnHide');
		$('#modifyBtn').removeClass('btnHide');
		$('#sn2Input').removeAttr('disabled');
		$('#dateInput').val('');
	}
	function showResult(obj) {
		let cnt = 0;
		console.log(obj.eventData)
		obj.eventData.forEach((element) => {
			const e = JSON.parse(element);
			cnt++;
			$('#tableResults').append(`
				<tr class="resultRows" id="${e.rma}">
					<th scope="row">${cnt}</th>
					<td>${e.rma}</td>
					<td>${e.sn1}</td>
					<td>${e.sn2}</td>
					<td>${e.date}</td>
					<td>${e.comment}</td>
				</tr>                
			`);
		});
	}

	ws.onmessage = (message) => {
		const obj = JSON.parse(message.data);		
		if (obj.event === 'show-results') {
			$('#tableResults').empty();
			showResult(obj);
		} 
		if (obj.event === 'show-inForm') {
			const e = JSON.parse(obj.eventData);			
			$("#option3").prop("checked", true);
			setModify();
			$('#rmaInput').val(e.rma),
			$('#sn1Input').val(e.sn1),
			$('#sn2Input').val(e.sn2),
			$('#dateInput').val(e.date),
			$('#msgTextarea').val(e.comment)
		} 
	};
	
});
