import * as http from "$lib/handlers/api";
import { JwtManager } from "$lib/handlers/jwt";
import type { LoginPageInfo } from "$lib/types/home.type";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const prerender = true;

export const load: PageServerLoad<LoginPageInfo> = async ({locals}) => {
    const accessToken = await JwtManager.getAccessToken();
    if (accessToken) {
        throw redirect(301, `/${locals.locale}/protected`);
    }
    const homePageInfo = await http.get(`v1/public/home`, [])
        .then(d => d as LoginPageInfo);
    return homePageInfo
};
