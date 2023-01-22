import * as http from '$lib/handlers/api';
import { JwtManager } from "$lib/handlers/jwt"
import type { LoginPageInfo } from "$lib/types/home.type"
import type { PageServerLoad } from './$types'

export const prerender = true;

export const load: PageServerLoad<LoginPageInfo> = async () => {
    const homePageInfo = await http.get(`v1/public/home`, [], JwtManager.getToken())
        .then(d => d as LoginPageInfo);
    return homePageInfo
};
