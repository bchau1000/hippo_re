import { firebaseAuth } from "./firebase";
import { createUserWithEmailAndPassword, signInWithEmailAndPassword, User, UserCredential, AuthError } from "firebase/auth";
import { apiUrl } from "./constant";
import { AuthModel, UserModel } from "../../models";

const authUserURL:URL = new URL(`${apiUrl}/user/auth`);
const loginUserURL:URL = new URL(`${apiUrl}/user/login`);

export const loginUser = async (email:string, password:string):Promise<Response> => {
    const response:UserCredential = await signInWithEmailAndPassword(firebaseAuth(), email, password);

    return signInWithEmailAndPassword(firebaseAuth(), email, password)
        .then(async (userCred:UserCredential) => {
            const info:RequestInit = {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "include",
                body: new AuthModel({
                    idToken: await userCred.user.getIdToken(),
                }).toJSON(),
            };

            const response:Response = await fetch(loginUserURL, info)
        
            return response;
        });
}

export const registerUser = (email:string, password:string):User|never => {
    let err:any = null;

    createUserWithEmailAndPassword(firebaseAuth(), email, password)
        .then((userCredential:UserCredential) => {
            const user:User = userCredential.user;
            return user;
        })
        .catch((error) => err = error);

    throw new Error(`${err.errorCode} - Error occurred while registering user: ${err.Message}`)
}

export const authUser = async ():Promise<boolean> => {
    const info:RequestInit = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: "include",
        body: "",
    };
    const response:Response = await fetch(authUserURL, info);
    return response.status === 200;
}