/**
 * 簡易コンポーネント作成ユーティリティ
 * @licence MIT
 */
class TinyComp extends HTMLElement {
	// 表示更新メソッド
	update() {
		// スタイル定義の変更時の適用
		let style = this._style;
		if (style) {
			let styleElem = this.shadowRoot.querySelector('style');
			let cur = styleElem.textContent;
			if (cur!=style) {
				styleElem.textContent = style;
			}
		}
		// コンテンツ更新ハンドルの実行
		if (this._updateContents) {
			this._updateContents();
		}
	}
	// 基底コンストラクタ
	constructor() {
		super();
		const shadow = this.attachShadow({mode: 'open'});
		// コンテンツ生成
		const template = document.getElementById(this.constructor.tagName);
		if (template) {
			// テンプレートから生成
			shadow.appendChild(template.content.cloneNode(true));
		} else {
			// エレメント生成
			shadow.appendChild(MakeElement(this._contents));
		}
		// Style 生成
		if (!shadow.querySelector('style')) {
			shadow.appendChild(document.createElement('style'));
		}
		this.update();
	}
	// 監視属性リストをマージして返す
	static get observedAttributes() {
		// クラスの _attributes を取得
		let attr = this._attributes;
		if (!attr) attr = [];
		// スーパークラスの observedAttributes をマージ
		const proto = Object.getPrototypeOf(this).observedAttributes;
		return attr.concat(proto);
	}
	// 監視属性更新イベントハンドル
	attributeChangedCallback(name, oldValue, newValue) {
		this.update();
	}
/*
	// Class.tagName でタグ名を登録
	// 関連付けの簡略化のため、テンプレートのIDと共用させる
	//     → <template id="component-name">...</template>
	static tagName = "component-name";

	// スタイル定義
	_style = "";
	get _style() { return `...`; }
	// コンテンツ定義 (tagName と同じIDのテンプレートが無い場合)
	_contents = {tag: 'span', props:{textContent:""}, child:{tag: 'span', ...}};
	// コンテンツ更新ハンドル
	_updateContents() {}

	// 監視属性リスト
	static _attributes = [];

	// DOM 接続イベント
	connectedCallback() {}
	// DOM 切断イベント
	disconnectedCallback() {}
	// ドキュメント移動イベント
	adoptedCallback() {}
*/
};
// 独自表現のオブジェクトからコンテンツのエレメントを作成する。
function MakeElement(contents) {
	if (!contents) {
		const mes = "contents was undefined.";
		console.log(mes);
		contents = {
			tag:"span"
			, style:"color:red; font-size:xx-large;"
			, props:{textContent: mes}
		};
	}
	// エレメント生成
	let elem = document.createElement(contents.tag);
	// 属性(srcなど)の設定
	Object.keys(contents).forEach((name) => {
		if (name!='tag' && name!='props' && name!='child') {
			elem.setAttribute(name, contents[name]);
		}
	});
	if (contents.props) {
		// プロパティ(elem.textContentなど)の設定
		Object.keys(contents.props).forEach((name) => {
			elem[name] = contents.props[name];
		});
	}
	if (contents.child) {
		// 子要素の埋め込み
		if (Array.isArray(contents.child)) {
			contents.child.forEach((child) => { elem.appendChild(MakeElement(child)); });
		} else {
			elem.appendChild(MakeElement(contents.child));
		}
	}
	return elem;
};
function ComponentDefine(cls) {
	customElements.define(cls.tagName, cls);
};
