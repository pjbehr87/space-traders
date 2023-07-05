async function orbitShip(shipSymbol) {
	const resp = await fetch(`/my/ships/${shipSymbol}/orbit`, {
		method: 'POST'
	});

	respJson = await resp.json();

	return respJson;
}

document.querySelectorAll("#orbit-ship").forEach(obitShipBtn => {
	let shipSymbol = obitShipBtn.dataset.shipSymbol;

	obitShipBtn.addEventListener('click', function () {
		obitShipBtn.textContent = 'Orbiting...';
		obitShipBtn.disabled = true;
		orbitShip(shipSymbol)
			.then(resp => {
				if (resp["error"] !== undefined) {
					alert(resp.error.message);
				}
				location.reload();
			});
	});
});

async function dockShip(shipSymbol) {
	const resp = await fetch(`/my/ships/${shipSymbol}/dock`, {
		method: 'POST'
	});

	respJson = await resp.json();

	return respJson;
}

document.querySelectorAll("#dock-ship").forEach(dockShipBtn => {
	let shipSymbol = dockShipBtn.dataset.shipSymbol;

	dockShipBtn.addEventListener('click', function () {
		dockShipBtn.textContent = 'Docking...';
		dockShipBtn.disabled = true;
		dockShip(shipSymbol)
			.then(resp => {
				if (resp["error"] !== undefined) {
					alert(resp.error.message);
				}
				location.reload();
			});
	});
});

async function navigateShip(shipSymbol, waypointSymbol) {
	let formData = new FormData();
	formData.append('waypointSymbol', waypointSymbol);
	const resp = await fetch(`/my/ships/${shipSymbol}/navigate`, {
		method: 'POST',
		body: formData
	});

	respJson = await resp.json();

	return respJson;
}

document.querySelectorAll(".navigate-ship").forEach(navigateShipBtn => {
	let
		shipSymbol = navigateShipBtn.dataset.shipSymbol,
		waypointSymbol = navigateShipBtn.dataset.waypointSymbol;

	navigateShipBtn.addEventListener('click', function () {
		navigateShipBtn.textContent = 'Navigating...';
		navigateShipBtn.disabled = true;
		navigateShip(shipSymbol, waypointSymbol)
			.then(resp => {
				if (resp["error"] !== undefined) {
					alert(resp.error.message);
				}
				location.reload();
			});
	});
});
