export class UserModel {
    private id:number;
    private uid:string;
    private email:string;
    private displayName:string;

    constructor({
        id,
        uid,
        email,
        displayName,
    }:{
        id:number,
        uid:string,
        email:string,
        displayName:string,
    }) {
        this.id = id;
        this.uid = uid;
        this.email = email;
        this.displayName = displayName;
    }

    public getId():number {
        return this.id;
    }

    public getUid():string {
        return this.uid;
    }

    public getEmail():string {
        return this.email;
    }

    public getDisplayName():string {
        return this.displayName;
    }

    public toJSON():string {
        return JSON.stringify({
            id: this.id,
            uid: this.uid,
            email: this.email,
            displayName: this.displayName,
        })
    }
}