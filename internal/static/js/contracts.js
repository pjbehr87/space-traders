document.querySelectorAll(".accept-contract").forEach($acceptBtn => {
	$acceptBtn.addEventListener('click', function () {
		btnAction($acceptBtn);

		let contractId = $acceptBtn.dataset.contractId
		fetchUrl(
			`/my/contracts/${contractId}/accept`,
			{
				successFn: () => {
					location.reload();
				},
				errFn: (err) => {
					alert(err);
					btnDefault($acceptBtn);
				}
			});
	});
});
