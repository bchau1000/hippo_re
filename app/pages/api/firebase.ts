import { FirebaseApp, FirebaseOptions, initializeApp,  getApp} from 'firebase/app';
import { Auth, getAuth } from "firebase/auth";

// TODO: Replace the following with your app's Firebase project configuration
const firebaseConfig = {
    apiKey: "AIzaSyC3mclhmR80s6gReKsWXPVoUA7sB_EOMSw",
    authDomain: "hippo-3b249.firebaseapp.com",
    projectId: "hippo-3b249",
    storageBucket: "hippo-3b249.appspot.com",
    messagingSenderId: "264264475909",
    appId: "1:264264475909:web:db138310da6a959c5ec36c",
    measurementId: "G-D99EENT984"
  };
const firebaseInstance = ():FirebaseApp => initializeApp(firebaseConfig);

export const firebaseAuth= ():Auth => {
    if(firebaseInstance.length === 0) {
        return getAuth(firebaseInstance())
    } else {
        return getAuth(getApp());
    }
};

