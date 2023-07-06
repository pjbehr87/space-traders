// Global JS

function fetchUrl(url, { method = 'POST', formData = new FormData(), successFn = ()=>{}, failFn = ()=>{}, errFn = ()=>{} }) {
	fetch(url, {
		method: method,
		body: formData
	}).then(resp => {
		resp.json().then(respJson => {
			console.log(respJson)
			if (resp.status !== 200) {
				failFn(respJson);
			}
			else {
				successFn(respJson);
			}
		});
	}).catch(err => {
		errFn(err);
	});
}

function btnAction(btn) {
	let btnIcon = btn.querySelectorAll('i')[0].outerHTML;
	btn.dataset.default = btn.textContent.trim();
	btn.innerHTML = btnIcon + ' ' + btn.dataset.action;
	btn.disabled = true;
}
function btnDefault(btn) {
	let btnIcon = btn.querySelectorAll('i')[0].outerHTML;
	btn.innerHTML = btnIcon + ' ' + btn.dataset.default;
	btn.disabled = false;
}
