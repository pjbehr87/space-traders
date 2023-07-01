async function purchaseShip(shipType, waypointSymbol) {
	let formData = new FormData();
	formData.append('shipType', shipType);
	formData.append('waypointSymbol', waypointSymbol);
	const resp = await fetch('/my/ships', { 
		method: 'POST',
		body: formData
	});

	if(resp.status != 200) {
		let respJson = await resp.json()
		throw respJson.error.message
	}
}

document.querySelectorAll(".purchase-ship").forEach(purchaseBtn => {
	let
		shipType = purchaseBtn.dataset.shipType,
		waypointSymbol = purchaseBtn.dataset.waypointSymbol;
	
	purchaseBtn.addEventListener('click', function () {
		purchaseShip(shipType, waypointSymbol)
			.then(() => {
				location.href = '/my/ships';
			})
			.catch(err => alert(err));
	});
});
