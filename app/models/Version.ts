export class Version {
    private version: string;
    private status: string;

    constructor({ 
        version,
        status
    }:{
        version: string, 
        status: string }) {

        this.version = version;
        this.status = status;
    }

    public getVersion(): string {
        return this.version;
    }

    public getStatus(): string {
        return this.status;
    }

    public toString(): string {
        return `{Version: ${this.version}, 
            Status: ${this.status}}`;
    }
}