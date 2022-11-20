import {apiUrl} from './constant'

const pingUrl:URL = new URL(`${apiUrl}/ping`);

export const pingServer = async ():Promise<Response> => {
    const info:RequestInit = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: "include"
    };
    const response:Response =  await fetch(pingUrl, info);
    return response;
}
