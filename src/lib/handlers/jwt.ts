import type { JwtToken } from "$lib/types/home.type"
import dayjs from "dayjs"
import * as http from '$lib/handlers/api';

export class JwtManager {
    private static Instance = new JwtManager();
    private _accessToken: string;
    private _accessTokenExp: number;
    private _refreshToken: string;
    private _refreshTokenExp: number;

    private constructor() {
        this._accessToken = "";
        this._accessTokenExp = 0;
        this._refreshToken = "";
        this._refreshTokenExp = 0;
    }

    private static async refreshAccessToken() {
        const inst = this;
        return http.post("", [], this.Instance._refreshToken)
            .then(j => j as JwtToken)
            .then(j => inst.setToken(j))
            .catch(e => {
                console.log(e);
            });
    }

    public static getInstance(): JwtManager {
        return this.Instance;
    }
    public static async getAccessToken(): Promise<string | null> {
        const now  = dayjs();
        if (this.Instance._accessToken && dayjs.unix(this.Instance._accessTokenExp).isAfter(now)) {
            return this.Instance._accessToken;
        }
        if(this.Instance._refreshToken && dayjs.unix(this.Instance._refreshTokenExp).isAfter(now)) {
            await this.refreshAccessToken();
            return this.Instance._accessToken;
        }
        return null;
    }
    public static setToken(token: JwtToken) {
        this.Instance._accessToken = token.AccessToken;
        this.Instance._accessTokenExp = token.AccessTokenExpiry;
        this.Instance._refreshToken = token.RefreshToken;
        this.Instance._refreshTokenExp = token.RefreshTokenExpiry;
    }
}
