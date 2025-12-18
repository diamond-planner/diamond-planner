<script lang="ts">
  import {Edit} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import UniformSetForm from "$lib/dp/components/forms/UniformSetForm.svelte";
  import Cap from "$lib/dp/components/icons/Cap.svelte";
  import Pants from "$lib/dp/components/icons/Pants.svelte";
  import Shirt from "$lib/dp/components/icons/Shirt.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedUniformSet} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    uniformSet: ExpandedUniformSet;
  }

  const {uniformSet}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("uniformsets").delete(id);
    invalidate("club:single");
  }
</script>

<article class="card preset-tonal-surface shadow-md">
  <header class="card-header">
    <h3 class="h4 font-semibold">{uniformSet.name}</h3>
  </header>

  <section class="p-4">
    <div class="grid grid-cols-3 place-items-center">
      <div class="max-w-20">
        <Cap classes="w-20" identifier={uniformSet.cap}/>
      </div>

      <div>
        <Shirt classes="w-16 h-16" fillColor={uniformSet.jersey}/>
      </div>

      <div>
        <Pants classes="w-16 h-16" fillColor={uniformSet.pants}/>
      </div>
    </div>
  </section>

  <footer class="card-footer">
    {#if uniformSet?.expand?.club?.admins.includes(authRecord.id)}
      <hr class="my-2">
      <div class="flex justify-end gap-2">
        <Dialog triggerClasses="btn btn-sm btn-icon preset-tonal-tertiary border border-tertiary-500">

          {#snippet triggerContent()}
            <Edit/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2>Edit Uniform Set "{uniformSet.name}"</h2>
            </header>
          {/snippet}

          <UniformSetForm {uniformSet} clubID={uniformSet.club}/>
        </Dialog>

        <DeleteButton id={uniformSet.id} modelName="Uniform Set" action={deleteAction}/>
      </div>
    {/if}
  </footer>
</article>