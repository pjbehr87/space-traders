'use strict';

const shipSymbol = document.querySelectorAll("#ship-symbol")[0].value;

document.querySelectorAll("#orbit-ship").forEach(obitShipBtn => {
	obitShipBtn.addEventListener('click', function () {
		let btnIcon = obitShipBtn.querySelectorAll('i')[0].outerHTML;
		obitShipBtn.innerHTML = btnIcon + ' Orbiting...';
		obitShipBtn.disabled = true;

		fetchUrl(
			`/my/ships/${shipSymbol}/orbit`,
			{
				successFn: () => {
					location.reload();
				},
				errFn: (err) => {
					alert(err);
					obitShipBtn.innerHTML = btnIcon + ' Orbit';
					obitShipBtn.disabled = false;
				}
			});
	});
});

document.querySelectorAll("#dock-ship").forEach(dockShipBtn => {
	dockShipBtn.addEventListener('click', function () {
		let btnIcon = dockShipBtn.querySelectorAll('i')[0].outerHTML;
		dockShipBtn.innerHTML = btnIcon + ' Docking...';
		dockShipBtn.disabled = true;

		fetchUrl(
			`/my/ships/${shipSymbol}/dock`,
			{
				successFn: () => {
					location.reload();
				},
				errFn: (err) => {
					alert(err);
					dockShipBtn.innerHTML = btnIcon + ' Dock';
					dockShipBtn.disabled = false;
				}
			});
	});
});

document.querySelectorAll(".navigate-ship").forEach(navigateShipBtn => {
	navigateShipBtn.addEventListener('click', function () {
		let btnIcon = navigateShipBtn.querySelectorAll('i')[0].outerHTML;
		navigateShipBtn.innerHTML = btnIcon + ' Navigating...';
		navigateShipBtn.disabled = true;

		let formData = new FormData();
		formData.append('waypointSymbol', navigateShipBtn.dataset.waypointSymbol);
		fetchUrl(
			`/my/ships/${shipSymbol}/navigate`,
			{
				formData: formData,
				successFn: () => {
					location.reload();
				},
				errFn: (err) => {
					alert(err);
					navigateShipBtn.innerHTML = btnIcon + ' Navagate';
					navigateShipBtn.disabled = false;
				}
			});
	});
});

document.querySelectorAll("#refuel-ship").forEach(fullRefuelShipBtn => {
	fullRefuelShipBtn.addEventListener('click', function () {
		let btnIcon = fullRefuelShipBtn.querySelectorAll('i')[0].outerHTML;
		fullRefuelShipBtn.innerHTML = btnIcon + ' Refueling...';
		fullRefuelShipBtn.disabled = true;

		let formData = new FormData();
		formData.append('units', fullRefuelShipBtn.dataset.units);
		fetchUrl(
			`/my/ships/${shipSymbol}/refuel`,
			{
				formData: formData,
				successFn: () => {
					location.reload();
				},
				errFn: (err) => {
					alert(err);
					fullRefuelShipBtn.innerHTML = btnIcon + ' Refuel';
					fullRefuelShipBtn.disabled = false;
				}
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
		let btnIcon = extractResourcesBtn.querySelectorAll('i')[0].outerHTML;
		extractResourcesBtn.innerHTML = btnIcon + ' Extracting...';
		extractResourcesBtn.disabled = true;

		fetchUrl(
			`/my/ships/${shipSymbol}/extract`,
			{
				successFn: (respJson) => {
					document.querySelectorAll('#ship-cargo-units')[0].textContent = respJson.cargo.units;
					miningCooldown(extractResourcesBtn, respJson.cooldown.remainingSeconds);
				},
				failFn: (respJson) => {
					miningCooldown(extractResourcesBtn, respJson.error.data.cooldown.remainingSeconds);
				},
				errFn: (err) => {
					alert(err);
					extractResourcesBtn.innerHTML = btnIcon + ' Extract';
					extractResourcesBtn.disabled = false;
				}
			});
	});
});

document.querySelectorAll(".sell-cargo").forEach(sellCargoBtn => {
	let
		goodsSymbol = sellCargoBtn.dataset.goodsSymbol,
		goodsUnits = sellCargoBtn.dataset.goodsUnits;

	sellCargoBtn.addEventListener('click', function () {
		let btnIcon = sellCargoBtn.querySelectorAll('i')[0].outerHTML;
		sellCargoBtn.innerHTML = btnIcon + ' Selling...';
		sellCargoBtn.disabled = true;

		let formData = new FormData();
		formData.append('symbol', goodsSymbol);
		formData.append('units', goodsUnits);
		fetchUrl(
			`/my/ships/${shipSymbol}/sell`,
			{
				formData: formData,
				successFn: (respJson) => {
					sellCargoBtn.closest('li').remove()
				},
				errFn: (err) => {
					alert(err);
					sellCargoBtn.innerHTML = btnIcon + ' Sell';
					sellCargoBtn.disabled = false;
				}
			});
	});
});
