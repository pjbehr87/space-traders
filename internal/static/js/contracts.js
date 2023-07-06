document.querySelectorAll(".accept-contract").forEach($thisBtn => {
	acceptBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		let contractId = $thisBtn.dataset.contractId
		fetchUrl(
			`/my/contracts/${contractId}/accept`,
			{
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