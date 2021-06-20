import firebase from 'firebase/app';
import 'firebase/firestore';
import 'firebase/auth';

const config = {
    apiKey: "<FIREBASE_API_KEY>",
    authDomain: "<YOUR_DOMAIN>.firebaseapp.com",
    databaseURL: "https://<YOUR_DOMAIN>.firebaseio.com",
    projectId: "<FIREBASE_PROJECT_ID>",
    storageBucket: "<FIREBASE_BUCKET_NAME>.appspot.com",
    messagingSenderId: "<SENDER_ID>",
    appId: "<APP_ID>"
};

export const createUserProfileDocument = async (userAuth, additionalData) => {
    if (!userAuth) return;

    const userRef = firestore.doc(`/users/${userAuth.uid}`);
    const snapshot = await userRef.get();

    if (!snapshot.exists) {
        const {displayName, email} = userAuth;
        const createdAt = new Date();

        try {
            await userRef.set({displayName, email, createdAt, ...additionalData});
        } catch (e) {
            console.log("error creating user...");
        }
    }

    return userRef;
};

firebase.initializeApp(config);

export const auth = firebase.auth();
export const firestore = firebase.firestore();

const provider = new firebase.auth.GoogleAuthProvider();
provider.setCustomParameters({prompt: 'select_account'});

export const signInWithGoogle = () => auth.signInWithPopup(provider);

export default firebase;
