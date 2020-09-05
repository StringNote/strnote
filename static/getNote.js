/**
 * getNotePromise gets the note
 * @param uid is note-id
 * @license MIT
 */
function getNotePromise(uid) {
    return new Promise(function (callback, onerror) {
        const apiURL = "https://strnote.com/api/v2/";
        const xhr = new XMLHttpRequest();
        xhr.open("GET", apiURL + uid, true);
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
        xhr.send();
    });
}
