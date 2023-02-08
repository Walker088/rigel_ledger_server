import type { PageServerLoad } from "./$types";

export const prerender = false;

type UserInfo = {
    userName: string,
}

export const load: PageServerLoad<UserInfo> = async () => {
    return {userName: "walker088"}
};