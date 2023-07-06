'use strict';

const shipSymbol = document.querySelectorAll("#ship-symbol")[0].value;

async function orbitShip() {
	const resp = await fetch(`/my/ships/${shipSymbol}/orbit`, {
		method: 'POST'
	});

	if (resp.status != 200) {
		let respJson = await resp.json()
		throw respJson.error.message
	}
	return await resp.text();
}

document.querySelectorAll("#orbit-ship").forEach(obitShipBtn => {
	obitShipBtn.addEventListener('click', function () {
		let btnIcon = obitShipBtn.querySelectorAll('i')[0].outerHTML;
		obitShipBtn.innerHTML = btnIcon + ' Orbiting...';
		obitShipBtn.disabled = true;
		orbitShip()
			.then(() => location.reload())
			.catch(err => {
				alert(err);
				obitShipBtn.innerHTML = btnIcon + ' Orbit';
				obitShipBtn.disabled = false;
			});
	});
});

async function dockShip() {
	const resp = await fetch(`/my/ships/${shipSymbol}/dock`, {
		method: 'POST'
	});

	if (resp.status != 200) {
		let respJson = await resp.json()
		throw respJson.error.message
	}
	return await resp.text();
}

document.querySelectorAll("#dock-ship").forEach(dockShipBtn => {
	dockShipBtn.addEventListener('click', function () {
		let btnIcon = dockShipBtn.querySelectorAll('i')[0].outerHTML;
		dockShipBtn.innerHTML = btnIcon + ' Docking...';
		dockShipBtn.disabled = true;
		dockShip()
			.then(() => location.reload())
			.catch(err => {
				alert(err);
				dockShipBtn.innerHTML = btnIcon + ' Dock';
				dockShipBtn.disabled = false;
			});
	});
});

async function navigateShip(waypointSymbol) {
	let formData = new FormData();
	formData.append('waypointSymbol', waypointSymbol);
	const resp = await fetch(`/my/ships/${shipSymbol}/navigate`, {
		method: 'POST',
		body: formData
	});

	if (resp.status != 200) {
		let respJson = await resp.json()
		throw respJson.error.message
	}
	return await resp.text();
}

document.querySelectorAll(".navigate-ship").forEach(navigateShipBtn => {
	let waypointSymbol = navigateShipBtn.dataset.waypointSymbol;

	navigateShipBtn.addEventListener('click', function () {
		let btnIcon = navigateShipBtn.querySelectorAll('i')[0].outerHTML;
		navigateShipBtn.innerHTML = btnIcon + ' Navigating...';
		navigateShipBtn.disabled = true;
		navigateShip(waypointSymbol)
			.then(() => location.reload())
			.catch(err => {
				alert(err);
				navigateShipBtn.innerHTML = btnIcon + ' Navagate';
				navigateShipBtn.disabled = false;
			});
	});
});

async function fullRefuelShip(units) {
	let formData = new FormData();
	formData.append('units', units);
	const resp = await fetch(`/my/ships/${shipSymbol}/refuel`, {
		method: 'POST',
		body: formData
	});

	if (resp.status != 200) {
		let respJson = await resp.json()
		throw respJson.error.message
	}
	return await resp.text();
}

document.querySelectorAll("#refuel-ship").forEach(fullRefuelShipBtn => {
	let fuelUnits = fullRefuelShipBtn.dataset.fuelUnits;

	fullRefuelShipBtn.addEventListener('click', function () {
		let btnIcon = fullRefuelShipBtn.querySelectorAll('i')[0].outerHTML;
		fullRefuelShipBtn.innerHTML = btnIcon + ' Refueling...';
		fullRefuelShipBtn.disabled = true;
		fullRefuelShip(fuelUnits)
			.then(() => location.reload())
			.catch(err => {
				alert(err);
				fullRefuelShipBtn.innerHTML = btnIcon + ' Refuel';
				fullRefuelShipBtn.disabled = false;
			});
	});
});

function miningCooldown(btn, remainingSeconds) {
	let
		btnIconMineCart = btn.querySelectorAll('i')[0],
		btnIconCooldown = btnIconMineCart.cloneNode(true);
	btnIconCooldown.classList.remove("bi-minecart-loaded");
	btn.classList.remove("btn-primary");
	btnIconCooldown.classList.add("bi-clock-history");
	btn.classList.add("btn-warning");
	btn.innerHTML = btnIconCooldown.outerHTML + ' Extracting (' + remainingSeconds + 's)';

	let countdown = setInterval(() => {
		btn.innerHTML = btnIconCooldown.outerHTML + ' Extracting (' + --remainingSeconds + 's)';
		if (remainingSeconds === 0) {
			btn.disabled = false;
			btn.innerHTML = btnIconCooldown.outerHTML + ' Extract';
			clearInterval(countdown);
		}
	}, 1000);

	setTimeout(() => {
		btn.classList.remove("btn-warning");
		btn.classList.add("btn-primary");
		btn.disabled = false;
		btn.innerHTML = btnIconMineCart.outerHTML + ' Extract';
	}, (remainingSeconds * 1000));
}

document.querySelectorAll("#extract-minerals").forEach(extractResourcesBtn => {
	extractResourcesBtn.addEventListener('click', function () {
		let
			btnIcon = extractResourcesBtn.querySelectorAll('i')[0],
			btnIconHtml = btnIcon.outerHTML;
		extractResourcesBtn.innerHTML = btnIconHtml + ' Extracting...';
		extractResourcesBtn.disabled = true;
		fetch(`/my/ships/${shipSymbol}/extract`, {
			method: 'POST'
		}).then(resp => {
			console.log(resp);
			if (resp.status !== 200) {
				resp.json().then(respJson => {
					// Extract on cooldown
					if (respJson.error.code === 4000) {
						miningCooldown(extractResourcesBtn, respJson.error.data.cooldown.remainingSeconds);
					}
					else {
						console.log(respJson)
					}
				});
			}
			else {
				resp.json().then(respJson => {
					console.log(respJson);
					document.querySelectorAll('#ship-cargo-units')[0].textContent = respJson.cargo.units;
					miningCooldown(extractResourcesBtn, respJson.cooldown.remainingSeconds);
				});
			}
		}).catch(err => {
			alert(err);
			extractResourcesBtn.innerHTML = btnIconHtml + ' Extract';
			extractResourcesBtn.disabled = false;
		});
	});
});

async function sellCargo(goodsSymbol, goodsUnits) {
	let formData = new FormData();
	formData.append('symbol', goodsSymbol);
	formData.append('units', goodsUnits);
	const resp = await fetch(`/my/ships/${shipSymbol}/sell`, {
		method: 'POST',
		body: formData
	});

	if (resp.status != 200) {
		let respJson = await resp.json();
		console.log(respJson);
		throw respJson.error.message;
	}
	return await resp.text();
}

document.querySelectorAll(".sell-cargo").forEach(sellCargoBtn => {
	let
		goodsSymbol = sellCargoBtn.dataset.goodsSymbol,
		goodsUnits = sellCargoBtn.dataset.goodsUnits;

	sellCargoBtn.addEventListener('click', function () {
		let btnIcon = sellCargoBtn.querySelectorAll('i')[0].outerHTML;
		sellCargoBtn.innerHTML = btnIcon + ' Selling...';
		sellCargoBtn.disabled = true;
		sellCargo(goodsSymbol, goodsUnits)
			.then(resp => {
				console.log(resp);
				sellCargoBtn.closest('li').remove()
			})
			.catch(err => {
				alert(err);
				sellCargoBtn.innerHTML = btnIcon + ' Sell';
				sellCargoBtn.disabled = false;
			});
	});
});
