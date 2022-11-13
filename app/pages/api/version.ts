import { Version } from "../../models/Version";

const GetVersion = async ():Promise<any> => {
    const info:RequestInit = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: "include"
    };
    const url:URL = new URL("http://127.0.0.1:4000/hippo/api/version");
    const response:Response =  await fetch(url, info);
    if(response.status == 200) {
        return await response.json()
    } else {
        return null
    }
}

export {
    GetVersion
}