import type { PageServerLoad } from "./$types";
import type { JwtToken } from "$lib/types/home.type";

import { JwtManager } from '$lib/handlers/jwt';
import * as http from '$lib/handlers/api';
import { redirect } from '@sveltejs/kit';
import dayjs from "dayjs";

export const load: PageServerLoad<JwtToken> = async ({ url, locals }) => {
    const authCode = url.searchParams.get("code");
    if (!authCode) {
        throw redirect(307, `/${locals.locale}`);
    }
    const now = dayjs();
    const jwtToken = await http.get(`v1/public/oauth/github/login?code=${authCode}`)
        .then(d => ({
            AccessToken: d.accessToken,
            AccessTokenExpiry: d.accessExpiry,
            RefreshToken: d.refreshToken,
            RefreshTokenExpiry: d.refreshTokenExpiry,
        }) as JwtToken);
    const isJwtValid = jwtToken.AccessToken
        && dayjs.unix(jwtToken.AccessTokenExpiry).isAfter(now)
        && jwtToken.RefreshToken
        && dayjs.unix(jwtToken.RefreshTokenExpiry).isAfter(now);
    if (!isJwtValid) {
        throw redirect(307, `/${locals.locale}`);
    } else {
        JwtManager.setToken(jwtToken);
        throw redirect(307, `/${locals.locale}/protected`);
    }
}
