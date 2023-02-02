import type { JwtToken } from "$lib/types/home.type"
import type { Dayjs } from "dayjs"
import dayjs from "dayjs"
import * as http from '$lib/handlers/api';

export class JwtManager {
    private static Instance = new JwtManager();
    private _accessToken: string;
    private _accessTokenExp: Dayjs | null;
    private _refreshToken: string;
    private _refreshTokenExp: Dayjs | null;

    private constructor() {
        this._accessToken = "";
        this._accessTokenExp = null;
        this._refreshToken = "";
        this._refreshTokenExp = null;
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
        if (this.Instance._accessToken && this.Instance._accessTokenExp?.isAfter(now)) {
            return this.Instance._accessToken;
        }
        if(this.Instance._refreshToken && this.Instance._refreshTokenExp?.isAfter(now)) {
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
