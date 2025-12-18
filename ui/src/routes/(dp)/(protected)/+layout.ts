import {redirect} from "@sveltejs/kit";
import {browser, dev} from "$app/environment";
import {authSettings, client} from "$lib/dp/client.svelte.ts";
import {toastController} from "$lib/dp/service/ToastController.svelte.ts";

export const ssr = false;

/**
 * This does not load any data, instead it acts as an authentication check for all protected routes.
 * @TODO: hacky solution
 */
export const load = ({url}) => {
  if (browser) {
    // we are referencing this here because load functions need to read at least one parameter
    // that changes on navigation to re-run (this function needs to run on every route).
    const {pathname} = url;

    if (dev) {
      console.log(pathname);
    }

    if (!client.authStore.isValid) {
      toastController.trigger({
        message: "You need to be logged in to access this page (session expired/never logged in).",
        background: "preset-filled-error-500",
      });
      authSettings.record = null;
      redirect(303, "/login");
    }
  }
};
