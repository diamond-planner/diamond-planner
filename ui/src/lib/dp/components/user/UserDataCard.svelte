<script lang="ts">
  import {Mail, User} from "lucide-svelte";
  import Avatar from "$lib/dp/components/utils/Avatar.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    model: CustomAuthModel;
  }

  const {model}: Props = $props();
</script>

<div class="card preset-tonal-surface lg:col-span-2 shadow-lg">
  <header class="card-header">
    <h2 class="h4 font-semibold">User Data</h2>
  </header>

  <section class="p-4 flex flex-col sm:flex-row gap-4 lg:gap-12 sm:items-center">
    <Avatar
      --size="4rem"
      fallback={`${model.first_name.charAt(0)?.toUpperCase()}${model.last_name.charAt(0)?.toUpperCase()}`}
      src={client.files.getURL(model, model?.avatar)}
    />

    <div class="grid grid-cols-1 gap-2 col-span-4">
      <div class="flex items-center gap-3">
        <User/>
        <div>
          <p>{`${model?.first_name ?? ""} ${model?.last_name ?? ""}`}</p>
          <p class="text-sm font-light">Name</p>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <Mail/>
        <div>
          <p>{model?.email}</p>
          <p class="text-sm font-light">email address</p>
        </div>
      </div>
    </div>
  </section>

  <footer class="card-footer">
    <p>Edit User Data:</p>
    <a class="btn preset-tonal-primary border border-primary-500" href="/account/{model.id}/settings"
       title="Go to User Settings">User
      Settings</a>
  </footer>
</div>

<style>
  .card-footer {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: calc(var(--spacing) * 4);
  }
</style>
