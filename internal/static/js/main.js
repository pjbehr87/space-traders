// Global JS

function fetchUrl(url, { formData = new FormData(), successFn = ()=>{}, failFn = ()=>{}, errFn = ()=>{} }) {
	fetch(url, {
		method: 'POST',
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
