import { firebaseAuth } from "./firebase";

const GetVersion = async ():Promise<Response> => {
    const info:RequestInit = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: "include"
    };
    const url:URL = new URL("http://127.0.0.1:4000/hippo/api/version");
    const response:Response =  await fetch(url, info);
    return response;
}

export {
    GetVersion
}