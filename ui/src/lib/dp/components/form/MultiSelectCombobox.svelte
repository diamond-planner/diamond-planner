<script generics="T extends UsersResponse" lang="ts">
  //@ts-expect-error

  import {X} from "lucide-svelte";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";

  // TODO: this should be even more generic

  interface Props {
    itemName: string;
    selectedItems: T[];
    allItems: T[];
    allowDeletionOfLastItem?: boolean;
  }

  const {
    itemName,
    selectedItems = $bindable(),
    allItems,
    allowDeletionOfLastItem = false,
  }: Props = $props();

  let selectElement: HTMLSelectElement | undefined = $state();

  function addItemToSelection(users: T[]) {
    if (!selectElement || selectElement?.value === "") {
      return;
    }
    const selectedUser = users.find((user) => user.id === selectElement?.value);
    const adminExists = selectedItems.find((admin) => admin.id === selectedUser?.id);

    if (selectedUser && selectElement && !adminExists) {
      selectedItems.push(selectedUser);
    }
    selectElement.value = "";
  }

  function removeItemFromSelection(itemToRemove: T) {
    if (!allowDeletionOfLastItem && selectedItems.length === 1) {
      toastController.trigger({
        message: `You cannot remove the last ${itemName}!`,
        background: "preset-filled-warning-500",
      });
      return;
    }
    const itemRef = selectedItems.find((entry) => entry.id === itemToRemove.id);

    if (itemRef) {
      const index = selectedItems.indexOf(itemRef);

      if (index !== -1) {
        selectedItems.splice(index, 1);
      }
    }
  }
</script>

{#each selectedItems as selectItem}
  <button
          type="button"
          class="chip preset-filled-primary-500 me-1 lg:me-2"
          onclick={() => removeItemFromSelection(selectItem)}
  >
    <span>{selectItem.first_name} {selectItem.last_name}</span>
    <X size="12"/>
  </button>
{/each}

<div>Select to add as {itemName}:</div>
<select
        bind:this={selectElement}
        class="select"
        onchange={() => addItemToSelection(allItems)}
>
  <option selected value="">None</option>
  {#each allItems as item}
    <option value={item.id}>{item.first_name} {item.last_name}</option>
  {/each}
</select>
