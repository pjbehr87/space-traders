async function purchaseShip(shipType, waypointSymbol) {
	let formData = new FormData();
	formData.append('shipType', shipType);
	formData.append('waypointSymbol', waypointSymbol);
	const resp = await fetch('/my/ships', { 
		method: 'POST',
		body: formData
	});

	respJson = await resp.json();

	return respJson;
}

document.querySelectorAll(".purchase-ship").forEach(purchaseBtn => {
	console.log(purchaseBtn)
	let
		shipType = purchaseBtn.dataset.shipType,
		waypointSymbol = purchaseBtn.dataset.waypointSymbol;
	
	purchaseBtn.addEventListener('click', function () {
		purchaseShip(shipType, waypointSymbol)
			.catch(resp => console.log(resp))
			.then(resp => {
				if(resp["error"] !== undefined) {
					alert(resp.error.message);
				}
				location.href('/my/ships');
			});
	});
});
