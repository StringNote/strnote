// Your web app's Firebase configuration
var firebaseConfig = {
  apiKey: "AIzaSyCOr4bemwGStuFyYgk8qrxZe_KUXIIq3X8",
  authDomain: "stringnote.firebaseapp.com",
  databaseURL: "https://stringnote.firebaseio.com",
  projectId: "stringnote",
  storageBucket: "stringnote.appspot.com",
  messagingSenderId: "447313251447",
  appId: "1:447313251447:web:8b7cbd7e514574be1dca72"
};
// Initialize Firebase
firebase.initializeApp(firebaseConfig);

var uiConfig = {
  callbacks: {
    signInSuccessWithAuthResult: function (authResult, redirectUrl) {
      return true;
    },
    uiShown: function () {
      // document.getElementById('loader').style.display = 'none';
    }
  },
  // ログイン完了時のリダイレクト先
  signInFlow: 'popup',
  signInSuccessUrl: '/',

  // 利用する認証機能
  signInOptions: [
    firebase.auth.GoogleAuthProvider.PROVIDER_ID
    // , firebase.auth.FacebookAuthProvider.PROVIDER_ID
    , firebase.auth.TwitterAuthProvider.PROVIDER_ID
    , firebase.auth.GithubAuthProvider.PROVIDER_ID
  ]
};

function getIdTokenPromise() {
  return firebase.auth().currentUser.getIdToken(true);
}
function signOutPromise() {
  return firebase.auth().signOut();
}

var onAuthStateChangedHandler = (user) => { }

firebase.auth().onAuthStateChanged((user) => {
  onAuthStateChangedHandler(user);
});

function uiStart(id) {
  // Firebase Auth UI 開始
  let ui = new firebaseui.auth.AuthUI(firebase.auth());
  ui.start("#" + id, uiConfig);
}