/**
 * @license MIT
 */

class Note {
	parent;
	uid;
	name;
	mes;
	utc;
	date;
	handle;
	constructor(parent, uid, handle) {
		this.parent = parent;
		this.uid = uid;
		this.handle = handle;
		this.Refresh();
		return this;
	}
	Refresh() {
		const note = this;
		note.getNotePromise(this.uid)
			.then(function (info) {
				const obj = JSON.parse(info.result);
				note.name = obj.name;
				note.mes = obj.mes;
				note.utc = obj.utc;
				note.date = note.local2str(note.utc2local(obj.utc));
				note.handle(note);
			}).catch(function (error) {
				note.parent.onError(error);
			});
	}
	getNotePromise(uid) {
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
	utc2local(ymdhms) {
		const ye = ymdhms.substring(0, 4);
		const mo = ymdhms.substring(4, 6) - 1;
		const da = ymdhms.substring(6, 8);
		const ho = ymdhms.substring(8, 10);
		const mi = ymdhms.substring(10, 12);
		const se = ymdhms.substring(12, 14);
		return new Date(Date.UTC(ye, mo, da, ho, mi, se));
	}
	local2str(dt) {
		const ye = dt.getFullYear();
		const mo = "" + (dt.getMonth() + 1);
		const da = "" + dt.getDate();
		const ho = "" + dt.getHours();
		const mi = "" + dt.getMinutes();
		const se = "" + dt.getSeconds();
		return ye + "/" + mo.padStart(2, "0") + "/" + da.padStart(2, "0")
			+ " " + ho.padStart(2, "0") + ":" + mi.padStart(2, "0") + ":" + se.padStart(2, "0");
	}
}

class Notes {
	notes = [];
	notedic = {};
	onError = error => { console.log(error); };
	timer;
	setInterval(interval) {
		if (this.timer) {
			clearInterval(this.timer);
			this.timer = "";
		}
		if (interval && interval>0) {
			this.timer = setInterval(this.Refresh.bind(this), interval);
		}
	}
	Add(uid, handle) {
		const note = new Note(this, uid, handle);
		this.notedic[uid] = note;
		this.notes.push(note);
		return note;
	}
	Get(uid) {
		return this.notedic[uid];
	}
	Refresh() {
		const onError = this.onError;
		const notedic = this.notedic;
		const uids = this.notes.map(note => note.uid);
		this.listUTCPromise(uids)
			.then(function (info, result) {
				const cutcs = JSON.parse(info.result)
				Object.keys(cutcs).forEach(uid => {
					const utc = cutcs[uid];
					const note = notedic[uid];
					if (utc != note.utc) {
						note.Refresh();
					}
				});
			}).catch(function (error) {
				onError(error);
			});
	}
	listUTCPromise(ids) {
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
}
