export class AuthModel {
    private idToken:string;

    public constructor({
        idToken
    }:{
        idToken:string
    }) {
        this.idToken = idToken;
    }

    public getIdToken():string {
        return this.idToken;
    }

    public toJSON():string {
        return JSON.stringify({
            idToken: this.idToken,
        });
    }
}