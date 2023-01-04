
export class JwtManager {
    private static Instance = new JwtManager();
    private token:string;

    private constructor() {
        this.token = "";
    }

    public static getInstance(): JwtManager {
        return this.Instance;
    }
    public static getToken() {
        return this.Instance.token;
    }
    public static setToken(token: string) {
        this.Instance.token = token;
    }
}
