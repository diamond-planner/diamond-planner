<script lang="ts">
  import {fade, slide} from "svelte/transition";
  import {goto} from "$app/navigation";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  import OAuthProviderButton from "$lib/dp/auth/OAuthProviderButton.svelte";
  import PasswordRequestButton from "$lib/dp/auth/PasswordRequestButton.svelte";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {Extension} from "$lib/dp/types/ExpandedResponse.js";
  import type {UsersUpdate} from "$lib/dp/types/pb-types.ts";
  import type {Toast} from "$lib/dp/types/Toast.ts";
  import {client} from "../client.svelte.js";
  //@ts-ignore
  import {Tabs} from "bits-ui";

  const {authCollection = "users", passwordLogin = true, signupAllowed = true} = $props();

  const coll = $derived(client.collection(authCollection));

  const failSettings: Toast = {
    message: "There was an error processing your authentication request.",
    background: "preset-filled-error-500",
  };

  let form: Extension<Partial<UsersUpdate>, { signup_key: string }> = $state({
    email: "",
    password: "",
    first_name: "",
    last_name: "",
    passwordConfirm: "",
    signup_key: "",
  });
  let signup = false;

  async function submit(e: SubmitEvent) {
    e.preventDefault();

    if (signup) {
      try {
        await coll.create({...form});
        const signupSuccessful = await coll.requestVerification(form.email ?? "");

        if (signupSuccessful) {
          await goto("/signupconfirm");
        } else {
          toastController.trigger(failSettings);
        }
      } catch {
        toastController.trigger(failSettings);
      }
    } else {
      try {
        const authResponse = await coll.authWithPassword(form.email ?? "", form.password ?? "", {
          expand: "club",
        });

        if (authResponse) {
          await goto("/account", {invalidateAll: true});
        }
      } catch (error) {
        console.error(error);
        toastController.trigger(failSettings);
      }
    }
  }

  let tabSet: "login" | "signup" | string = $state("login");
  let forgotPassword = $state(false);
</script>

{#snippet login()}
  <label class="label">
    <span>Your email</span>
    <input
      class="input"
      bind:value={form.email}
      required
      type="email"
      placeholder="name@provider.com"
      autocomplete="email"
    />
  </label>

  {#if !forgotPassword}
    <label class="label" transition:slide>
      <span>Your password</span>
      <input
        bind:value={form.password}
        class="input"
        required
        type="password"
        placeholder="**********"
        minlength="8"
        maxlength="72"
        autocomplete="current-password"
      />
    </label>

    <button
      transition:slide
      class="btn preset-tonal-primary border border-primary-500 block mt-4"
      type="submit"
      onclick={() => (signup = false)}
      disabled={form.email === "" || form.password === ""}
    >
      Login to your account
    </button>
  {/if}

  <div class="block md:flex gap-6 items-center justify-start my-6">
    <Switch
      name="forgot password"
      onCheckedChange={(e) => (forgotPassword = e.checked)}
    >
      Forgot Password?
    </Switch>

    {#if forgotPassword}
      <div class="mt-2 md:mt-0" transition:fade>
        <PasswordRequestButton
          email={form.email ?? ""}
          disabled={form.email === ""}
          classes="btn anchor p-0"
        >
          Reset Password
        </PasswordRequestButton>
      </div>
    {/if}
  </div>
{/snippet}

<div class="outer-wrapper">
  <div class="inner-wrapper">
    <article>
      <header>
        <h1 class="h2">Sign in to Diamond Planner</h1>
      </header>

      <form onsubmit={submit}>
        {#if passwordLogin}
          {#if signupAllowed}
            <Tabs.Root
              bind:value={tabSet}
              class="my-2"
            >
              <Tabs.List
                class="tabs-list preset-tonal-surface"
              >
                <Tabs.Trigger
                  value="login"
                  class="tabs-trigger btn"
                >Log In
                </Tabs.Trigger>
                <Tabs.Trigger
                  value="signup"
                  class="tabs-trigger btn"
                >Create Account
                </Tabs.Trigger>
              </Tabs.List>

              <Tabs.Content value="login" class="pt-3">
                <h2>Login</h2>

                {@render login()}
              </Tabs.Content>

              <Tabs.Content value="signup" class="pt-3">

                {@render login()}

                {#if tabSet === "signup"}
                  <label class="label">
                    <span class="">Your email</span>
                    <input
                      class="input"
                      bind:value={form.email}
                      required
                      type="email"
                      placeholder="name@provider.com"
                      autocomplete="email"
                    />
                  </label>

                  <div class="grid sm:grid-cols-2 gap-2">

                    <label class="label">
                      <span class="">First Name</span>
                      <input
                        class="input"
                        bind:value={form.first_name}
                        placeholder="John"
                        required
                        type="text"
                        autocomplete="given-name"
                      />
                    </label>

                    <label class="label">
                      <span class="">Last Name</span>
                      <input
                        class="input"
                        bind:value={form.last_name}
                        placeholder="Doe"
                        required
                        type="text"
                        autocomplete="family-name"
                      />
                    </label>
                  </div>

                  <div class="grid sm:grid-cols-2 gap-2">

                    <label class="label">
                      <span class="">Your password</span>
                      <input
                        class="input"
                        bind:value={form.password}
                        required
                        type="password"
                        placeholder="**********"
                        minlength="8"
                        maxlength="72"
                        autocomplete="new-password"
                      />
                    </label>

                    <label class="label">
                      <span class="">Confirm password</span>
                      <input
                        class="input"
                        bind:value={form.passwordConfirm}
                        required
                        type="password"
                        placeholder="confirm password"
                        minlength="8"
                        maxlength="72"
                        autocomplete="new-password"
                      />
                    </label>
                  </div>

                  <label class="label">
                    <span class="">Signup Key</span>
                    <input
                      bind:value={form.signup_key}
                      class="input"
                      name="signup_key"
                      placeholder="minimum 8 characters"
                      minlength="8"
                      required
                      type="text"
                      autocomplete="one-time-code"
                    />
                    <span class="block text-sm font-light">
                                          A valid signup key needs to be entered upon user account creation.
                                          If you do not have a signup key, please contact your team manager.
                                      </span>
                  </label>

                  <input type="hidden" name="register" value={true}/>

                  <button
                    class="btn preset-tonal-primary border border-primary-500 my-2"
                    type="submit"
                    onclick={() => (signup = true)}
                    disabled={form.email === "" || form.password === "" || form.passwordConfirm === "" || form.signup_key === ""}
                  >
                    Register new account
                  </button>
                {/if}
              </Tabs.Content>

            </Tabs.Root>

          {/if}
        {/if}

        {#await coll.listAuthMethods({requestKey: null}) then methods}
          {#if methods.oauth2.providers.length > 0}
            <hr class="my-2">

            <div class="mx-2 mt-3 text-surface-700-300 font-light flex justify-center">
              <span>or sign in with</span>
            </div>

            <div class="my-3 md:my-5 flex flex-wrap gap-6 justify-center">
              {#each methods.oauth2.providers as provider}

                {#if tabSet === "login"}
                  <!--Login Buttons - no signupKey-->
                  <OAuthProviderButton authProvider={provider} collection={coll} disabled={false}/>
                {:else}
                  <!--Signup Buttons - with signupKey state prop-->
                  <OAuthProviderButton authProvider={provider} collection={coll}
                                       signup_key={form.signup_key} disabled={form.signup_key === ""}/>
                {/if}

              {/each}
            </div>

            {#if tabSet === "signup" && form.signup_key === ""}
              <div class="mx-2 mt-3 text-surface-700-300 font-light flex justify-center" transition:slide>
              <span>
                Even when using an external login provider,
                a signup key (see above) is still required for account creation.
              </span>
              </div>
            {/if}
            <hr class="my-6">

            <p class="font-light text-sm mt-10">
              Note: it is possible to associate a local account with an
              external provider later by using the same email address. A single account can be associated with more
              than one external provider.
            </p>
          {/if}
        {:catch error}
          <!-- pocketbase not working -->
          <p>There seems to be an error while contacting the backend. Please try again later.</p>
        {/await}
      </form>
    </article>
  </div>

</div>

<style>
  .outer-wrapper {
    display: flex;
    justify-content: center;
  }

  .inner-wrapper {
    display: flex;
    justify-content: center;
    margin-block-start: calc(var(--spacing) * 3) !important;
    margin-block-end: calc(var(--spacing) * 10) !important;

    @media (min-width: 48rem) {
      max-width: 60%;
    }

    @media (min-width: 64rem) {
      max-width: 55%;
    }
  }

  .label {
    margin-block: calc(var(--spacing) * 1);
    @media (width >= 48rem /* 768px */
    ) {
      margin-block: calc(var(--spacing) * 2) /* 0.5rem = 8px */
    ;
    }
  }

  header {
    margin-bottom: 2em;
  }
</style>
