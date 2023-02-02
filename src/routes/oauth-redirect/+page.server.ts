import type { PageServerLoad } from "./$types";
import type { JwtToken } from "$lib/types/home.type";

import * as http from '$lib/handlers/api';
import { redirect } from '@sveltejs/kit';
import dayjs from "dayjs";

export const load: PageServerLoad<JwtToken> = async ({ url, locals }) => {
    const authCode = url.searchParams.get("code");
    if (!authCode) {
        throw redirect(307, `/${locals.locale}`);
    }
    const now = dayjs();
    const jwtToken = await http.get(`/oauth/github/login?code=${authCode}`)
        .then(d => ({
            AccessToken: d.AccessToken,
            AccessTokenExpiry: dayjs(d.AccessTokenExpiry),
            RefreshToken: d.RefreshToken,
            RefreshTokenExpiry: dayjs(d.RefreshTokenExpiry),
        }));
    const isJwtValid = jwtToken.AccessToken
        && jwtToken.AccessTokenExpiry.isAfter(now)
        && jwtToken.RefreshToken
        && jwtToken.RefreshTokenExpiry.isAfter(now);
    if (!isJwtValid) {
        throw redirect(307, `/${locals.locale}`);
    }
    return jwtToken;
}
