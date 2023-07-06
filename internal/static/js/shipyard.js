document.querySelectorAll(".purchase-ship").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn)

		let formData = new FormData();
		formData.append('shipType', $thisBtn.dataset.shipType);
		formData.append('waypointSymbol', $thisBtn.dataset.waypointSymbol);
		fetchUrl(
			'/my/ships',
			{
				formData: formData,
				successFn: () => {
					location.reload();
				},
				errFn: (err) => {
					alert(err);
					btnDefault($thisBtn);
				}
			});
	});
});
