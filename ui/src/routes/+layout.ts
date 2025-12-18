import type {LayoutLoad} from "./$types";
import {authSettings, client} from "$lib/dp/client.svelte.ts";
import {browser} from "$app/environment";
import {loadLocale} from "wuchale/load-utils";
import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
import {locales} from "../locales/data";
import {appLocale} from "$lib/dp/locale.svelte.ts";

export const ssr = false;

export const load: LayoutLoad = async ({fetch, depends}) => {
  let clubs: ExpandedClub[] = [];
  let teams: ExpandedTeam[] = [];

  if (browser) {
    let locale: string;
    if (appLocale.current) {
      // locale has been set before
      locale = appLocale.current;
    } else {
      // locale is unset so far, try to read from browser settings
      const browserPreferredLocale = navigator.language.slice(0, 2); // Safari has `de-DE`, Chrome and Firefox use `de`

      if (locales.includes(browserPreferredLocale)) {
        locale = browserPreferredLocale;
      } else {
        locale = "en";
      }
    }
    await loadLocale(locale);
  }

  if (browser && client.authStore.isValid) {
    /**
     * setting requestKey to `null` means "do not autocancel" => nav is always up to date
     */
    teams = await client.collection("teams").getFullList<ExpandedTeam>({
      expand: "club,admins",
      fetch: fetch,
      requestKey: null,
      sort: "+name",
    });

    const authRecord = authSettings.record as CustomAuthModel;
    clubs = await client.collection("clubs").getFullList<ExpandedClub>({
      filter: `"${authRecord?.club}" ?~ id`,
      fetch: fetch,
      requestKey: null,
      expand: "admins",
    });
  }
  depends("nav:load");

  return {
    clubs: clubs,
    teams: teams,
  };
};
