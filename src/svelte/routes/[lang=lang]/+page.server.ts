import * as http from "$lib/handlers/api";
import { JwtManager } from "$lib/handlers/jwt";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

type LoginPageInfo = {
    apiVersion: string,
    changeLogMd: string,
    ghOauthUrl?: string,
}

export const load: PageServerLoad<LoginPageInfo> = async ({locals, url}) => {
    const accessToken = await JwtManager.getAccessToken();
    if (accessToken) {
        throw redirect(301, `/${locals.locale}/protected`);
    }
    const homePageInfo = await http.get(`v1/public/home`, [])
        .then(d => d as LoginPageInfo);

    const ghAuthUrl = import.meta.env.VITE_GH_AUTH_URL;
    const ghClientId = import.meta.env.VITE_GH_CLIENT_ID;
    const ghRedirect = `http://${url.host}/oauth-redirect`;
    homePageInfo.ghOauthUrl = `${ghAuthUrl}?client_id=${ghClientId}&redirect_uri=${ghRedirect}&scope=user:email`;
    return homePageInfo;
};
