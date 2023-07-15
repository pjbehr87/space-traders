'use strict';

const 
	SHIP_SYMBOL = document.querySelectorAll("#ship-symbol")[0].value,
	MYSHIP_URL = `/my/ships/${SHIP_SYMBOL}`;
// ////////////////////
// Ship controls
// ////////////////////
document.querySelectorAll("#orbit-ship").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		fetchUrl(
			`${MYSHIP_URL}/orbit`,
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

document.querySelectorAll("#dock-ship").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		fetchUrl(
			`${MYSHIP_URL}/dock`,
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

document.querySelectorAll(".navigate-ship").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		let formData = new FormData();
		formData.append('waypointSymbol', $thisBtn.dataset.waypointSymbol);
		fetchUrl(
			`${MYSHIP_URL}/navigate`,
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

function miningCooldown($btn, remainingSeconds) {
	let
		$btnIconMineCart = $btn.querySelectorAll('i')[0],
		$btnIconCooldown = $btnIconMineCart.cloneNode(true);
	$btnIconCooldown.classList.remove("bi-minecart-loaded");
	$btn.classList.remove("btn-primary");
	$btnIconCooldown.classList.add("bi-clock-history");
	$btn.classList.add("btn-warning");
	$btn.innerHTML = $btnIconCooldown.outerHTML + ' Extracting (' + remainingSeconds + 's)';

	let countdown = setInterval(() => {
		$btn.innerHTML = $btnIconCooldown.outerHTML + ' Extracting (' + --remainingSeconds + 's)';
		if (remainingSeconds === 0) {
			$btn.disabled = false;
			$btn.innerHTML = $btnIconCooldown.outerHTML + ' Extract';
			clearInterval(countdown);
		}
	}, 1000);

	setTimeout(() => {
		$btn.classList.remove("btn-warning");
		$btn.classList.add("btn-primary");
		$btn.disabled = false;
		$btn.innerHTML = $btnIconMineCart.outerHTML + ' Extract';
	}, (remainingSeconds * 1000));
}

document.querySelectorAll("#extract-minerals").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		fetchUrl(
			`${MYSHIP_URL}/extract`,
			{
				successFn: (respJson) => {
					document.querySelectorAll('#ship-cargo-units')[0].textContent = respJson.cargo.units;
					miningCooldown($thisBtn, respJson.cooldown.remainingSeconds);
				},
				failFn: (respJson) => {
					miningCooldown($thisBtn, respJson.error.data.cooldown.remainingSeconds);
				},
				errFn: (err) => {
					alert(err);
					btnDefault($thisBtn);
				}
			});
	});
});

// ////////////////////
// Marketplace
// ////////////////////
document.querySelectorAll("#refuel-ship").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		let formData = new FormData();
		formData.append('units', $thisBtn.dataset.fuelUnits);
		fetchUrl(
			`${MYSHIP_URL}/refuel`,
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

document.querySelectorAll(".sell-cargo").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		let formData = new FormData();
		formData.append('symbol', $thisBtn.dataset.goodsSymbol);
		formData.append('units', $thisBtn.dataset.goodsUnits);
		fetchUrl(
			`${MYSHIP_URL}/sell`,
			{
				formData: formData,
				successFn: () => {
					$thisBtn.closest('li').remove()
				},
				errFn: (err) => {
					alert(err);
					btnDefault($thisBtn);
				}
			});
	});
});

// ////////////////////
// Shipyard
// ////////////////////
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



document.querySelectorAll(".deliver-cargo").forEach($thisBtn => {
	$thisBtn.addEventListener('click', function () {
		btnAction($thisBtn);

		let
			btnData = $thisBtn.dataset,
			contractId = $thisBtn,
			formData = new FormData();
		formData.append('shipSymbol', btnData.shipSymbol);
		formData.append('tradeSymbol', btnData.tradeSymbol);
		formData.append('units', btnData.units);
		fetchUrl(
			`/my/contracts/${contractId}/deliver`,
			{
				formData: formData,
				successFn: () => {
					$thisBtn.closest('li').remove()
				},
				errFn: (err) => {
					alert(err);
					btnDefault($thisBtn);
				}
			});
	});
});
