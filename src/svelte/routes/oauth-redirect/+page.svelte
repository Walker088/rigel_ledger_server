<script lang="ts">
	import type { PageData } from "./$types";
	import { onMount } from "svelte";
	import { JwtManager } from "$lib/handlers/jwt";
    import dayjs from "dayjs";
	import { error } from "@sveltejs/kit";

    export let data: PageData

    onMount(async () => {
        if(data) {
            const now = dayjs();
            const {AccessToken, AccessTokenExpiry, RefreshToken, RefreshTokenExpiry} = data;
            const isJwtValid = AccessToken
                && dayjs.unix(AccessTokenExpiry).isAfter(now)
                && RefreshToken
                && dayjs.unix(RefreshTokenExpiry).isAfter(now);
            if (isJwtValid) {
                JwtManager.setToken({
                    AccessToken, 
                    AccessTokenExpiry,
                    RefreshToken,
                    RefreshTokenExpiry
                });
                const userId = JwtManager.getUidFromToken(AccessToken);
                location.href = `/${data.locale}/protected/${userId}`;
            } else {
                throw error(401, {message: "The jwt token is invalid"});
            }
        }
    });
</script>