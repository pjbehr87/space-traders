async function acceptContract(contractId) {
	const resp = await fetch(`/my/contracts/${contractId}/accept`, { method: 'POST' });
	console.log(resp);
}

document.querySelectorAll(".accept-contract").forEach(acceptBtn => {
	let contractId = this.dataset.contractId
	acceptBtn.addEventListener('click', function () {
		acceptContract(contractId);
		location.href(`/my/contracts/${contractId}`);
	});
})