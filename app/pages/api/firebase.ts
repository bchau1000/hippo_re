import { FirebaseApp, FirebaseOptions, initializeApp } from 'firebase/app';
import { Auth, getAuth } from "firebase/auth";

// TODO: Replace the following with your app's Firebase project configuration
const firebaseConfig:FirebaseOptions = {
};
const firebaseInstance:FirebaseApp = initializeApp(firebaseConfig);

export const firebaseAuth:Auth = getAuth(firebaseInstance);

