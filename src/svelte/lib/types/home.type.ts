
export type MainLayoutInfo = {
    locale: Locales,
};

export type JwtToken = {
    AccessToken: string,
    AccessTokenExpiry: number,
    RefreshToken: string,
    RefreshTokenExpiry: number,
};
