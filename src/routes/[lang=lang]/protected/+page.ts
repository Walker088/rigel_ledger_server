import type { PageLoad } from "./$types";

export const prerender = false;

export const load: PageLoad = async () => {
    return {userName: "walker088"}
};