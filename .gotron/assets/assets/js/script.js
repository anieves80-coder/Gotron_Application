$(document).ready(function() {
	let ws = new WebSocket('ws://localhost:' + global.backendPort + '/web/app/events');
	const dt = new Date();
	const date = dt.getMonth() + 1 + '/' + dt.getDate() + '/' + dt.getFullYear();
	$('#dateInput').val(date);

	$('#frm').on('submit', function(e) {
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

	$('#searchBtn').on('click', function(e) {
		e.preventDefault();
		const data = {
			rma: $('#rmaInput').val(),
			sn1: $('#sn1Input').val(),
			frmDate: $('#dateInput').val()
		};
		if (!data.rma && !data.sn1 && !data.Date) {
			ws.send(
				JSON.stringify({
					event: 'get-all'
				})
			);
		} else {
			ws.send(
				JSON.stringify({
					event: 'get-search',
					data
				})
			);
		}
	});

	$('#modifyBtn').on('click', function(e) {
		e.preventDefault();
		alert('ok there');
	});

	$('#frm').on('reset', function(e) {
		setTimeout(function() {
			$('#dateInput').val(date);
		});
	});

	$('input[type=radio]').change(function() {
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

	ws.onmessage = (message) => {
		const obj = JSON.parse(message.data);
		let cnt = 0;
		$('#tableResults').empty();
		if (obj.event === 'get-all') {
			obj.eventData.forEach((element) => {
				const e = JSON.parse(element);
				cnt++;
				$('#tableResults').append(`
                    <tr>
                        <th scope="row">${cnt}</th>
                        <td>${e.rma}</td>
                        <td>${e.sn1}</td>
                        <td>${e.sn2}</td>
                        <td>${e.comment}</td>
                    </tr>                
                `);
			});
		}
	};
});
