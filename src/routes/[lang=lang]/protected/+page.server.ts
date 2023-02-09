import type { PageServerLoad } from "./$types";
import type { Dayjs } from "dayjs";
import * as http from "$lib/handlers/api";
import { JwtManager } from "$lib/handlers/jwt";

export const prerender = false;

type UserInfo = {
    userId: string,
    userName: string,
    userMail: string,
    userTypeCode: number,
    userTypeName: string,
    mainCountry: string,
    mainCurrency: string,
    mainLanguage: Locales,

    ledgerSummary: {
        balance: {
            assets: number,
            liabilityies: number,
            networth: number,
        },
        income: {
            monthlyIncome: number,
            monthlyExpence: number,
        },
        transactions: {
            availableRange: {
                firstAt: Dayjs,
                lastAt: Dayjs,
            },
            monthlyTxs: {
                transacId: number,
                transacDate: Dayjs,
                debitAccountCode: string,
                debitAccountName: string,
                creditAccountCode: string,
                creditAccountName: string,
                currency: string,
                originalAmt: number,
                adjustedAmt: number, // in user main currency
            }[],
        }
    }
}

export const load: PageServerLoad<UserInfo> = async () => {
    const token = await JwtManager.getAccessToken() || undefined;
    await http.get(`v1/public/home`, [], token);
    return {userName: "walker088"}
};
