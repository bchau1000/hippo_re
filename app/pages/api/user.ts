import { firebaseAuth } from "./firebase";
import { Auth, createUserWithEmailAndPassword, User, UserCredential } from "firebase/auth";

const loginUser = () => {
    
}

const registerUser =  (auth:Auth, email:string, password:string):User|never => {
    let errorCode:Nullable<string> = null;
    let errorMessage:Nullable<string> = null;

    createUserWithEmailAndPassword(auth, email, password)
        .then((userCredential:UserCredential) => {
            const user:User = userCredential.user;
            return user;
        })
        .catch((error) => {
            errorCode = error.errorCode;
            errorMessage = error.errorMessage;
        });

    throw new Error(`${errorCode} - Error occurred while registering user: ${errorMessage}`)
}