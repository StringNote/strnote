function sendTokenPromise(token, name, mes) {
    return new Promise(function (callback, onerror) {
        const xhr = new XMLHttpRequest();
        xhr.open("POST", "/api/v1/user", true);
        xhr.setRequestHeader('X-Requested-With', token);
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
        form.append("name", name);
        form.append("mes", mes);
        xhr.send(form);
    });
}
