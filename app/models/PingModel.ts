export class PingModel {
    private version: string;
    private status: string;
    private message: string;

    constructor({ 
        version,
        status,
        message,
    }:{
        version: string, 
        status: string,
        message: string,
    }) {

        this.version = version;
        this.status = status;
        this.message = message;
    }

    public getVersion(): string {
        return this.version;
    }

    public getStatus(): string {
        return this.status;
    }

    public getMessage(): string {
        return this.message;
    }

    public toJSON(): string {
        return JSON.stringify({
            version: this.version,
            status: this.status,
            message: this.message,
        });
    }
}