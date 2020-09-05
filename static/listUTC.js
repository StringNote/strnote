function listUTCPromise(ids) {
	const apiURL = "https://strnote.com/api/v2/";
	return new Promise(function (callback, onerror) {
		const xhr = new XMLHttpRequest();
		xhr.open("POST", apiURL + "list", true);
		xhr.onreadystatechange = function () {
			if (xhr.readyState == 4) {
				const info = {
					result: xhr.response,
					request: xhr
				};
				switch (xhr.status) {
					case 200:	// OK
						try {
							callback(info);
						} catch (e) {
							info.error = e;
							onerror(info);
						}
						break;

					default:
						onerror(info);
						break;
				}
			}
		};
		const form = new FormData();
		form.append("ids", JSON.stringify(ids));
		xhr.send(form);
	});
}
